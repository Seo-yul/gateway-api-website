---
title: "크로스 네임스페이스 라우팅"
linkTitle: "Cross-Namespace routing"
weight: 10
description: "Routing traffic across namespaces with Gateway API"
---

# 크로스 네임스페이스 라우팅

Gateway API는 크로스 **네임스페이스** 라우팅을 핵심 기능으로 지원한다.
이는 둘 이상의 사용자 또는 팀이 기반 네트워킹 인프라를 공유하면서도 접근 및
장애 도메인을 최소화하기 위해 제어와 구성을 분리해야 할 때 유용하다.

게이트웨이와 라우트는 서로 다른 네임스페이스에 배포할 수 있으며, 라우트는
네임스페이스 경계를 넘어 게이트웨이에 연결할 수 있다. 이를 통해 라우트와 게이트웨이에
대해 네임스페이스별로 서로 다른 사용자 접근 제어를 적용할 수 있으며,
**클러스터** 전체 라우팅 구성의 다양한 부분에 대한 접근과 제어를 효과적으로
분리할 수 있다. 라우트가 네임스페이스 경계를 넘어 게이트웨이에 연결하는 기능은
[_라우트 연결_](#cross-namespace-route-attachment)에 의해 관리된다.
이 가이드에서는 라우트 연결을 살펴보고, 독립적인 팀들이 동일한 게이트웨이를
안전하게 공유하는 방법을 시연한다.

이 가이드에는 동일한 쿠버네티스 클러스터의 `store-ns`와 `site-ns` 네임스페이스에서
운영하는 두 개의 독립 팀, _store_와 _site_가 있다. 이들의 목표와
Gateway API 리소스를 사용하여 이를 달성하는 방법은 다음과 같다:

- site 팀은 _home_과 _login_ 두 개의 애플리케이션을 보유하고 있다. 팀은 접근 및
장애 도메인을 최소화하기 위해 앱 간의 접근과 구성을 가능한 한 격리하고자 한다.
카나리 롤아웃과 같은 라우팅 구성을 격리하기 위해 동일한 게이트웨이에 연결된
별도의 HTTPRoute를 사용하면서도, 동일한 IP 주소, 포트, DNS 도메인 및
TLS 인증서를 공유한다.
- store 팀은 `store-ns` 네임스페이스에 배포한 _store_라는 단일 Service를
보유하고 있으며, 이 역시 동일한 IP 주소와 도메인 뒤에 노출되어야 한다.
- Foobar Corporation은 모든 앱에 대해 `foo.example.com` 도메인을 사용한다.
이는 `infra-ns` 네임스페이스에서 운영하는 중앙 인프라 팀에 의해 관리된다.
- 마지막으로, 보안 팀은 `foo.example.com`의 인증서를 관리한다.
단일 공유 게이트웨이를 통해 이 인증서를 관리함으로써 애플리케이션 팀을
직접 관여시키지 않고도 보안을 중앙에서 제어할 수 있다.

Gateway API 리소스 간의 논리적 관계는 다음과 같다:

![Cross-Namespace routing](/images/cross-namespace-routing.svg)

## 크로스 네임스페이스 라우트 연결 <a name="cross-namespace-route-attachment"></a>

[라우트 연결][attachment]은 라우트가 게이트웨이에 연결되어 라우팅 규칙을
프로그래밍하는 방식을 결정하는 중요한 개념이다. 이는 네임스페이스에 걸쳐 하나
이상의 게이트웨이를 공유하는 라우트가 있을 때 특히 중요하다.
게이트웨이와 라우트 연결은 양방향이다 - 연결은 게이트웨이 소유자와 라우트 소유자
모두가 관계에 동의해야만 성공할 수 있다. 이러한 양방향 관계가 존재하는 이유는
두 가지이다:

- 라우트 소유자는 인지하지 못하는 경로를 통해 애플리케이션이 과도하게 노출되는 것을
원하지 않는다.
- 게이트웨이 소유자는 특정 앱이나 팀이 허가 없이 게이트웨이를 사용하는 것을
원하지 않는다. 예를 들어, 내부 서비스가 인터넷 게이트웨이를 통해
접근 가능해서는 안 된다.

게이트웨이는 _연결 제약 조건_을 지원하며, 이는 어떤 라우트가 연결될 수 있는지를
제한하는 게이트웨이 **리스너**의 필드이다. 게이트웨이는 네임스페이스와
라우트 유형을 연결 제약 조건으로 지원한다. 연결 제약 조건을 충족하지 못하는
라우트는 해당 게이트웨이에 연결할 수 없다. 마찬가지로, 라우트는 라우트의
`parentRef` 필드를 통해 연결하려는 게이트웨이를 명시적으로 참조한다. 이러한 요소들이
함께 인프라 소유자와 애플리케이션 소유자 간의 핸드셰이크를 생성하여
애플리케이션이 게이트웨이를 통해 노출되는 방식을 독립적으로 정의할 수 있게 한다.
이는 관리 오버헤드를 줄이는 효과적인 **정책**이다. 앱 소유자는 앱이
사용할 게이트웨이를 지정할 수 있고, 인프라 소유자는 게이트웨이가 수락하는
네임스페이스와 라우트 유형을 제한할 수 있다.


## 공유 게이트웨이

인프라 팀은 `infra-ns` 네임스페이스에 `shared-gateway` 게이트웨이를 배포한다:

```yaml
{{< include file="examples/standard/cross-namespace-routing/gateway.yaml" >}}
```

위 게이트웨이의 `https` 리스너는 `foo.example.com` 도메인에 대한 트래픽을 매칭한다.
이를 통해 인프라 팀이 도메인의 모든 측면을 관리할 수 있다. 아래의 HTTPRoute는
도메인을 지정할 필요가 없으며, `hostname`이 설정되지 않은 경우 기본적으로 모든
트래픽을 매칭한다. 이렇게 하면 HTTPRoute를 도메인에 구애받지 않고 관리할 수 있어,
애플리케이션 도메인이 고정적이지 않은 경우에 유용하다.

이 게이트웨이는 또한 `infra-ns` 네임스페이스의 `foo-example-com` Secret을 사용하여
HTTPS를 구성한다. 이를 통해 인프라 팀이 앱 소유자를 대신하여 TLS를 중앙에서
관리할 수 있다. `foo-example-com` 인증서는 연결된 라우트로 향하는 모든 트래픽을
종료하며, HTTPRoute 자체에는 별도의 TLS 구성이 필요하지 않다.

이 게이트웨이는 네임스페이스 셀렉터를 사용하여 어떤 HTTPRoute가 연결될 수 있는지를
정의한다. 이를 통해 인프라 팀은 네임스페이스 집합을 허용 목록에 추가하여
누가 또는 어떤 앱이 이 게이트웨이를 사용할 수 있는지를 제한할 수 있다.


```yaml
apiVersion: gateway.networking.k8s.io/v1beta1
kind: Gateway
spec:
  listeners:
  - allowedRoutes:
      namespaces:
        from: Selector
        selector:
          matchLabels:
            shared-gateway-access: "true"
...
```

`shared-gateway-access: "true"` 레이블이 지정된 네임스페이스_만_ 해당 라우트를
`shared-gateway`에 연결할 수 있다. 다음 네임스페이스 집합에서, `no-external-access`
네임스페이스에 `infra-ns/shared-gateway`에 대한 `parentRef`가 있는 HTTPRoute가
존재하더라도, 연결 제약 조건(네임스페이스 레이블)이 충족되지 않았기 때문에
게이트웨이에 의해 무시된다.

```yaml
{{< include file="examples/standard/cross-namespace-routing/0-namespaces.yaml" >}}
```

게이트웨이의 연결 제약 조건은 필수가 아니지만, 여러 팀과 네임스페이스가 있는
클러스터를 운영할 때 모범 사례이다. 클러스터의 모든 앱이 게이트웨이에 연결할
권한이 있는 환경에서는 `listeners[].allowedRoutes` 필드를 구성할 필요가 없으며,
모든 라우트가 게이트웨이를 자유롭게 사용할 수 있다.


## 라우트 연결

store 팀은 `store-ns` 네임스페이스에 `store` Service에 대한 라우트를 배포한다:

```yaml
{{< include file="examples/standard/cross-namespace-routing/store-route.yaml" >}}
```

이 라우트는 `/store` 트래픽을 매칭하여 `store` Service로 전송하는
단순한 라우팅 로직을 가지고 있다.

이제 site 팀은 애플리케이션에 대한 라우트를 배포한다. `site-ns` 네임스페이스에
두 개의 HTTPRoute를 배포한다:

- `home` HTTPRoute는 기본 라우팅 규칙으로 작동하며, 기존 라우팅 규칙에 의해
매칭되지 않은 `foo.example.com/*`에 대한 모든 트래픽을 매칭하여 `home` Service로
전송한다.
- `login` HTTPRoute는 `foo.example.com/login`에 대한 트래픽을
`service/login-v1`과 `service/login-v2`로 라우팅한다. 가중치를 사용하여
이들 간의 트래픽 분배를 세밀하게 제어한다.

이 두 라우트 모두 동일한 게이트웨이 연결 구성을 사용하며, `infra-ns` 네임스페이스의
`gateway/shared-gateway`를 이 라우트들이 연결하려는 유일한 게이트웨이로 지정한다.

```yaml
{{< include file="examples/standard/cross-namespace-routing/site-route.yaml" >}}
```

이 세 개의 라우트가 배포된 후, 모두 `shared-gateway` 게이트웨이에 연결된다.
게이트웨이는 이러한 라우트를 단일 평면 라우팅 규칙 목록으로 병합한다.
[라우팅 우선순위]({{< ref "/reference/spec#httprouterule" >}})는 가장 구체적인 매치에 의해
결정되며, 충돌은 [충돌 해결]({{< ref "/guides/api-design#conflicts" >}})에 따라 처리된다.
이를 통해 독립적인 사용자 간의 라우팅 규칙에 대한 예측 가능하고 결정적인 병합이
제공된다.

크로스 네임스페이스 라우팅 덕분에 Foobar Corporation은 중앙 집중식 제어를
유지하면서도 인프라 소유권을 보다 균등하게 분배할 수 있다. 이를 통해
선언적이고 오픈 소스인 API를 통해 양쪽의 장점을 모두 누릴 수 있다.

[attachment]: {{< ref "/overview/concepts/api-overview#attaching-routes-to-gateways" >}}
