---
title: "구현자 가이드"
weight: 5
description: "Guidelines and tips for building a Gateway API implementation"
---

Gateway API 구현체를 만드는 것에 대해 알고 싶었지만 차마 물어보지 못했던 모든 것.

이 문서는 기본 타입의 godoc 필드 내에 적절한 위치가 없는, _Gateway API 구현체를 작성하기_ 위한 팁과 요령을 모아두는 곳이다.

또한 이 API의 구현자가 흔한 실수를 피하는 데 도움이 되는 가이드라인을 기록하기 위한 곳이기도 하다.

이 API를 _사용하는_ 최종 사용자라면, 무언가를 _구축하는_ 것이 아닌 한 크게 관련이 없을 수 있다.

이 문서는 살아있는 문서이며, 빠진 내용이 있다면 PR을 환영한다!

## Gateway API에 대해 기억해야 할 중요한 사항

이 내용의 대부분은 놀라운 것이 아니길 바라지만, 때때로 명확하지 않은 함의가 있으며 여기서 이를 설명하고자 한다.

### Gateway API는 `kubernetes.io` API이다

Gateway API는 `gateway.networking.k8s.io` API 그룹을 사용한다. 이는 쿠버네티스 코어 바이너리에 포함된 API와 마찬가지로, 릴리스가 발생할 때마다 업스트림 쿠버네티스 리뷰어가 API를 검토했다는 것을 의미한다.

### Gateway API는 CRD를 사용하여 배포된다

Gateway API는 [버전 관리 정책][versioning]을 사용하여 버전이 관리되는 CRD 세트로 제공된다.

해당 버전 관리 정책에서 가장 중요한 부분은, _동일해 보이는_ 객체(즉, 동일한 `group`, `version`, `kind`를 가진 객체)가 약간 다른 스키마를 가질 수 있다는 것이다. 변경은 _호환 가능한_ 방식으로 이루어지므로 일반적으로 "그냥 작동"하지만, "그냥 작동"하는 것을 더 안정적으로 만들기 위해 구현체가 취해야 할 조치가 있다. 이에 대한 자세한 내용은 아래에 설명되어 있다.

CRD 기반 배포는 또한 CRD가 설치되지 _않은_ 상태에서 구현체가 Gateway API 객체를 사용(즉, get, list, watch 등)하려고 하면 쿠버네티스 클라이언트 코드가 심각한 오류를 반환할 가능성이 높다는 것을 의미한다. 이를 처리하기 위한 팁도 아래에 자세히 설명되어 있다.

Gateway API 객체에 대한 CRD 정의에는 모두 두 가지 특정 어노테이션이 포함되어 있다:

- `gateway.networking.k8s.io/bundle-version: <semver-release-version>`
- `gateway.networking.k8s.io/channel: <channel-name>`

"번들 버전"과 "채널"(릴리스 채널의 줄임말)의 개념은 [버전 관리][versioning] 문서에서 설명된다.

구현체는 이를 사용하여 클러스터에 설치된 스키마 버전이 무엇인지(있는 경우) 확인할 수 있다.

[versioning]: {{< ref "/overview/concepts/versioning" >}}

### Standard 채널 CRD에 대한 변경은 하위 호환이 된다

Standard 채널 CRD에 대한 계약의 일부는 _API 버전 내의_ 변경이 _호환 가능_해야 한다는 것이다. Experimental 채널에 속하는 CRD는 하위 호환성을 보장하지 않는다는 점에 유의한다.

[Gateway API 버전 관리 정책]({{< ref "/overview/concepts/versioning" >}})은 대체로 업스트림 쿠버네티스 API와 일치하지만, "검증 수정"을 허용한다. 예를 들어, API 사양에서 값이 유효하지 않다고 명시했지만 해당 검증이 이를 다루지 않았다면, 향후 릴리스에서 해당 유효하지 않은 입력을 방지하는 검증을 추가할 수 있다.

이 계약은 또한 구현체가 작성된 버전보다 높은 버전의 API에서도 실패하지 않는다는 것을 의미한다. 쿠버네티스가 저장하는 새로운 스키마는 구현체가 코드에서 사용하는 이전 버전으로 확실히 직렬화할 수 있기 때문이다.

