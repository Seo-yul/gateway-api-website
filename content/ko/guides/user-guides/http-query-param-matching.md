---
title: "HTTP 쿼리 파라미터 매칭"
linkTitle: "HTTP query parameter matching"
weight: 6
description: "Matching HTTP requests based on query parameters"
---

# HTTP 쿼리 파라미터 매칭

{{< collapsible-alert color="info" title="확장 지원 기능: HTTPRouteQueryParamMatching" open="true" >}}
이 기능은 확장 지원(Extended Support)의 일부이다. 지원 수준에 대한 자세한 내용은 [적합성 가이드]({{< ref "/overview/concepts/conformance" >}})를 참고한다.
{{< /collapsible-alert >}}

[HTTPRoute 리소스]({{< ref "/reference/api-types/httproute" >}})를 사용하면 쿼리 파라미터를
기반으로 요청을 매칭할 수 있다. 이 가이드에서는 이 기능의 사용 방법을 설명한다.

## 단일 쿼리 파라미터를 기반으로 요청 매칭

다음 `HTTPRoute`는 `animal` 쿼리 파라미터의 값에 따라 두 백엔드로 트래픽을
분할한다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: query-param-matching
  namespace: gateway-conformance-infra
spec:
  parentRefs:
  - name: same-namespace
  rules:
  - matches:
    - queryParams:
      - name: animal
        value: whale
    backendRefs:
    - name: infra-backend-v1
      port: 8080
  - matches:
    - queryParams:
      - name: animal
        value: dolphin
    backendRefs:
    - name: infra-backend-v2
      port: 8080
```

- 쿼리 파라미터 `animal=whale`을 포함한 `/` 경로에 대한 요청은 `infra-backend-v1`로 라우팅된다.
- 쿼리 파라미터 `animal=dolphin`을 포함한 `/` 경로에 대한 요청은 `infra-backend-v2`로 라우팅된다.

## 다중 쿼리 파라미터를 기반으로 요청 매칭

규칙(rule)은 여러 쿼리 파라미터에 대해 매칭할 수도 있다. 다음 규칙은
쿼리 파라미터 `animal=dolphin`과 `color=blue`가 모두 존재하는 경우
트래픽을 `infra-backend-v3`로 라우팅한다.

```yaml
  - matches:
    - queryParams:
      - name: animal
        value: dolphin
      - name: color
        value: blue
    backendRefs:
    - name: infra-backend-v3
      port: 8080
```

## OR 매칭

규칙에 여러 `matches`가 있는 경우, 요청이 그 중 하나라도 만족하면 라우팅된다.
다음 규칙은 아래 조건 중 하나를 만족하면 트래픽을 `infra-backend-v3`로 라우팅한다.

- 쿼리 파라미터 `animal=dolphin`과 `color=blue`가 모두 존재하는 경우
- 또는(OR) 쿼리 파라미터 `ANIMAL=Whale`이 존재하는 경우

```yaml
  - matches:
    - queryParams:
      - name: animal
        value: dolphin
      - name: color
        value: blue
    - queryParams:
      - name: ANIMAL
        value: Whale
    backendRefs:
    - name: infra-backend-v3
      port: 8080
```

## 다른 매칭 타입과의 조합

쿼리 파라미터 매칭은 경로(path) 매칭이나 헤더(header) 매칭과 같은 다른 매칭 타입과
조합하여 사용할 수 있다. 다음 규칙들은 이러한 조합을 보여준다.

```yaml
  - matches:
    - path:
        type: PathPrefix
        value: /path1
      queryParams:
      - name: animal
        value: whale
    backendRefs:
    - name: infra-backend-v1
      port: 8080
  - matches:
    - headers:
      - name: version
        value: one
      queryParams:
      - name: animal
        value: whale
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
      queryParams:
      - name: animal
        value: whale
    backendRefs:
    - name: infra-backend-v3
      port: 8080
```
