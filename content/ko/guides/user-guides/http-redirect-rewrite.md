---
title: "HTTP 경로 리다이렉트와 리라이트"
linkTitle: "HTTP redirects and rewrites"
weight: 2
description: "Configuring HTTP redirects and URL rewrites with HTTPRoute"
---

# HTTP 경로 리다이렉트(Redirect)와 리라이트(Rewrite)

[HTTPRoute 리소스]({{< ref "/reference/api-types/httproute" >}})는
[필터]({{< ref "/reference/api-types/httproute#filters-optional" >}})를 사용하여
클라이언트에게 리다이렉트를 발행하거나 업스트림으로 전송되는 경로를 리라이트할 수 있다.
이 가이드는 이러한 기능의 사용 방법을 보여준다.

리다이렉트 필터와 리라이트 필터는 상호 호환되지 않는다. 규칙에서 두 필터 타입을
동시에 사용할 수 없다.

## 리다이렉트(Redirect)

리다이렉트는 클라이언트에게 HTTP 3XX 응답을 반환하여 다른 리소스를
가져오도록 지시한다.
[`RequestRedirect` 규칙 필터]({{< ref "/reference/spec#httprequestredirectfilter" >}})는
필터링된 HTTPRoute 규칙과 일치하는 요청에 대해 게이트웨이가
리다이렉트 응답을 보내도록 지시한다.

### 지원되는 상태 코드

게이트웨이 API는 다음과 같은 HTTP 리다이렉트 상태 코드를 지원한다.

- **301 (Moved Permanently)**: 리소스가 새 위치로 영구적으로 이동했음을 나타낸다. 검색 엔진과 클라이언트는 새 URL을 사용하도록 참조를 업데이트한다. HTTP에서 HTTPS로의 업그레이드나 영구적인 URL 변경과 같은 영구 리다이렉트에 사용한다.

- **302 (Found)**: 리소스가 일시적으로 다른 위치에서 사용 가능함을 나타낸다. 상태 코드를 지정하지 않으면 이것이 기본 상태 코드이다. 원래 URL이 향후 다시 유효해질 수 있는 임시 리다이렉트에 사용한다.

- **303 (See Other)**: 요청에 대한 응답을 GET 메서드를 사용하여 다른 URL에서 찾을 수 있음을 나타낸다. 이는 일반적으로 POST 요청 후 확인 페이지로 리다이렉트하여 중복 폼 제출을 방지하는 데 사용된다.

- **307 (Temporary Redirect)**: 302와 유사하지만, 리다이렉트를 따를 때 HTTP 메서드가 변경되지 않음을 보장한다. 리다이렉트에서 원래 HTTP 메서드(POST, PUT 등)를 보존해야 할 때 사용한다.

- **308 (Permanent Redirect)**: 301과 유사하지만, 리다이렉트를 따를 때 HTTP 메서드가 변경되지 않음을 보장한다. HTTP 메서드를 보존해야 하는 영구 리다이렉트에 사용한다.

리다이렉트 필터는 다양한 URL 구성 요소를 독립적으로 대체할 수 있다. 예를 들어,
HTTP에서 HTTPS로의 영구 리다이렉트(301)를 발행하려면
`requestRedirect.statusCode=301`과 `requestRedirect.scheme="https"`를 설정한다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-redirect-http.yaml" >}}
```

리다이렉트는 설정된 URL 구성 요소를 리다이렉트 설정과 일치하도록 변경하면서
원래 요청 URL의 다른 구성 요소는 보존한다. 이 예제에서
`GET http://redirect.example/cinnamon` 요청은
`location: https://redirect.example/cinnamon` 헤더가 포함된 301 응답을 반환한다.
호스트네임(`redirect.example`), 경로(`/cinnamon`), 포트(암시적)는
변경되지 않는다.

### 메서드 보존 리다이렉트