마찬가지로, 구현체가 _더 높은_ 버전으로 작성된 경우, 이해하는 새로운 값은 이전 버전에 존재하지 않으므로 단순히 _사용되지 않는다_.

## 구현 규칙 및 가이드라인

### CRD 관리

Gateway API CRD를 관리하는 방법과 구현체에 CRD 설치를 번들로 제공하는 것이 허용되는 경우에 대한 정보는 [CRD 관리 가이드]({{< ref "/guides/crd-management" >}})를 참조한다.

### 적합성 및 버전 호환성

적합한 Gateway API 구현체란 각 Gateway API 번들 버전 릴리스에 포함된 적합성 테스트를 통과하는 구현체이다.

구현체는 적합하려면 _건너뛴_ 테스트 _없이_ 적합성 테스트 스위트를 통과해야 한다(MUST). 개발 중에는 테스트를 건너뛸 수 있지만, 적합하다고 인정받고자 하는 버전에는 건너뛴 테스트가 없어야 한다(MUST).

Extended 기능은 Extended 상태에 대한 계약에 따라 비활성화할 수 있다.

Gateway API 적합성은 버전별로 다르다. 버전 N에 대해 적합성을 통과한 구현체가 변경 없이 버전 N+1에 대해 적합성을 통과하지 못할 수 있다.

구현체는 테스트 세부 정보가 포함된 적합성 테스트 스위트 보고서를 Gateway API GitHub 저장소에 제출하는 것이 좋다(SHOULD).

적합성 테스트 스위트 출력에는 지원되는 Gateway API 버전이 포함된다.

#### 버전 호환성

v1.0이 릴리스되면, **Gateway**(게이트웨이)와 **GatewayClass**(게이트웨이 클래스)를 지원하는 구현체는 새로운 Condition인 `SupportedVersion`을 설정해야 한다(MUST). `status: true`는 설치된 CRD 버전이 지원됨을 의미하고, `status: false`는 지원되지 않음을 의미한다.

### 표준 상태 필드 및 Condition

Gateway API에는 많은 리소스가 있지만, 이를 설계할 때 Condition 타입과 `status.conditions` 필드를 사용하여 객체 간에 가능한 한 일관된 상태 경험을 유지하기 위해 노력했다.

대부분의 리소스에는 `status.conditions` 필드가 있지만, 일부는 `conditions` 필드를 _포함하는_ 네임스페이스가 지정된 필드도 가지고 있다.

후자의 경우, 게이트웨이의 `status.listeners`와 라우트의 `status.parents` 필드가 슬라이스의 각 항목이 일부 구성 하위 집합과 관련된 Condition을 식별하는 예이다.

게이트웨이의 경우, _리스너_별로 Condition을 허용하기 위한 것이고, 라우트의 경우, _구현체_별로 Condition을 허용하기 위한 것이다(라우트 객체가 여러 게이트웨이에서 사용될 수 있고, 해당 게이트웨이가 서로 다른 구현체에 의해 조정될 수 있기 때문이다).

이 모든 경우에서 유사한 의미를 가진 비교적 일반적인 Condition 타입이 있다:

- `Accepted` - 리소스 또는 그 일부에 구현체가 제어하는 기본 데이터 플레인에서 일부 구성을 생성할 수 있는 허용 가능한 설정이 포함되어 있다. 이는 _전체_ 구성이 유효하다는 것을 의미하는 것이 아니라, 일부 효과를 생성하기에 _충분한_ 구성이 유효하다는 것을 의미한다.
- `Programmed` - 이는 `Accepted` 이후의 후속 단계를 나타내며, 리소스 또는 그 일부가 Accepted되고 기본 데이터 플레인에 프로그래밍된 상태이다. 사용자는 _가까운 미래의 어느 시점에_ 트래픽이 흐를 준비가 된 구성을 기대해야 한다. 이 Condition은 설정될 _때_ 데이터 플레인이 준비되었다는 것을 말하는 것이 _아니라_, 모든 것이 유효하고 _곧 준비될 것이라는_ 의미이다. "곧"은 구현체에 따라 다른 의미를 가질 수 있다.
- `ResolvedRefs` - 이 Condition은 리소스 또는 그 일부의 모든 참조가 유효하고, 존재하면서 해당 참조를 허용하는 객체를 가리키고 있음을 나타낸다. 이 Condition이 `status: false`로 설정된 경우, 리소스 또는 그 일부의 _최소 하나의_ 참조가 어떤 이유로 유효하지 않으며, `message` 필드가 어떤 것이 유효하지 않은지 나타내야 한다.

