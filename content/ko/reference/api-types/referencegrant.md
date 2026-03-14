---
title: "ReferenceGrant"
weight: 60
description: "Enabling cross-namespace references in Gateway API"
---

{{< channel-version channel="standard" version="v0.6.0" >}}

`ReferenceGrant` 리소스는 베타이며 `v0.6.0` 부터
표준 채널의 일부이다. 릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하라.
{{< /channel-version >}}

{{< note >}}
이 리소스는 원래 "ReferencePolicy"라는 이름이었다. 정책 연결과의
혼동을 피하기 위해 "ReferenceGrant"로 이름이 변경되었다.
{{< /note >}}

**ReferenceGrant**(참조 허가)는 Gateway API 내에서 교차 네임스페이스 참조를
활성화하는 데 사용할 수 있다. 특히 라우트가 다른 네임스페이스의 백엔드로
트래픽을 전달하거나, 게이트웨이가 다른 네임스페이스의 Secret을 참조할 수 있다.

![Reference Grant](/images/referencegrant-simple.svg)
<!-- Source: https://docs.google.com/presentation/d/11HEYCgFi-aya7FS91JvAfllHiIlvfgcp7qpi_Azjk4E/edit#slide=id.g13c18e3a7ab_0_171 -->

과거에 네임스페이스 경계를 넘어 트래픽을 전달하는 것이 필요한 기능이었지만,
ReferenceGrant와 같은 보호 장치 없이는
[취약점](https://github.com/kubernetes/kubernetes/issues/103675)이
발생할 수 있었다.

객체가 자신의 네임스페이스 외부에서 참조되는 경우, 객체의 소유자는
해당 참조를 명시적으로 허용하기 위해 ReferenceGrant 리소스를 생성해야 한다.
ReferenceGrant가 없으면 교차 네임스페이스 참조는 유효하지 않다.

`ReferenceGrant`는 주의하여 사용하는 것이 권장되며, 이 리소스의 올바른 사용을
보장하기 위해 클러스터 관리자가 유효성 검사 및 제한을 적용해야 한다.

자세한 내용은 [보안 고려 사항]({{< ref "/overview/concepts/security" >}}#limiting-cross-namespace-references)을
참조하라.

## 구조
기본적으로 ReferenceGrant는 두 개의 목록으로 구성된다. 참조가 올 수 있는
리소스 목록과 참조될 수 있는 리소스 목록이다.

`from` 목록을 사용하면 `to` 목록에 설명된 항목을 참조할 수 있는 리소스의
group, kind, namespace를 지정할 수 있다.

`to` 목록을 사용하면 `from` 목록에 설명된 항목이 참조할 수 있는 리소스의
group과 kind를 지정할 수 있다. `to` 목록에는 네임스페이스가 필요하지 않은데,
이는 ReferenceGrant가 ReferenceGrant와 동일한 네임스페이스에 있는 리소스에
대한 참조만 허용하는 데 사용될 수 있기 때문이다.

## 예시
다음 예시는 네임스페이스 `foo`의 HTTPRoute가 네임스페이스 `bar`의
Service를 참조하는 방법을 보여준다. 이 예시에서 `bar` 네임스페이스의
ReferenceGrant는 `foo` 네임스페이스의 HTTPRoute에서 Service로의 참조를
명시적으로 허용한다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: foo
  namespace: foo
spec:
  rules:
  - matches:
    - path: /bar
    backendRefs:
      - name: bar
        namespace: bar
---
apiVersion: gateway.networking.k8s.io/v1beta1
kind: ReferenceGrant
metadata:
  name: bar
  namespace: bar
spec:
  from:
  - group: gateway.networking.k8s.io
    kind: HTTPRoute
    namespace: foo
  to:
  - group: ""
    kind: Service
```

## API 설계 결정
API는 본질적으로 단순하지만, 몇 가지 주목할 만한 결정 사항이 있다:

1. 각 ReferenceGrant는 단일 From과 To 섹션만 지원한다. 추가적인
   신뢰 관계는 추가 ReferenceGrant 리소스로 모델링해야 한다.
1. 리소스 이름은 ReferenceGrant의 "From" 섹션에서 의도적으로 제외되었는데,
   이는 의미 있는 보호를 거의 제공하지 않기 때문이다. 네임스페이스 내에서
   특정 종류의 리소스에 쓸 수 있는 사용자는 항상 리소스의 이름을 변경하거나
   구조를 변경하여 주어진 허가와 일치시킬 수 있다.
1. "From" 구조체당 단일 네임스페이스만 허용된다. 셀렉터가 더 강력하겠지만,
   불필요하게 안전하지 않은 구성을 조장하게 된다.
1. 이러한 리소스의 효과는 순수하게 추가적이며, 서로 위에 쌓인다.
   이로 인해 서로 충돌하는 것이 불가능하다.

자세한 내용은 특정 ReferenceGrant 필드가 어떻게 해석되는지에 대해
[API 사양]({{< ref "/reference/spec" >}}#referencegrant)을 참조하라.

## 구현 가이드라인
이 API는 런타임 검증에 의존한다. 구현은 이러한 리소스에 대한 변경을
반드시 감시하고(MUST) 각 변경 또는 삭제 후 교차 네임스페이스 참조의
유효성을 재계산해야 한다.

교차 네임스페이스 참조의 상태를 전달할 때, 구현은 참조가 허용되는
ReferenceGrant가 존재하지 않는 한 다른 네임스페이스의 리소스 존재에 대한
정보를 노출해서는 안 된다(MUST NOT). 이는 ReferenceGrant 없이 존재하지 않는
리소스에 대한 교차 네임스페이스 참조가 이루어진 경우, 모든 상태 조건이나
경고 메시지는 이 참조를 허용하는 ReferenceGrant가 존재하지 않는다는 사실에
초점을 맞춰야 한다는 것을 의미한다. 참조된 리소스가 존재하는지 여부에 대한
힌트를 제공해서는 안 된다.

## 예외
교차 네임스페이스 라우트 -> 게이트웨이 바인딩은 핸드셰이크 메커니즘이
게이트웨이 리소스에 내장된 약간 다른 패턴을 따른다. 이 접근 방식에 대한
자세한 정보는 관련 [보안 모델 문서]({{< ref "/overview/concepts/security" >}})를 참조하라.
ReferenceGrant와 개념적으로 유사하지만, 이 구성은 게이트웨이 리스너에
직접 내장되어 있으며, ReferenceGrant로는 불가능한 리스너별
세밀한 구성을 허용한다.

ReferenceGrant를 무시하고 다른 보안 메커니즘을 선호하는 것이 허용될 수 있는
(MAY) 상황이 있다. 이는 NetworkPolicy와 같은 다른 메커니즘이 구현에 의한
교차 네임스페이스 참조를 효과적으로 제한할 수 있는 경우에만 수행할 수 있다(MAY).

이 예외를 선택하는 구현은 ReferenceGrant가 해당 구현에서 존중되지 않는다는
것을 명확히 문서화하고 어떤 대안적 보호 장치가 사용 가능한지
상세히 설명해야 한다(MUST). 이는 API의 인그레스 구현에는 적용되지 않을
가능성이 높으며, 모든 메시 구현에 적용되지는 않는다.

교차 네임스페이스 참조에 관련된 위험의 예시는
[CVE-2021-25740](https://github.com/kubernetes/kubernetes/issues/103675)을
참조하라. 이 API의 구현은 혼동된 대리인(confused deputy) 공격을 피하기 위해
매우 주의해야 한다. ReferenceGrant는 이에 대한 보호 장치를 제공한다.
예외는 동등하게 효과적인 다른 보호 장치가 확실히 갖추어진 구현에서만
만들어야 한다(MUST).

## 적합성 수준
ReferenceGrant 지원은 다음 객체에서 시작되는 교차 네임스페이스 참조에 대한
"CORE" 적합성 수준 요구 사항이다:

- Gateway
- GRPCRoute
- HTTPRoute
- TLSRoute
- TCPRoute
- UDPRoute

즉, 모든 구현은 위의 예외 섹션에 명시된 경우를 제외하고, 게이트웨이 및
모든 핵심 xRoute 타입에서의 교차 네임스페이스 참조에 대해 이 흐름을
반드시 사용해야 한다(MUST).

다른 "ImplementationSpecific" 객체 및 참조도 위의 예외 섹션에 명시된
경우를 제외하고, 교차 네임스페이스 참조에 대해 이 흐름을 반드시
사용해야 한다(MUST).

## 향후 API 그룹 변경 가능성

ReferenceGrant는 Gateway API 및 SIG Network 사용 사례 외부에서도 관심을
받기 시작하고 있다. 이 리소스가 더 중립적인 위치로 이동할 수 있다.
ReferenceGrant API 사용자는 향후 어느 시점에 다른 API 그룹
(`gateway.networking.k8s.io` 대신)으로 전환해야 할 수 있다.
