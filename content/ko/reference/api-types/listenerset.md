---
title: "ListenerSet"
weight: 80
description: "Specifying additional listeners for a Gateway"
---

{{< channel-version channel="standard" version="v1.5.0" >}}

`ListenerSet` 리소스는 GA(정식 출시)되었으며 `v1.5.0` 부터 표준 채널의
일부이다. 릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하자.
{{< /channel-version >}}

**ListenerSet**(리스너셋)은 Gateway에 추가 리스너를 지정하기 위한 Gateway API 타입이다.
포트, 호스트네임, TLS 종료 등의 네트워크 리스너 구성을 중앙 Gateway 리소스에서 분리한다.

## 배경

리스너셋을 사용하면 팀이 중앙의 공유 Gateway에 리스너 그룹을 독립적으로 정의하고 연결할 수 있다.
이를 통해 셀프 서비스 TLS 구성이 가능하고, 분산된 리스너 관리를 통해 멀티테넌시가 개선되며,
단일 Gateway 리소스의 64개 리스너 제한을 넘어 확장할 수 있다.

리스너셋은 다음과 같은 이점을 제공한다:

- *멀티테넌시*: 서로 다른 팀이 동일한 Gateway와 백엔드 로드밸런싱 인프라를 공유하면서
자체 리스너셋을 생성할 수 있다.

- *대규모 배포*: 리스너셋을 사용하면 Gateway에 64개 이상의 리스너를 연결할 수 있다.
또한 팀은 중복을 피하기 위해 동일한 리스너셋 구성을 공유할 수 있다.

- *게이트웨이당 더 많은 리스너를 위한 인증서*: 이제 Gateway당 64개 이상의 리스너를 가질 수 있으므로,
단일 Gateway가 자체 인증서를 가진 더 많은 백엔드로 보안 트래픽을 전달할 수 있다.
이 접근 방식은 Istio Ambient Mesh나 Knative와 같이 서비스 수준 인증서가 필요한 프로젝트에 적합하다.

## 사양

`ListenerSet` 사양은 다음을 정의한다:

*   `ParentRef`- 이 리스너셋이 연결하고자 하는 Gateway를 정의한다.
*   `Listeners`- 호스트네임, 포트, 프로토콜, 종료, TLS 설정 및
    리스너에 연결할 수 있는 라우트를 정의한다.

## 리스너셋 연결
### Gateway 구성

기본적으로 `Gateway`는 `ListenerSet`의 연결을 허용하지 않는다. 사용자는
`spec.allowedListeners`를 통해 `ListenerSet` 연결을 허용하도록 `Gateway`를 구성하여
이 동작을 활성화할 수 있다:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: parent-gateway
spec:
  allowedListeners:
    namespaces:
      from: Same
```

`AllowedListeners` 내의 `namespaces.from` 필드는 다음 네 가지 값을 가질 수 있다:

- `None` (기본값): 외부 ListenerSet의 연결이 허용되지 않는다. Gateway 리소스 내에서
직접 정의된 리스너만 사용된다.

- `Same`: Gateway와 동일한 네임스페이스에 위치한 ListenerSet만 연결할 수 있다.

- `All`: 클러스터의 모든 네임스페이스에서 ListenerSet을 연결할 수 있으며,
Gateway를 가리키는 유효한 parentRef가 있어야 한다.

- `Selector`: 특정 레이블 셀렉터와 일치하는 네임스페이스에 있는 ListenerSet만 허용된다.
이 값을 사용할 때는 selector 필드도 함께 제공해야 한다.

### ListenerSet 구성
리스너셋은 parentRef를 사용하여 특정 Gateway를 가리킨다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: ListenerSet
metadata:
  name: workload-listeners
spec:
  parentRef:
    name: parent-gateway
    kind: Gateway
    group: gateway.networking.k8s.io
```

### 라우트 연결