구현자는 각 리소스 또는 그 일부에 대한 이러한 Condition의 정확한 세부 사항을 확인하기 위해 각 타입의 godoc을 확인해야 한다.

또한, 업스트림 `Conditions` 구조체에는 선택적 `observedGeneration` 필드가 포함되어 있다 - 구현체는 이 필드를 사용하고(MUST) 상태가 생성되는 시점에 객체의 `metadata.generation` 필드로 설정해야 한다. 이를 통해 API 사용자는 상태가 객체의 현재 버전과 관련이 있는지 판단할 수 있다.

## TLS

TLS는 Gateway API에서 큰 주제이며, 기능 세트가 계속 확장되고 있다. 사용자 관점에서 이 주제를 더 깊이 다루는 [TLS 가이드]({{< ref "/guides/user-guides/tls" >}})가 있지만, 이 섹션에서는 구현자 관점에서의 일부 공백을 메우고자 한다.

### 리스너 격리 {#listener-isolation}
게이트웨이 내에서 TLS 구성은 현재 리스너에만 독점적으로 연결되어 있다. 이 접근 방식을 관리 가능하게 만들기 위해, 모든 구현체가 아래에 정의된 완전한 "**Listener**(리스너) 격리"를 제공하는 목표를 향해 노력하도록 권장하고 있다:

요청은 최대 하나의 리스너와 매칭되어야 한다(SHOULD). 예를 들어, "foo.example.com"과 "*.example.com"에 대해 리스너가 정의된 경우, "foo.example.com"에 대한 요청은 "foo.example.com" 리스너에 연결된 라우트만 사용하여 라우팅되어야 한다(SHOULD)("*.example.com" 리스너가 아님).

