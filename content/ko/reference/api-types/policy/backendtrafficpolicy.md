---
title: "BackendTrafficPolicy"
weight: 20
description: "Configuring backend traffic behavior including retry budgets"
aliases:
  - /ko/reference/api-types/backendtrafficpolicy/
---

{{< channel-version channel="experimental" version="v1.3.0" >}}

`BackendTrafficPolicy` 리소스는 `v1.3.0` 부터 실험 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하라.
{{< /channel-version >}}

[BackendTrafficPolicy(백엔드 트래픽 정책)][backendtrafficpolicy]는 유효한 백엔드 리소스를
대상으로 할 때 클라이언트의 동작을 구성하기 위한 Gateway API 타입이다.

## 배경

`BackendTrafficPolicy`는 백엔드(유사한 엔드포인트의 그룹)를 대상으로 요청을 보낼 때
클라이언트의 동작을 구성하는 데 사용된다. 현재 이 정책에는 지정된 백엔드를
대상으로 하는 활성 재시도 수를 "재시도 예산"을 정의하여 동적으로 제한할 수 있는
`RetryConstraint`를 구성하는 기능이 포함되어 있다. 향후 추가 기능이 추가될 수 있다.

### 재시도 제약(Retry Constraint)

클라이언트 측 재시도는 간헐적 장애 기간 동안 요청을 성공적으로 재시도할 수
있도록 하는 데 중요하지만, 과도한 재시도는 시스템을 빠르게 압도하여 재시도 폭풍과
같은 연쇄 장애로 이어질 수 있다. `BackendTrafficPolicy` 내에 `RetryConstraint`를
지정하면 애플리케이션 개발자가 전체 활성 요청 볼륨의 백분율로 활성 클라이언트 측
재시도의 제한을 동적으로 계산할 수 있다. [HTTPRouteRule][httproute] 리소스 내의
retry 스탠자가 요청을 재시도*해야 하는지* 여부와 실패한 요청에 대해 수행할 수 있는
최대 재시도 횟수를 지정하는 반면, 예산 기반 재시도는 대상 백엔드가 클라이언트로부터의
*전체* 재시도 볼륨에 의해 압도되지 않도록 보장하기 위한 안전 장치 역할을 한다.

재시도 예산은 주어진 시간 간격 동안 재시도로 구성된 수신 트래픽의 백분율을
기반으로 계산된다. 재시도 예산 간격 내에서 동일한 원래 요청을 여러 번
재시도하면 각 재시도가 예산 계산에 포함된다. 예산이 초과되면 백엔드에 대한
추가 재시도는 거부되며 클라이언트에 503 응답을 반드시 반환해야 한다(MUST).
재시도 예산 계산의 파라미터는 `RetryConstraint` 내에서 구성할 수 있다.

BackendTrafficPolicy는 직접(Direct)
[PolicyAttachment(정책 연결)]({{< ref "/reference/policy-attachment" >}})이다.
참조된 백엔드를 대상으로 하는 모든 Gateway API 라우트는 구성된
BackendTrafficPolicy를 존중해야 한다. 재시도에 대한 제약을 정의하기 위한
추가 구성은 향후 정의될 수 있다(MAY).


## 사양

[BackendTrafficPolicy][backendtrafficpolicy]의 사양은 다음으로 구성된다:

- [TargetRefs][localpolicytargetreference] - 정책의 대상 API 객체를 정의한다.
  백엔드(Service, ServiceImport 또는 구현별 backendRef와 같은 유사한 엔드포인트의
  그룹)만이 유효한 API 대상 참조이다.
- [RetryConstraint][retryConstraint] - BudgetDetails와 MinRetryRate를 지정하여 재시도 예산이 계산되는 방식에 대한 구성을 정의한다.
- [SessionPersistence][sessionPersistence] - 백엔드에 대한 세션 지속성을 정의한다.

다음 차트는 객체 정의와 관계를 보여준다:
{{< mermaid >}}
flowchart LR
    backendTrafficPolicy[["<b>backendTrafficPolicy</b> <hr><align=left>BackendTrafficPolicySpec: spec<br>PolicyStatus: status</align>"]]
    spec[["<b>spec</b><hr>LocalPolicyTargetReference: targetRefs <br> RetryConstraint: retryConstraint" <br> SessionPersistence: sessionPersistence]]
    status[["<b>status</b><hr>[ ]PolicyAncestorStatus: ancestors"]]
    budget[["<b>budget</b><hr>int: budgetPercent<br>Duration: budgetInterval"]]
    retryConstraint[["<b>retryConstraint</b><hr>BudgetDetails: budget<br>MinRetryRate: RequestRate"]]
    ancestorStatus[["<b>ancestors</b><hr>AncestorRef: parentReference<br>GatewayController: controllerName<br>[]Condition: conditions"]]
    targetRefs[[<b>targetRefs</b><hr>]]
    backendRef["one of:<hr>" Service OR ServiceImport OR an implementation-specific backendRef]
    backendTrafficPolicy -->spec
    backendTrafficPolicy -->status
    spec -->targetRefs & retryConstraint
    retryConstraint -->budget
    status -->ancestorStatus
    targetRefs -->backendRef
{{< /mermaid >}}

### 백엔드 대상 지정

`BackendTrafficPolicy`는 Service, ServiceImport 또는 구현별 backendRef와 같은
하나 이상의 TargetRefs를 통해 백엔드 파드 그룹을 대상으로 한다.
TargetRefs는 Name, Kind, Group을 통해 백엔드를 지정하는 필수 객체 참조이다.
현재 TargetRefs는 서비스의 특정 포트로 범위를 지정할 수 없다.

[backendtrafficpolicy]: /reference/specx/#xbackendtrafficpolicy
[localpolicytargetreference]: /references/spec/#localpolicytargetreference
[retryConstraint]: /reference/specx/#retryconstraint
[sessionPersistence]: /references/spec/#sessionpersistence
[httproute]: /references/spec/#httproute
