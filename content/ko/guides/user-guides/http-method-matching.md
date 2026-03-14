---
title: "HTTP 메서드 매칭"
linkTitle: "HTTP method matching"
weight: 7
description: "Matching HTTP requests based on the HTTP method"
---

# HTTP 메서드 매칭

{{< collapsible-alert color="info" title="확장 지원 기능: HTTPRouteMethodMatching" open="true" >}}
이 기능은 확장 지원의 일부이다. 지원 수준에 대한 자세한 내용은 [적합성 가이드]({{< ref "/overview/concepts/conformance" >}})를 참조한다.
{{< /collapsible-alert >}}

[HTTPRoute 리소스]({{< ref "/reference/api-types/httproute" >}})는 HTTP 메서드를 기반으로
요청을 매칭하는 데 사용할 수 있다. 이 가이드에서는 이 기능을 사용하는 방법을
설명한다.

## HTTP 메서드 기반 요청 매칭

다음 `HTTPRoute`는 요청의 HTTP 메서드에 따라 트래픽을 두 개의 백엔드로
분할한다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: method-matching
  namespace: gateway-conformance-infra
spec:
  parentRefs:
  - name: same-namespace
  rules:
  - matches:
    - method: POST
    backendRefs:
    - name: infra-backend-v1
      port: 8080
  - matches:
    - method: GET
    backendRefs:
    - name: infra-backend-v2
      port: 8080
```

- `/`에 대한 `POST` 요청은 `infra-backend-v1`로 라우팅된다.
- `/`에 대한 `GET` 요청은 `infra-backend-v2`로 라우팅된다.

## 다른 매칭 유형과의 조합

메서드 매칭은 경로 매칭이나 헤더 매칭과 같은 다른 매칭 유형과 결합할 수 있다.
다음 규칙은 이를 보여준다.

```yaml
  # 코어 매칭 유형과의 조합
  - matches:
    - path:
        type: PathPrefix
        value: /path1
      method: GET
    backendRefs:
    - name: infra-backend-v1
      port: 8080
  - matches:
    - headers:
      - name: version
        value: one
      method: PUT
    backendRefs:
    - name: infra-backend-v2
      port: 8080
  - matches:
    - path:
        type: PathPrefix
        value: /path2
      headers:
      - name: version
        value: two
      method: POST
    backendRefs:
    - name: infra-backend-v3
      port: 8080
```

## OR(논리합) 매칭

하나의 규칙에 여러 `matches`가 있는 경우, 요청이 그 중 하나라도 만족하면
라우팅된다. 다음 규칙은 아래 조건을 만족하면 트래픽을 `infra-backend-v1`로
라우팅한다.

- 요청이 `/path3`에 대한 `PATCH`인 경우
- OR(또는) 요청이 `version: three` 헤더를 가진 `/path4`에 대한 `DELETE`인 경우

```yaml
  # (조건1 AND 조건2) OR (조건3 AND 조건4 AND 조건5) 형태의 매칭
  - matches:
    - path:
        type: PathPrefix
        value: /path3
      method: PATCH
    - path:
        type: PathPrefix
        value: /path4
      headers:
      - name: version
        value: three
      method: DELETE
    backendRefs:
    - name: infra-backend-v1
      port: 8080
```
