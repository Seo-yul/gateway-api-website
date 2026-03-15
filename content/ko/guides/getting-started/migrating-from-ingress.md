---
title: "Ingress에서 마이그레이션"
weight: 3
description: "A general guide for migrating from any Ingress implementation to Gateway API"
---

Gateway API는 [Ingress API][ing]의 후속이다. 이 가이드는 모든 Ingress 구현에서 Gateway API로 마이그레이션하기 위한 일반적인 개요를 제공한다. Ingress와 Gateway API 사이의 핵심 개념, 차이점, 기능 매핑에 초점을 맞추고 수동 변환 예제를 제공한다.

[ing]:https://kubernetes.io/docs/concepts/services-networking/ingress/

{{< note >}}
**Ingress-NGINX 사용자인가?**

Ingress-NGINX를 사용하고 있다면, 더 맞춤화된 조언과 리소스를 위해 [Ingress-NGINX 마이그레이션 가이드]({{< ref "/guides/getting-started/migrating-from-ingress-nginx" >}})도 확인해 보기 바란다.
{{< /note >}}

이 가이드는 변환을 돕는다. 다음 내용을 다룬다:

*   왜 Gateway API로 전환하는 것이 좋은지 설명한다.
*   Ingress API와 Gateway API 사이의 주요 차이점을 설명한다.
*   Ingress 기능을 Gateway API 기능에 매핑한다.
*   Ingress 리소스를 Gateway API 리소스로 변환하는 예제를 보여준다.
*   자동 변환을 위한 [ingress2gateway](https://github.com/kubernetes-sigs/ingress2gateway)를
    소개한다.

동시에, 이 가이드는 실시간 마이그레이션을 준비하거나 Ingress 컨트롤러의
구현별 기능을 변환하는 방법은 설명하지 않는다.
또한, Ingress API는 HTTP/HTTPS 트래픽만 다루므로, 이 가이드는
다른 프로토콜에 대한 Gateway API 지원을 다루지 않는다.

## Gateway API로 전환하는 이유

[Ingress API](https://kubernetes.io/docs/concepts/services-networking/ingress/)는
서비스에 대한 외부 HTTP/HTTPS 로드 밸런싱을 구성하는 표준 쿠버네티스 방법이다.
쿠버네티스 사용자에 의해 널리 채택되었으며, 많은 구현체([Ingress 컨트롤러](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/))가
제공되어 벤더들의 지원이 잘 되고 있다. 또한, [cert-manager](https://cert-manager.io/)와
[ExternalDNS](https://github.com/kubernetes-sigs/external-dns)와 같은 여러
클라우드 네이티브 프로젝트가 Ingress API와 통합된다.

그러나, Ingress API에는 몇 가지 한계가 있다:

- *제한된 기능*. Ingress API는 TLS 종료와 HTTP 트래픽의 간단한
  콘텐츠 기반 요청 라우팅만 지원한다.
- *확장성에 대한 어노테이션 의존*. 어노테이션 접근 방식의 확장성은
  제한된 이식성으로 이어지며, 각 구현체마다 고유한 지원 확장이 있어
  다른 구현체로 변환되지 않을 수 있다.
- *불충분한 권한 모델*. Ingress API는 공유 로드 밸런싱 인프라를 가진
  다중 팀 클러스터에 적합하지 않다.

Gateway API는 이러한 한계를 해결하며, 다음 섹션에서 보여줄 것이다.

> Gateway API의 [설계 목표]({{< ref "/overview#gateway-api-concepts" >}})에 대해
> 자세히 읽어보라.

## Ingress API와 Gateway API의 주요 차이점

Ingress API와 Gateway API 사이에는 세 가지 주요 차이점이 있다:

*   페르소나
*   사용 가능한 기능
*   확장성 접근 방식 (구현별 기능)

### 페르소나

처음에 Ingress API에는 Ingress라는 단일 리소스 종류만 있었다. 그 결과,
Ingress 리소스의 소유자인 사용자라는 하나의 페르소나만 있었다.
Ingress 기능은 사용자에게 TLS 종료 구성 및 로드 밸런싱 인프라
프로비저닝(일부 Ingress 컨트롤러에서 지원) 등 애플리케이션이 외부
클라이언트에 노출되는 방식에 대한 많은 제어를 제공한다.
이러한 수준의 제어를 셀프 서비스 모델이라고 한다.

동시에, Ingress API에는 Ingress 컨트롤러를 프로비저닝하고 관리하는
담당자를 설명하는 두 개의 암묵적 페르소나도 포함되어 있었다:
제공자 관리 Ingress 컨트롤러를 위한 인프라 제공자와 자체 호스팅
Ingress 컨트롤러를 위한 클러스터 운영자(또는 관리자)이다.
[IngressClass](https://kubernetes.io/docs/concepts/services-networking/ingress/#ingress-class)
리소스가 나중에 추가되면서, 인프라 제공자와 클러스터 운영자는
해당 리소스의 소유자가 되었고, 따라서 Ingress API의 명시적 페르소나가
되었다.

Gateway API에는
[네 가지 명시적 페르소나]({{< ref "/overview/concepts/security#roles-and-personas" >}})가
포함된다: 애플리케이션 개발자, 애플리케이션 관리자, 클러스터 운영자,
인프라 제공자. 이를 통해 사용자 페르소나의 책임을 이러한 페르소나들에게
분배하여(인프라 제공자 제외) 셀프 서비스 모델에서 벗어날 수 있다:

*   클러스터 운영자/애플리케이션 관리자가 TLS 종료 구성을 포함한
    외부 클라이언트 트래픽의 진입점을 정의한다.
*   애플리케이션 개발자가 해당 진입점에 연결되는 애플리케이션의
    라우팅 규칙을 정의한다.

이러한 분리는 여러 팀이 동일한 로드 밸런싱 인프라를 공유하는 일반적인
조직 구조를 따른다. 동시에, 셀프 서비스 모델을 포기할 필요는 없다 --
애플리케이션 개발자, 애플리케이션 관리자, 클러스터 운영자 책임을
충족하는 단일 RBAC 역할을 구성하는 것이 여전히 가능하다.

아래 표는 Ingress API와 Gateway API 페르소나 간의 매핑을 요약한다:

| Ingress API 페르소나 | Gateway API 페르소나 |
|-|-|
| 사용자 | 애플리케이션 개발자, 애플리케이션 관리자, 클러스터 운영자 |
| 클러스터 운영자 | 클러스터 운영자 |
| 인프라 제공자 | 인프라 제공자 |

### 사용 가능한 기능

Ingress API는 기본 기능만 제공한다: TLS 종료와 호스트 헤더 및 요청의
URI를 기반으로 한 HTTP 트래픽의 콘텐츠 기반 라우팅이다. 더 많은 기능을
제공하기 위해, Ingress 컨트롤러는 Ingress 리소스의
[어노테이션][anns]을 통해 지원하며, 이는 Ingress API의 구현별 확장이다.

[anns]:https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
어노테이션 접근 방식의 확장성은 Ingress API 사용자에게 두 가지 부정적인
결과를 가져온다:

*   *제한된 이식성*. 너무 많은 기능이 어노테이션을 통해 제공되기 때문에,
    Ingress 컨트롤러 간 전환이 어렵거나 불가능해진다.
    한 구현의 어노테이션을 다른 것으로 변환해야 하기 때문이다
    (다른 구현이 첫 번째 구현의 일부 기능을 지원하지 않을 수도 있다).
    이는 Ingress API의 이식성을 제한한다.
*   *API의 어색함*. 어노테이션은 키-값 문자열이고(Ingress 리소스의 spec과
    같은 구조화된 스키마와는 대조적으로) 리소스 상단에 적용되기 때문에
    (spec의 관련 부분이 아닌), 특히 많은 어노테이션이 Ingress 리소스에
    추가될 때 Ingress API를 사용하기 어색해질 수 있다.

Gateway API는 Ingress 리소스의 모든 기능과 어노테이션을 통해서만
사용할 수 있는 많은 기능을 지원한다. 결과적으로, Gateway API는
Ingress API보다 이식성이 높다. 또한, 다음 섹션에서 보여주듯이,
어노테이션을 전혀 사용할 필요가 없어 어색함 문제를 해결한다.

### 확장성 접근 방식 {#extensibility-approach}

Ingress API에는 두 가지 확장 지점이 있다:

*   Ingress 리소스의 어노테이션 (이전 섹션에서 설명)
*   [리소스 백엔드](https://kubernetes.io/docs/concepts/services-networking/ingress/#resource-backend),
    서비스 이외의 백엔드를 지정할 수 있는 기능

Gateway API는 Ingress API에 비해 기능이 풍부하다. 그러나, 인증과 같은
고급 기능이나 연결 타임아웃 및 헬스 체크와 같은 일반적이지만 데이터
플레인 간에 이식할 수 없는 기능을 구성하려면, Gateway API의 확장에
의존해야 한다.

Gateway API에는 다음과 같은 주요 확장 지점이 있다:

*   *외부 참조.* Gateway API 리소스의 기능(필드)은 해당 기능을 구성하는
    게이트웨이 구현에 특화된 사용자 정의 리소스를 참조할 수 있다.
    예를 들어:
    *   [HTTPRouteFilter]({{< ref "/reference/spec#httproutefilter" >}})는
        `extensionRef` 필드를 통해 외부 리소스를 참조할 수 있어,
        구현별 필터를 구성한다.
    *   [BackendObjectReference]({{< ref "/reference/spec#backendobjectreference" >}})는
        서비스 이외의 리소스를 지원한다.
    *   [SecretObjectReference]({{< ref "/reference/spec#secretobjectreference" >}})는
        Secret 이외의 리소스를 지원한다.
*   *사용자 정의 구현*. 일부 기능의 경우, 지원 방법을 정의하는 것은
    구현에 맡겨진다. 해당 기능은 구현별
    (사용자 정의) [호환성 수준]({{< ref "/overview/concepts/conformance#2-support-levels" >}})에
    해당한다. 예를 들어:
    *   [HTTPPathMatch]({{< ref "/reference/spec#httppathmatch" >}})의
        `RegularExpression` 타입.
*   *정책.* 게이트웨이 구현은 인증과 같은 데이터 플레인 기능을 노출하기
    위해 정책이라는 사용자 정의 리소스를 정의할 수 있다. Gateway API는
    이러한 리소스의 세부 사항을 규정하지 않는다. 그러나, 표준 UX를
    규정한다. 자세한 내용은 [정책 연결 가이드]({{< ref "/reference/policy-attachment" >}})를
    참조하라. 위의 *외부 참조*와 달리, Gateway API 리소스는 정책을
    참조하지 않는다. 대신, 정책이 Gateway API 리소스를 참조해야 한다.

확장 지점에는 Gateway API 리소스의 어노테이션이 포함되지 않는다.
이 접근 방식은 API 구현에 대해 강력히 권장되지 않는다.

## Ingress API 기능을 Gateway API 기능에 매핑

이 섹션에서는 Ingress API 기능을 해당하는 Gateway API 기능에 매핑하며,
세 가지 주요 영역을 다룬다:

*   진입점
*   TLS 종료
*   라우팅 규칙

### 진입점

대략적으로 말하면, 진입점은 데이터 플레인이 외부 클라이언트에 접근
가능한 IP 주소와 포트의 조합이다.

모든 Ingress 리소스에는 두 개의 암묵적 진입점이 있다 -- HTTP와
HTTPS 트래픽용. Ingress 컨트롤러가 이러한 진입점을 제공한다.
일반적으로, 모든 Ingress 리소스가 공유하거나, 각 Ingress 리소스마다
전용 진입점을 받는다.

Gateway API에서는 진입점을
[Gateway]({{< ref "/reference/api-types/gateway" >}}) 리소스에 명시적으로 정의해야 한다.
예를 들어, 데이터 플레인이 포트 80에서 HTTP 트래픽을 처리하도록 하려면,
해당 트래픽을 위한
[리스너]({{< ref "/reference/spec#listener" >}})를
정의해야 한다. 일반적으로, 게이트웨이 구현은 각 Gateway 리소스에 대해
전용 데이터 플레인을 제공한다.

게이트웨이 리소스는 클러스터 운영자와 애플리케이션 관리자가 소유한다.

### TLS 종료

Ingress 리소스는
[TLS 섹션](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls)을 통해
TLS 종료를 지원하며, TLS 인증서와 키는 Secret에 저장된다.

Gateway API에서 TLS 종료는
[게이트웨이 리스너]({{< ref "/reference/spec#listener" >}})의
속성이며, Ingress와 마찬가지로 TLS 인증서와 키도 Secret에 저장된다.

리스너가 Gateway 리소스의 일부이므로, 클러스터 운영자와 애플리케이션
관리자가 TLS 종료를 소유한다.

### 라우팅 규칙

Ingress 리소스의
[경로 기반 라우팅 규칙](https://kubernetes.io/docs/concepts/services-networking/ingress/#path-types)은
[HTTPRoute]({{< ref "/reference/spec#httprouterule" >}})의
[라우팅 규칙]({{< ref "/reference/api-types/httproute" >}})에
직접 매핑된다.

[호스트 헤더 기반 라우팅 규칙](https://kubernetes.io/docs/concepts/services-networking/ingress/#name-based-virtual-hosting)은
HTTPRoute의
[호스트이름]({{< ref "/reference/spec#hostname" >}})에
매핑된다. 그러나, Ingress에서는 각 호스트이름에 별도의 라우팅 규칙이
있는 반면, HTTPRoute에서는 라우팅 규칙이 모든 호스트이름에 적용된다.

> Ingress API는 호스트라는 용어를 사용하고 Gateway API는 호스트이름을
> 사용한다. 이 가이드에서는 Ingress 호스트를 지칭할 때 Gateway API
> 용어를 사용한다.

> HTTPRoute의 `hostnames`는 [게이트웨이 리스너]({{< ref "/reference/spec#listener" >}})의
> `hostname`과 일치해야 한다. 그렇지 않으면, 리스너는 일치하지 않는
> 호스트이름에 대한 라우팅 규칙을 무시한다.
> [HTTPRoute 문서]({{< ref "/reference/spec#httproutespec" >}})를 참조하라.

HTTPRoute는 애플리케이션 개발자가 소유한다.

다음 세 섹션에서는 Ingress 라우팅 규칙의 추가 기능을 매핑한다.

#### 규칙 병합 및 충돌 해결

일반적으로, Ingress 컨트롤러는 모든 Ingress 리소스의 라우팅 규칙을
병합하고(Ingress 리소스별로 데이터 플레인을 프로비저닝하지 않는 한)
규칙 간의 잠재적 충돌을 해결한다. 그러나, 병합과 충돌 해결 모두
Ingress API에 의해 규정되지 않으므로, Ingress 컨트롤러마다
다르게 구현할 수 있다.

반면, Gateway API는 규칙 병합 및 충돌 해결 방법을 규정한다:

*   게이트웨이 구현은 리스너에 연결된 모든 HTTPRoute의 라우팅 규칙을
    병합해야 한다.
*   충돌은
    [API 설계 가이드: 충돌]({{< ref "/guides/api-design#conflicts" >}})에 규정된 대로 처리해야 한다.
    예를 들어, 라우팅 규칙에서 더 구체적인 매치가 덜 구체적인 것보다
    우선한다.

#### 기본 백엔드

Ingress [기본 백엔드](https://kubernetes.io/docs/concepts/services-networking/ingress/#default-backend)는
해당 Ingress 리소스와 관련된 모든 일치하지 않는 HTTP 요청에 응답하는
백엔드를 구성한다. Gateway API에는 직접적인 동등물이 없다:
이러한 라우팅 규칙을 명시적으로 정의해야 한다. 예를 들어, 경로 접두사
`/`를 가진 요청을 기본 백엔드에 해당하는 서비스로 라우팅하는 규칙을
정의한다.

#### 연결할 데이터 플레인 선택

Ingress 리소스는 사용할 Ingress 컨트롤러를 선택하기 위해
[클래스](https://kubernetes.io/docs/concepts/services-networking/ingress/#ingress-class)를
지정해야 한다. HTTPRoute는
[parentRef]({{< ref "/reference/spec#parentreference" >}})를 통해
연결할 게이트웨이(또는 게이트웨이들)를 지정해야 한다.

### 구현별 Ingress 기능 (어노테이션)

Ingress 어노테이션은 구현별 기능을 구성한다. 따라서, 이를
Gateway API로 변환하는 것은 Ingress 컨트롤러와 게이트웨이
구현에 따라 달라진다.

다행히, 어노테이션을 통해 지원되는 기능 중 일부는 이제 Gateway API
(HTTPRoute)의 일부이다, 주로:

*   요청 리다이렉트 (TLS 리다이렉트 포함)
*   요청/응답 조작
*   트래픽 분할
*   헤더, 쿼리 파라미터, 또는 메서드 기반 라우팅

그러나, 나머지 기능은 대부분 구현별로 남아 있다. 이를 변환하려면,
게이트웨이 구현 문서를 참조하여 어떤
[확장 지점](#extensibility-approach)을 사용할지 확인하라.

## 예제

이 섹션에서는 Ingress 리소스를 Gateway API 리소스로 변환하는 예제를
보여준다.

### 가정

이 예제에는 다음과 같은 가정이 포함된다:

*   모든 리소스가 같은 네임스페이스에 속한다.
*   Ingress 컨트롤러:
    *   클러스터에 해당 IngressClass 리소스 `prod`가 있다.
    *   `example-ingress-controller.example.org/tls-redirect` 어노테이션을 통해
        TLS 리다이렉트 기능을 지원한다.
*   게이트웨이 구현에는 클러스터에 해당 GatewayClass 리소스 `prod`가 있다.

또한, 참조된 Secret과 서비스의 내용 및 IngressClass와 GatewayClass는
간결함을 위해 생략된다.

### Ingress 리소스

아래 Ingress는 다음 구성을 정의한다:

*   `example-ingress-controller.example.org/tls-redirect` 어노테이션을 사용하여
    `foo.example.com` 및 `bar.example.com` 호스트이름에 대한 모든 HTTP 요청에
    TLS 리다이렉트를 구성한다.
*   `example-com` Secret의 TLS 인증서와 키를 사용하여
    `foo.example.com` 및 `bar.example.com` 호스트이름에 대해 TLS를 종료한다.
*   URI 접두사 `/orders`를 가진 `foo.example.com` 호스트이름에 대한
    HTTPS 요청을 `foo-orders-app` 서비스로 라우팅한다.
*   다른 모든 접두사를 가진 `foo.example.com` 호스트이름에 대한
    HTTPS 요청을 `foo-app` 서비스로 라우팅한다.
*   모든 URI를 가진 `bar.example.com` 호스트이름에 대한
    HTTPS 요청을 `bar-app` 서비스로 라우팅한다.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example-ingress
  annotations:
    some-ingress-controller.example.org/tls-redirect: "True"
spec:
  ingressClassName: prod
  tls:
  - hosts:
    - foo.example.com
    - bar.example.com
    secretName: example-com
  rules:
  - host: foo.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: foo-app
            port:
              number: 80
      - path: /orders
        pathType: Prefix
        backend:
          service:
            name: foo-orders-app
            port:
              number: 80
  - host: bar.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: bar-app
            port:
              number: 80
```

다음 세 섹션에서는 이 Ingress를 Gateway API 리소스로 변환한다.

### 변환 단계 1 - 게이트웨이 정의

다음 게이트웨이 리소스는:

*   GatewayClass `prod`에 속한다.
*   로드 밸런싱 인프라를 프로비저닝한다 (이는 게이트웨이 구현에 따라
    달라진다).
*   Ingress 리소스에 암묵적으로 포함된 HTTP 및 HTTPS 리스너(진입점)를
    구성한다:
    *   포트 `80`의 HTTP 리스너 `http`
    *   Ingress에서 사용한 것과 동일한 `example-com` Secret에 저장된 인증서와
        키를 사용한 포트 `443`의 HTTPS 리스너 `https` (TLS 종료 포함)

또한, 두 리스너 모두 같은 네임스페이스의 모든 HTTPRoute를 허용하며
(기본 설정), HTTPRoute 호스트이름을 `example.com` 하위 도메인으로
제한한다 (`foo.example.com`과 같은 호스트이름은 허용하지만
`foo.kubernetes.io`는 허용하지 않음).

{{< include file="examples/standard/simple-http-https/gateway.yaml" >}}

### 변환 단계 2 - HTTPRoute 정의

Ingress는 두 개의 HTTPRoute로 분할된다 -- `foo.example.com`용과
`bar.example.com`용.

{{< include file="examples/standard/simple-http-https/foo-route.yaml" >}}

{{< include file="examples/standard/simple-http-https/bar-route.yaml" >}}

두 HTTPRoute 모두:

*   단계 1의 게이트웨이 리소스의 `https` 리스너에 연결된다.
*   해당 호스트이름에 대한 Ingress 규칙과 동일한 라우팅 규칙을 정의한다.

### 단계 3 - TLS 리다이렉트 구성

다음 HTTPRoute는 Ingress 리소스가 어노테이션을 통해 구성한 TLS 리다이렉트를
구성한다. 아래 HTTPRoute는:

*   게이트웨이의 `http` 리스너에 연결된다.
*   `foo.example.com` 또는 `bar.example.com` 호스트이름에 대한 모든 HTTP 요청에
    TLS 리다이렉트를 발행한다.

{{< include file="examples/standard/simple-http-https/tls-redirect-route.yaml" >}}

## Ingress 자동 변환

[Ingress to Gateway](https://github.com/kubernetes-sigs/ingress2gateway) 프로젝트는
Ingress 리소스를 Gateway API 리소스, 특히 HTTPRoute로 변환하는 데 도움을 준다.
변환 결과는 항상 테스트하고 검증해야 한다.