리스너 격리를 지원하지 않는 구현체는 이를 명확하게 문서화해야 한다(MUST). 향후, 이 기능에 대한 지원을 주장하는 구현체 간에 이 동작이 일관되도록 HTTPS 리스너 격리 적합성 테스트를 추가할 계획이다.
[#2803](https://github.com/kubernetes-sigs/gateway-api/issues/2803)에서 이러한 테스트에 대한 최신 업데이트를 확인할 수 있다.

### 간접 구성
TLS 인증서가 게이트웨이 소유자에 의해 직접 관리되지 않을 수 있는 다양한 사례가 있다. 이것이 완전한 목록은 아니지만, Gateway API로 TLS 인증서를 관리하는 데 사용될 것으로 예상되는 몇 가지 접근 방식을 문서화한다:

#### 1. 다른 곳의 인증서
일부 제공자는 쿠버네티스 외부에서 TLS 인증서를 구성하고 호스팅하는 기능을 제공한다. 해당 외부 제공자에 연결할 수 있는 구현체는 게이트웨이 리스너의 TLS 옵션을 통해 해당 기능을 노출할 수 있다. 예를 들어:

```
  listeners:
  - name: https
    protocol: HTTPS
    port: 443
    tls:
      mode: Terminate
      options:
        vendor.example.com/certificate-name: store-example-com
```

이 예에서 `store-example-com` 이름은 외부 `vendor.example.com` TLS 인증서 제공자가 저장한 인증서의 이름을 참조한다.

#### 2. 나중에 채워지는 자동 생성 TLS 인증서
많은 사용자는 TLS 인증서가 자동으로 생성되기를 선호한다. 이에 대한 잠재적인 구현 중 하나는 게이트웨이와 HTTPRoute를 감시하고, TLS 인증서를 생성하여 게이트웨이에 첨부하는 컨트롤러를 포함한다. 구현 세부 사항에 따라, 게이트웨이 소유자는 이 기능에 명시적으로 옵트인하기 위해 게이트웨이 또는 리스너 수준에서 무언가를 구성해야 할 수 있다. 예를 들어, 누군가가 이 패턴을 따라 인증서를 생성하는 `acme-cert-generator`를 만들었다고 가정해 보자. 해당 생성기는 `tls.options`에 `acme.io/cert-generator`가 설정되었거나 전체 게이트웨이에 유사한 어노테이션이 설정된 게이트웨이 리스너에서만 인증서를 생성하고 채우도록 선택할 수 있다.

이것은 실제로 [현재 Cert Manager가 작동하는 방식](https://cert-manager.io/docs/usage/gateway/)과 상당히 유사하지만, 그 방식은 게이트웨이 소유자가 이후에 채울 쿠버네티스 Secret을 참조하도록 요구한다. 이 특정 접근 방식은 Gateway API v1.1까지 TLS CertificateRef를 지정해야 했기 때문에 필요했다.

Gateway API v1.1에서 검증이 완화되면서, TLS 인증서를 생성 시 미지정 상태로 남길 수 있어 생성된 TLS 인증서를 다룰 때 구성이 줄어든다.

#### 3. 다른 페르소나에 의해 지정되는 인증서
일부 조직에서는 애플리케이션 개발자가 TLS 인증서를 관리하는 역할을 담당한다(이 역할과 다른 역할에 대한 자세한 내용은 [역할 및 페르소나]({{< ref "/overview/concepts/roles-and-personas" >}})를 참조한다).

이 사용 사례를 가능하게 하려면 새로운 컨트롤러와 CRD가 생성되어야 한다. 이 CRD는 호스트명을 사용자가 제공한 인증서에 연결하고, 컨트롤러는 해당 CRD에 지정된 인증서를 해당 호스트명과 일치하는 게이트웨이 리스너에 채운다. 이 또한 리스너 또는 게이트웨이 수준의 동작 옵트인이 도움이 될 것이다.

### TLS 확장에 대한 전반적인 가이드라인
Gateway API 위에 TLS 확장을 구축할 때 다음 가이드라인을 따르는 것이 중요하다:

1. TLS 옵션이나 어노테이션에 구현체에 고유한 도메인 접두사가 있는 이름을 사용한다.
   (예를 들어, `certificate-name` 대신 `example.com/certificate-name`을 사용한다).
2. 옵션이나 어노테이션 값에 인증서와 같은 민감한 정보를 인코딩하지 않는다.
   대신 이해하기 쉬운 간결한 이름을 통한 참조를 선호한다. 이러한 값은 기술적으로
   최대 253자까지 가능하지만, 전반적인 가독성과 UX를 유지하기 위해 50자 미만으로
   유지하는 것을 강력히 권장한다.
3. 이러한 확장을 활성화하기 위해 Gateway API v1.1+에서는 더 이상 게이트웨이
   리스너에 TLS 구성을 지정하도록 요구하지 않는다. 게이트웨이 리스너에 충분한
   TLS 구성이 지정되지 않은 경우, 구현체는 해당 리스너에 대해 `Programmed`
   condition을 `InvalidTLSConfig` 이유와 함께 `False`로 설정해야 한다(MUST).
4. 지원하기로 선택한 확장에 관계없이, 모든 구현체에서 이식 가능하도록 의도된
   핵심 TLS 구성을 지원하는 것이 중요하다. 확장은 이 API에서 자리가 있지만,
   모든 구현체는 여전히 API의 핵심 기능을 지원해야 한다(MUST).

## 리소스 세부 사항

현재 사용 가능한 각 적합성 프로필에 대해, 구현체가 조정하도록 기대되는 리소스 세트가 있다.

다음 섹션에서는 각 Gateway API 객체를 살펴보고 예상되는 동작을 설명한다.

### GatewayClass

GatewayClass에는 하나의 주요 `spec` 필드인 `controllerName`이 있다. 각 구현체는 도메인 접두사가 있는 문자열 값(예: `example.com/example-ingress`)을 `controllerName`으로 주장해야 한다.

구현체는 _모든_ GatewayClass를 감시하고(MUST), 일치하는 `controllerName`을 가진 GatewayClass를 조정해야 한다. 구현체는 일치하는 `controllerName`을 가진 GatewayClass 세트에서 최소 하나의 호환 가능한 GatewayClass를 선택하고, 각각에 `Accepted` Condition을 `status: true`로 설정하여 해당 GatewayClass의 처리를 수락한다는 것을 나타내야 한다. 일치하는 `controllerName`을 가지지만 Accepted되지 _않은_ GatewayClass에는 `Accepted` Condition을 `status: false`로 설정해야 한다.

구현체는 하나의 GatewayClass만 조정할 수 있는 경우 허용 가능한 GatewayClass 풀에서 하나만 선택할 수 있으며(MAY), 여러 GatewayClass를 조정할 수 있는 경우 원하는 만큼 선택할 수도 있다.

GatewayClass에 호환되지 않는 요소가 있는 경우(작성 시점에서 이에 대한 유일한 가능한 이유는 구현체가 지원하지 않는 `paramsRef` 객체에 대한 포인터가 있는 것이다), 구현체는 호환되지 않는 GatewayClass를 `Accepted`가 아닌 것으로 표시해야 한다(SHOULD).

### Gateway

게이트웨이 객체는 구현체가 이를 조정하기 위해 `spec.gatewayClassName` 필드에서 존재하고 구현체에 의해 `Accepted`된 GatewayClass를 참조해야 한다(MUST).

조정 범위를 벗어나는 게이트웨이 객체(예를 들어, 참조하는 GatewayClass가 삭제된 경우)는 삭제 프로세스의 일부로 구현체에 의해 상태가 제거될 수 있지만(MAY), 이는 필수가 아니다.

### 라우트

모든 라우트 객체는 몇 가지 속성을 공유한다:

- 구현체가 조정 가능한 것으로 간주하려면 범위 내의 부모에 연결되어야 한다(MUST).
- 구현체는 네임스페이스가 지정된 `parents` 필드를 사용하여 범위 내의 각 라우트에 대한 상태를 관련 Condition으로 업데이트해야 한다(MUST). 세부 사항은 특정 라우트 타입을 참조하되, 일반적으로 `Accepted`, `Programmed`, `ResolvedRefs` Condition이 포함된다.
- 범위를 벗어나는 라우트는 상태가 업데이트되지 않아야 한다(SHOULD NOT). 이러한 업데이트가 새로운 소유자의 상태를 덮어쓸 수 있기 때문이다. `observedGeneration` 필드는 남아 있는 상태가 오래되었음을 나타낸다.

#### HTTPRoute

HTTPRoute는 _암호화되지 않고_ 검사에 사용 가능한 HTTP 트래픽을 라우팅한다. 이는 게이트웨이에서 종료된 HTTPS 트래픽(이후 복호화되므로)을 포함하며, HTTPRoute가 경로, 메서드, 헤더와 같은 HTTP 속성을 라우팅 지시어에서 사용할 수 있게 한다.

#### TLSRoute

TLSRoute는 _트래픽 스트림을 복호화하지 않고_ SNI 헤더를 사용하여 암호화된 TLS 트래픽을 관련 백엔드로 라우팅한다.

#### TCPRoute

TCPRoute는 리스너에 도착하는 TCP 스트림을 주어진 백엔드 중 하나로 라우팅한다.

#### UDPRoute

UDPRoute는 리스너에 도착하는 UDP 패킷을 주어진 백엔드 중 하나로 라우팅한다.

### ReferenceGrant

**ReferenceGrant**(참조 승인)는 한 네임스페이스의 리소스 소유자가 다른 네임스페이스의 Gateway API 객체로부터의 참조를 _선택적으로_ 허용하는 데 사용되는 특수한 리소스이다.

ReferenceGrant는 참조 접근을 부여하는 대상과 동일한 네임스페이스에 생성되며, 다른 네임스페이스, 다른 Kind, 또는 둘 다로부터의 접근을 허용한다.

교차 네임스페이스 참조를 지원하는 구현체는 ReferenceGrant를 감시하고(MUST), 범위 내 Gateway API 객체가 참조하는 객체를 가리키는 모든 ReferenceGrant를 조정해야 한다.