라우트는 `ListenerSet`을 `parentRef`로 지정할 수 있다. 라우트는 `ParentReference`의
`sectionName` 필드를 사용하여 특정 리스너를 대상으로 지정할 수 있다. 리스너가
대상으로 지정되지 않은 경우(`sectionName`이 설정되지 않은 경우) 라우트는
`ListenerSet`의 모든 리스너에 연결된다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: httproute-example
spec:
  parentRefs:
  - name: workload-listeners
    kind: ListenerSet
    group: gateway.networking.k8s.io
    sectionName: second
```

드문 경우이지만, ListenerSet과 그 부모 Gateway의 별도 리스너에 라우트를 연결하고
싶을 수 있다. 이 경우 구성은 다음과 같다:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: httproute-example
spec:
  parentRefs:
  - name: workload-listeners
    kind: ListenerSet
    group: gateway.networking.k8s.io
    sectionName: second
  - name: parent-gateway
    kind: Gateway
    sectionName: foo
```

## 리스너 충돌

여러 리스너가 동일한 포트, 프로토콜 및/또는 호스트네임을 요청할 때 충돌이 발생한다. 컨트롤러는 다음 우선순위를 사용하여 이를 해결한다:

1. 부모 Gateway 리스너: Gateway 사양에 직접 정의된 리스너가 항상 가장 높은 우선순위를 가진다.

2. ListenerSet 생성 시간: 두 ListenerSet이 충돌하는 경우, 더 오래된 creationTimestamp를 가진 것이 우선한다.

3. ListenerSet 알파벳 순서: 생성 타임스탬프가 동일한 경우, 리소스의 `{namespace}`/`{name}` 알파벳 순서에 따라 우선순위가 부여된다.

가장 높은 우선순위를 가진 리스너는 Accepted 및 Programmed 상태가 된다.
낮은 우선순위의 리스너는 상태에서 `Conflicted: True` 조건으로 표시된다.

{{< note >}}
**부분적인 ListenerSet 수락**

일부 리스너만 충돌하는 경우 ListenerSet이 부분적으로 수락될 수 있다. 유효한 리스너는 계속 트래픽을 라우팅하지만, 충돌하는 리스너는 트래픽을 라우팅하지 않는다.
{{< /note >}}

## 상태 업데이트

다음 세 가지 조건이 모두 충족되면 ListenerSet이 Gateway에 성공적으로 연결된다:

1. Gateway AllowedListeners 구성: 기본적으로 Gateway는 외부 ListenerSet의 연결을 허용하지 않는다. Gateway의 사양에 ListenerSet의 네임스페이스를 선택하는 `allowedListeners` 필드가 있어야 한다.

2. 유효한 부모 참조: ListenerSet은 대상 Gateway를 명시적으로 가리켜야 한다.

3. 리소스 수준 수락: Gateway 컨트롤러가 리소스를 검증하고 "수락"해야 한다(모든 리스너가 유효해야 하는 등).

{{< note >}}
**부분적인 ListenerSet 수락**

개별 리스너 중 하나가 다른 셋과 충돌하더라도 ListenerSet 전체가 Accepted 상태가 될 수 있다. 이 경우 충돌하지 않는 리스너만 데이터 플레인에 "Programmed"된다.
{{< /note >}}

### Gateway 상태
부모 `Gateway` 상태는 성공적으로 연결된 리스너 수를 `.status.attachedListenerSets`에 보고한다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: parent-gateway
...
status:
  attachedListenerSets: 2
```

### ListenerSet 상태
`ListenerSet`에는 최상위 `Accepted` 및 `Programmed` 조건이 있다. 세부 사항은 다음과 같다:

`Accepted: True` 조건은 ListenerSet이 수락된 경우 설정된다.

`Accepted: False` 조건은 부모 Gateway가 ListenerSet을 허용하지 않거나 모든 리스너가 유효하지 않는 등 여러 이유로 설정될 수 있다.

ListenerSet에는 여러 리스너가 포함될 수 있으므로, 각 리스너는 자체 상태 항목을 가지며 Gateway 리스너와 동일한 로직을 따른다.
