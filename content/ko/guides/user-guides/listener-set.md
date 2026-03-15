---
title: "ListenerSet"
weight: 17
description: "Delegated listener management with ListenerSets"
---

# ListenerSet

ListenerSet을 사용하면 팀이 모든 것을 64개의 리스너 제한이 있는 하나의 거대한 Gateway 객체에 넣는 대신, 별도의 리소스에서 포트, 호스트명, TLS 인증서를 정의할 수 있다. 이를 통해 대규모 멀티 테넌트 환경에 적합한 위임 관리 모델을 구현할 수 있다.

ListenerSet을 사용하면 다음과 같은 이점을 얻을 수 있다:

- *멀티테넌시*: 동일한 Gateway와 로드 밸런싱 인프라를 공유하면서 각 팀이 자체 ListenerSet을 생성할 수 있다.

- *대규모 배포*: ListenerSet을 사용하면 Gateway에 64개 이상의 리스너를 연결할 수 있다. 또한 팀은 중복을 피하기 위해 동일한 ListenerSet 설정을 공유할 수 있다.

- *Gateway당 더 많은 리스너를 위한 인증서*: 이제 Gateway당 64개 이상의 리스너를 가질 수 있으므로, 단일 Gateway가 자체 인증서를 가진 더 많은 백엔드로 보안 트래픽을 전달할 수 있다. 이 접근 방식은 Istio Ambient Mesh나 Knative와 같이 서비스 수준의 인증서가 필요한 프로젝트에 적합하다.

다음 다이어그램은 멀티 테넌트 환경에서 ListenerSet이 라우트 설정을 분산하는 데 어떻게 도움이 되는지 보여준다:

- Team 1과 Team 2는 각각의 네임스페이스 내에서 자체 Service와 HTTPRoute 리소스를 관리한다.

- 각 HTTPRoute는 네임스페이스 로컬 ListenerSet을 참조한다. 이를 통해 각 팀은 프로토콜, 포트, TLS 인증서 설정 등 라우트가 노출되는 방식을 제어한다.

- 양 팀의 ListenerSet은 별도의 네임스페이스에 있는 공통 Gateway를 공유한다. 별도의 Gateway 팀이 중앙 집중식 인프라를 설정 및 관리하거나 필요에 따라 정책을 적용할 수 있다.

{{< mermaid >}}
flowchart TD

  subgraph team1 namespace
    SVC1[Services]
    HR1[HTTPRoutes]
    LS1[ListenerSet]
  end

  subgraph team2 namespace
    SVC2[Services]
    HR2[HTTPRoutes]
    LS2[ListenerSet]
  end

  subgraph shared namespace
    GW[Gateway]
  end

  HR1 -- "parentRef" --> LS1
  LS1 -- "parentRef" --> GW
  HR1 -- "backendRef" --> SVC1

  HR2 -- "parentRef" --> LS2
  LS2 -- "parentRef" --> GW
  HR2 -- "backendRef" --> SVC2
{{< /mermaid >}}

## ListenerSet 사용하기
### Gateway 설정

기본적으로 Gateway는 ListenerSet의 연결을 허용하지 않는다. 사용자는 Gateway 스펙에 `allowedListeners` 필드를 추가하여 ListenerSet을 허용하도록 Gateway를 설정할 수 있다. 이 필드는 리스너를 추가할 수 있는 네임스페이스를 정의한다.

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

### ListenerSet 설정
ListenerSet은 `parentRef`를 통해 상위 Gateway를 참조한다:

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

HTTPRoute(또는 다른 라우트 타입)를 ListenerSet에 연결하려면, `parentRefs`를 통해 ListenerSet을 상위 리소스로 참조해야 한다.

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

## 리스너 충돌

여러 ListenerSet이 하나의 Gateway에 연결될 수 있으므로, 충돌을 관리하는 규칙이 있다. 리스너는 Port, Protocol, 그리고 (Protocol에 따라) Hostname의 조합을 기반으로 모든 리스너 그룹(단일 ListenerSet 객체 내 또는 Gateway에 연결된 리스너 그룹) 내에서 **고유**(distinct)할 수 있다.

리스너가 고유하지 않은 경우, 충돌을 해결하고 어떤 리스너가 적용될지 선택하기 위해 리스너 우선순위 규칙이 사용된다. 이를 경쟁으로 생각하면 가장 쉬우며, "승자"가 적용되는 리스너이다. 규칙은 다음과 같으며, 동점인 경우 다음 규칙으로 넘어간다.

* 상위 Gateway의 리스너가 다른 모든 리스너보다 우선한다.
* 생성 시간이 가장 빠른 ListenerSet이 우선한다.
* 알파벳순으로 가장 앞에 오는 ListenerSet이 우선한다.

승리한 ListenerSet은 `Accepted: true`로 표시되고, 패배한 ListenerSet은 `Accepted: false`, `Conflicted: true`로 표시된다.

Gateway API의 다른 모든 충돌 해결 규칙과 마찬가지로, 이는 트래픽 안정성을 제공하기 위한 것이다. 따라서 새로운 충돌하는 ListenerSet을 추가해도 기존 설정을 대체하지 않는다.

## 예시

다음 예시는 HTTP 리스너가 있는 Gateway와 고유한 호스트명 및 인증서를 가진 두 개의 하위 HTTPS ListenerSet을 보여준다. `belongs: shared-gateway` 레이블이 있는 네임스페이스의 ListenerSet만 허용된다:

{{< include file="examples/standard/listenerset/listenerset.yaml" >}}
