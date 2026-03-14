---
title: "HTTP 타임아웃"
linkTitle: "HTTP timeouts"
weight: 8
description: "Configuring HTTP request timeouts with HTTPRoute"
---

# HTTP 타임아웃

{{< collapsible-alert color="info" title="확장 지원 기능: HTTPRouteRequestTimeout" open="true" >}}
이 기능은 확장 지원(Extended Support)의 일부이다. 지원 수준에 대한 자세한 내용은 [적합성 가이드]({{< ref "/overview/concepts/conformance" >}})를 참조한다.
{{< /collapsible-alert >}}

[**HTTPRoute**(HTTP라우트) 리소스]({{< ref "/reference/api-types/httproute" >}})는
HTTP 요청에 대한 타임아웃을 구성하는 데 사용할 수 있다. 이는 장시간 실행되는 요청이
리소스를 소모하는 것을 방지하고 더 나은 사용자 경험을 제공하는 데 유용하다.

`HTTPRouteRule`의 `timeouts` 필드를 사용하여 요청
타임아웃을 지정할 수 있다.

## 요청 타임아웃 설정

다음 HTTPRoute는 `/request-timeout` 경로 접두사를 가진 모든
요청에 대해 500밀리초의 요청 타임아웃을 설정한다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: request-timeout
  namespace: gateway-conformance-infra
spec:
  parentRefs:
  - name: same-namespace
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /request-timeout
    backendRefs:
    - name: infra-backend-v1
      port: 8080
    timeouts:
      request: 500ms
```

이 경로에 대한 요청이 500밀리초보다 오래 걸리면, 게이트웨이는
타임아웃 오류를 반환한다.

## 요청 타임아웃 비활성화

요청 타임아웃을 비활성화하려면, `request` 필드를 `"0s"`로 설정한다.

```yaml
  - matches:
    - path:
        type: PathPrefix
        value: /disable-request-timeout
    backendRefs:
    - name: infra-backend-v1
      port: 8080
    timeouts:
      request: "0s"
```