리다이렉트 중에 HTTP 메서드가 보존되어야 하는 경우, 상태 코드 307 또는 308을 사용한다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-redirect-307.yaml" >}}
```

HTTP 메서드를 보존해야 하는 영구 리다이렉트의 경우, 상태 코드 308을 사용한다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-redirect-308.yaml" >}}
```

### POST-Redirect-GET 패턴

POST-Redirect-GET 패턴을 구현하려면, 상태 코드 303을 사용하여 POST 요청을 GET 엔드포인트로 리다이렉트한다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-redirect-303.yaml" >}}
```

### HTTP에서 HTTPS로의 리다이렉트

HTTP 트래픽을 HTTPS로 리다이렉트하려면, HTTP와 HTTPS 리스너를 모두 가진
게이트웨이(Gateway)가 필요하다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/gateway-redirect-http-https.yaml" >}}
```
게이트웨이를 보호하는 방법은 여러 가지가 있다. 이 예제에서는
쿠버네티스 시크릿(Secret)을 사용하여 보호한다(`certificateRefs` 섹션의 `redirect-example`).

HTTP 리스너에 연결되어 HTTPS로 리다이렉트하는 HTTPRoute가 필요하다.
여기서 `sectionName`을 `http`로 설정하여 `http`라는 이름의
리스너만 선택하도록 한다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-redirect-http.yaml" >}}
```

또한 HTTPS 리스너에 연결되어 HTTPS 트래픽을 애플리케이션 백엔드로
전달하는 HTTPRoute도 필요하다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-redirect-https.yaml" >}}
```

### 경로 리다이렉트

경로 리다이렉트는 HTTP 경로 수정자(HTTP Path Modifier)를 사용하여 전체 경로 또는 경로
접두사를 교체한다. 예를 들어, 아래의 HTTPRoute는 `/cayenne`로 시작하는 경로를 가진
모든 `redirect.example` 요청을 `/paprika`로 302 리다이렉트한다.
특정 요구사항에 따라 지원되는 상태 코드(301, 302, 303, 307, 308) 중
어떤 것이든 사용할 수 있다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-redirect-full.yaml" >}}
```

`https://redirect.example/cayenne/pinch`와
`https://redirect.example/cayenne/teaspoon`에 대한 요청 모두
`location: https://redirect.example/paprika` 리다이렉트를 수신하게 된다.

다른 경로 리다이렉트 타입인 `ReplacePrefixMatch`는 `matches.path.value`와
일치하는 경로 부분만 교체한다. 위의 필터를 다음과 같이 변경하면:

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-redirect-prefix.yaml" >}}
```

`location: https://redirect.example/paprika/pinch`와
`location: https://redirect.example/paprika/teaspoon`
응답 헤더가 포함된 리다이렉트가 발생한다.

## 리라이트(Rewrite)

리라이트는 업스트림으로 프록시하기 전에 클라이언트 요청의 구성 요소를 수정한다.
[`URLRewrite` 필터]({{< ref "/reference/spec#httpurlrewritefilter" >}})는
업스트림 요청의 호스트네임 및/또는 경로를 변경할 수 있다.
예를 들어, 다음 HTTPRoute는
`https://rewrite.example/cardamom` 요청을 수신하여
`host: rewrite.example` 대신 `host: elsewhere.example`
요청 헤더와 함께 `example-svc` 업스트림으로 전송한다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-rewrite.yaml" >}}
```

경로 리라이트도 HTTP 경로 수정자(HTTP Path Modifier)를 활용한다.
아래의 HTTPRoute는 `https://rewrite.example/cardamom/smidgen` 요청을
`example-svc` 업스트림의 `https://elsewhere.example/fennel`로 프록시한다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-rewrite-full-path.yaml" >}}
```

대신 `type: ReplacePrefixMatch`와 `replacePrefixMatch: /fennel`을 사용하면
업스트림에 `https://elsewhere.example/fennel/smidgen`을 요청하게 된다.

```yaml
{{< include file="examples/standard/http-redirect-rewrite/httproute-rewrite-prefix-path.yaml" >}}
```
