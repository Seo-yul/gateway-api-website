---
title: "HTTP 헤더 수정"
linkTitle: "HTTP header modifier"
weight: 3
description: "Modifying HTTP request and response headers with HTTPRoute"
---

# HTTP 헤더 수정

[HTTPRoute 리소스]({{< ref "/reference/api-types/httproute" >}})는 클라이언트의 HTTP 요청 헤더와 HTTP 응답 헤더를 수정할 수 있다.
이러한 요구사항을 충족하기 위해 두 가지 유형의 [필터]({{< ref "/reference/api-types/httproute#filters-optional" >}})를 사용할 수 있다: `RequestHeaderModifier`와 `ResponseHeaderModifier`.

이 가이드에서는 이러한 기능의 사용 방법을 설명한다.

이 기능들은 서로 호환된다는 점에 유의한다. 수신 요청의 HTTP 헤더와 해당 응답의 헤더를 단일 [HTTPRoute 리소스]({{< ref "/reference/api-types/httproute" >}})를 사용하여 모두 수정할 수 있다.

## HTTP 요청 헤더 수정

HTTP 헤더 수정은 수신 요청의 HTTP 헤더를 추가, 제거 또는 수정하는 과정이다.

HTTP 헤더 수정을 구성하려면 하나 이상의 HTTP 필터가 포함된 게이트웨이 객체를 정의한다. 각 필터는 사용자 정의 헤더 추가 또는 기존 헤더 수정과 같이 수신 요청에 대해 수행할 특정 수정 사항을 지정한다.

HTTP 요청에 헤더를 추가하려면 `RequestHeaderModifier` 유형의 필터를 `add` 액션과 함께 헤더의 이름 및 값을 지정하여 사용한다:

{{< include file="examples/standard/http-request-header-add.yaml" >}}

기존 헤더를 편집하려면 `set` 액션을 사용하고 수정할 헤더의 값과 설정할 새 헤더 값을 지정한다.

```yaml
    filters:
    - type: RequestHeaderModifier
      requestHeaderModifier:
        set:
          - name: my-header-name
            value: my-new-header-value
```

`remove` 키워드와 헤더 이름 목록을 사용하여 헤더를 제거할 수도 있다.

```yaml
    filters:
    - type: RequestHeaderModifier
      requestHeaderModifier:
        remove: ["x-request-id"]
```

위 예제를 사용하면 HTTP 요청에서 `x-request-id` 헤더가 제거된다.

### HTTP 응답 헤더 수정

요청 헤더를 편집하는 것이 유용한 것처럼, 응답 헤더도 마찬가지이다. 예를 들어, 특정 백엔드에 대해서만 쿠키를 추가/제거할 수 있으며, 이는 이전에 해당 백엔드로 리디렉션된 특정 사용자를 식별하는 데 도움이 될 수 있다.

또 다른 잠재적 사용 사례로, 프론트엔드가 안정 버전의 백엔드 서버와 통신하는지 베타 버전과 통신하는지를 알아야 하는 경우가 있을 수 있으며, 이를 통해 다른 UI를 렌더링하거나 응답 파싱을 조정할 수 있다.

HTTP 응답 헤더 수정은 원래 요청을 수정하는 데 사용되는 구문과 매우 유사한 구문을 사용하지만, 다른 필터(`ResponseHeaderModifier`)를 사용한다.

헤더를 추가, 편집 및 제거할 수 있다. 다음 예제에서 보여주는 것처럼 여러 헤더를 추가할 수 있다:

```yaml
    filters:
    - type: ResponseHeaderModifier
      responseHeaderModifier:
        add:
        - name: X-Header-Add-1
          value: header-add-1
        - name: X-Header-Add-2
          value: header-add-2
        - name: X-Header-Add-3
          value: header-add-3
```
