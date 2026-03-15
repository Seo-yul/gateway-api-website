---
title: "CORS(교차 출처 리소스 공유)"
linkTitle: "HTTP CORS"
weight: 9
description: "Configuring CORS policies with HTTPRoute"
---

# CORS(교차 출처 리소스 공유)

{{< collapsible-alert color="info" title="확장 지원 기능: HTTPRouteCORS" open="true" >}}
이 기능은 확장 지원의 일부이며, 구현체가 `HTTPRouteCORS` 기능을 지원해야 한다. 지원 수준에 대한 자세한 내용은 [적합성 가이드]({{< ref "/overview/concepts/conformance" >}})를 참고한다.
{{< /collapsible-alert >}}

{{< channel-version channel="standard" version="v1.5.0" >}}
`HTTPRouteCORS` 기능은 `v1.5.0`부터 표준 채널의 일부이다. 릴리스 채널에 대한
자세한 내용은 [버전 관리 가이드]({{< ref "/reference/api-types/httproute" >}})를 참고한다.
{{< /channel-version >}}

[**HTTPRoute**(HTTP 라우트) 리소스]({{< ref "/reference/api-types/httproute" >}})를 사용하여
CORS(교차 출처 리소스 공유)를 구성할 수 있다. CORS는 한 도메인에서 실행되는
웹 애플리케이션이 다른 도메인의 리소스에 대한 요청을 허용하거나 거부하는
보안 기능이다.

`HTTPRouteRule`의 `CORS` 필터를 사용하여 CORS 정책을 지정할 수 있다.

## 특정 출처의 요청 허용

다음 HTTPRoute는 `https://app.example`에서의 요청을 허용한다:

{{< include file="examples/standard/http-cors/httproute-specific-origin-no-creds.yaml" >}}

특정 출처 목록을 지정하는 대신, 단일 와일드카드(`"*"`)를 지정하여
모든 출처를 허용할 수도 있다:

{{< include file="examples/standard/http-cors/httproute-all-origins-no-creds.yaml" >}}

목록에 반지정(semi-specified) 출처를 사용할 수도 있다.
와일드카드는 스킴 뒤, 호스트명의 시작 부분에 위치한다(예: `https://*.bar.com`):

{{< include file="examples/standard/http-cors/httproute-origins-with-wildcards-no-creds.yaml" >}}

## 자격 증명 허용

`allowCredentials` 필드는 브라우저가 CORS 요청에 자격 증명(쿠키 및 HTTP 인증 등)을
포함할 수 있는지 여부를 지정한다.

다음 규칙은 `https://app.example`에서 자격 증명을 포함한 요청을 허용한다:

{{< include file="examples/standard/http-cors/httproute-credentials-true.yaml" >}}

## 기타 CORS 옵션

`CORS` 필터를 사용하면 다음과 같은 기타 CORS 옵션도 지정할 수 있다:

- `allowMethods`: CORS 요청에 허용되는 HTTP 메서드.
- `allowHeaders`: CORS 요청에 허용되는 HTTP 헤더.
- `exposeHeaders`: 클라이언트에 노출되는 HTTP 헤더.
- `maxAge`: 브라우저가 사전 요청(preflight) 응답을 캐시해야 하는 최대 시간(초 단위).

`allowMethods`, `allowHeaders`, `exposeHeaders`의 경우 특정 이름 목록 대신
단일 와일드카드(`"*"`)를 사용할 수도 있다.

종합 예시:

{{< include file="examples/standard/http-cors/httproute-all-fields-set.yaml" >}}
