---
title: "트래픽 매칭"
weight: 90
description: "Gateway API에서 트래픽이 리스너와 라우트에 매칭되는 방식"
---


## 리스너 선택

**라우트**(Route)가 연결(attach)될 때, 이 과정은 라우트의 `parentRef`에 지정된 **리스너**(Listener)와 해당 리스너에서 어떤 라우트가 연결될 수 있는지 제한하는 설정 간의 협상으로 생각할 수 있다.

라우트의 `parentRef`를 고려할 때, 라우트가 연결될 _수 있는_ 리스너 집합을 **관련**(relevant) 리스너 집합이라고 한다.

여러 리스너를 포함하는 **게이트웨이**(Gateway)의 `parentRef`를 지정하는 라우트는 사실상 해당 게이트웨이의 _모든_ 리스너에 연결을 시도하는 것이며, _모든_ 리스너가 **관련**(relevant) 리스너이다.

게이트웨이의 `parentRef`와 해당 `parentRef`의 `sectionName`을 모두 지정하는 라우트의 경우, 유일한 **관련**(relevant) 리스너는 해당 `sectionName`과 일치하는 `name` 필드를 가진 리스너이다.
일치하는 `name` 필드를 가진 리스너가 없으면, **관련**(relevant) 리스너 집합은 비어 있게 되며, 해당 `parentRef`는 무시된다.

라우트가 연결에 실패하는 다른 경우도 있다는 점에 유의해야 한다. 예를 들어, HTTPRoute는 호스트 이름 교차(Hostname Intersection)를 기반으로 리스너를 매칭할 수도 있으며, 이는 별도의 페이지에서 설명한다.

리스너의 경우, 어떤 라우트가 연결될 수 있는지 제한하는 두 가지 주요 방법이 있다.

* 라우트 그룹 또는 종류 제한 (예: `AllowedRoutes.Kinds`에 "HTTPRoutes"를 지정하여 HTTPRoute만 연결을 허용할 수 있다. 또는 자체 라우트를 만들어 해당 라우트만 허용할 수도 있다).
* 라우트가 연결할 수 있는 네임스페이스 제한 (예: `AllowedRoutes.Namespaces` 필드에 "All", "Same" 또는 "Selector"를 지정한다).

라우트의 **관련**(relevant) 리스너는 라우트 연결이 성공하기 위해 이러한 제한 사항에 부합해야 한다.

연결이 _성공_ 하면, 해당 라우트는 리스너 상태의 `attachedRoutes` 필드에 포함된다.
이 필드는 해당 특정 리스너에 성공적으로 연결된 라우트의 총 수를 기록한다.

## 트래픽 매칭

게이트웨이를 통과하는 트래픽이 반드시 단일 리스너에만 매칭되어야 한다는 요구사항은 라우트 기반 구성에도 확장 적용된다.

대부분의 사용 사례에서, 트래픽이 다양한 객체에 매칭되는 방식은 다음과 같다.

* 트래픽은 IP 주소로 유입되어 게이트웨이를 선택한다 (주소를 가지는 것은 게이트웨이뿐이기 때문이다).
* 해당 IP 주소에서 트래픽은 포트로 향하며, 이를 통해 하나 이상의 리스너를 선택한다. 리스너가 둘 이상인 경우, 추가 정보도 사용될 수 있다.
    * `hostname`은 HTTP, HTTPS, TLS 및 관련 라우트에 사용된다
* 트래픽이 통과할 단일 후보 라우트가 선택된다. (라우트 간 라우트 _매치_ 충돌이 발생하는 경우, 가장 오래된 라우트의 매치가 선택된다).

이러한 요구사항의 중요한 결론은 **트래픽이 특정 라우트에 지정된 어떤 트래픽과도 일치하지 않는 경우, 재라우팅을 위해 역시 매칭되는 다른 리스너를 선택할 수 없다**는 것이다.

예를 들어, 게이트웨이에 두 개의 `HTTP` 리스너가 있고, 하나는 `specific.example.com`용이고 다른 하나는 `*.example.com`용인 경우, `specific.example.com`에 대한 트래픽은 반드시 `specific.example.com` 리스너에 연결된 HTTPRoute에 의해 캡처되어야 하며, 그렇지 않으면 404를 반환한다.

따라서 게이트웨이와 라우트가 다음과 같은 경우:
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: example-com
  namespace: default
spec:
  listeners:
    - name: specific
      hostname: specific.example.com
      protocol: HTTP
      port: 80
    - name: wildcard
      hostname: *.example.com
      protocol: HTTP
      port: 80
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: specific
  namespace: default
spec:
  parentRefs:
    - name: example-com
  rules:
    - matches:
      - path:
          type: Exact
          value: /specific
      backendRefs:
      - name: specific
        port: 8080
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: wildcard
  namespace: default
spec:
  parentRefs:
    - name: example-com
  rules:
    - matches:
      - path:
          type: prefix
          value: /
      backendRefs:
      - name: prefix
        port: 8080


```

`specific` 라우트는 URL `http://specific.example.com/specific`으로 향하는 트래픽에_만_ 매칭된다.

`http://specific.example.com/otherpath`와 같은 다른 요청은 와일드카드 리스너에 매칭될 수 _있다고_ 이해될 수 있음에도 불구하고 (`*.wildcard.com`이 `specific.example.com`과도 일치하고, `/otherpath`가 `wildcard` HTTPRoute의 `/` 접두사 경로와 일치하므로) 404를 반환한다.
이는 리스너의 단일 매칭 속성 때문이며, 트래픽은 _다른 리스너_ 에 연결된 HTTPRoute와 _추가로_ 매칭될 수 없다.
