---
title: "HTTP 라우팅"
linkTitle: "HTTP routing"
weight: 1
description: "Configuring HTTP routing with HTTPRoute"
---

# HTTP 라우팅

[HTTPRoute 리소스]({{< ref "/reference/api-types/httproute" >}})를 사용하면 HTTP **트래픽**을 매칭하고
이를 쿠버네티스 **백엔드**로 전달할 수 있다. 이 가이드에서는 HTTPRoute가 호스트, 헤더,
경로 필드를 기반으로 트래픽을 매칭하고 서로 다른 쿠버네티스 Service로 전달하는 방법을 설명한다.

다음 다이어그램은 세 가지 서로 다른 Service에 걸친 필수 트래픽 흐름을 설명한다:

- `foo.example.com/login`으로 향하는 트래픽은 `foo-svc`로 전달된다
- `env: canary` 헤더가 포함된 `bar.example.com/*`로 향하는 트래픽은
`bar-svc-canary`로 전달된다
- 해당 헤더가 없는 `bar.example.com/*`로 향하는 트래픽은 `bar-svc`로 전달된다

![HTTP Routing](/images/http-routing.png)

점선은 이 라우팅 동작을 구성하기 위해 배포된 **게이트웨이** 리소스를 나타낸다.
동일한 `prod-web` 게이트웨이에 라우팅 규칙을 생성하는 두 개의 HTTPRoute 리소스가 있다.
이는 하나 이상의 **라우트**가 게이트웨이에 바인딩될 수 있음을 보여주며,
충돌하지 않는 한 라우트가 게이트웨이에서 병합될 수 있다. 라우트 병합에 대한
자세한 내용은 [HTTPRoute 문서]({{< ref "/reference/api-types/httproute#merging" >}})를 참조한다.

[게이트웨이][gateway]로부터 트래픽을 수신하려면 `HTTPRoute` 리소스에
연결할 상위 게이트웨이를 참조하는 `ParentRefs`를 구성해야 한다. 다음 예제는
`Gateway`와 `HTTPRoute`의 조합이 HTTP 트래픽을 서비스하도록 구성되는 방법을 보여준다:

{{< include file="examples/standard/http-routing/gateway.yaml" >}}

HTTPRoute는 [단일 호스트명 집합][spec]에 대해 매칭할 수 있다.
이러한 호스트명은 HTTPRoute 내의 다른 매칭보다 먼저 매칭된다.
`foo.example.com`과 `bar.example.com`은 서로 다른 라우팅 요구사항을 가진
별개의 호스트이므로, 각각 자체 HTTPRoute인 `foo-route`와 `bar-route`로 배포된다.

다음 `foo-route`는 `foo.example.com`에 대한 모든 트래픽을 매칭하고 라우팅 규칙을
적용하여 트래픽을 올바른 백엔드로 전달한다. 매치가 하나만 지정되어 있으므로
`foo.example.com/login/*` 트래픽만 전달된다. `/login`으로 시작하지 않는
다른 경로로의 트래픽은 이 라우트에 의해 매칭되지 않는다.

{{< include file="examples/standard/http-routing/foo-httproute.yaml" >}}

마찬가지로 `bar-route` HTTPRoute는 `bar.example.com`에 대한 트래픽을 매칭한다.
이 호스트명에 대한 모든 트래픽은 라우팅 규칙에 따라 평가된다. 가장 구체적인
매치가 우선하므로, `env: canary` 헤더가 있는 트래픽은 `bar-svc-canary`로
전달되고, 헤더가 없거나 `canary`가 아닌 경우 `bar-svc`로 전달된다.

{{< include file="examples/standard/http-routing/bar-httproute.yaml" >}}

[gateway]: {{< ref "/reference/spec#gateway" >}}
[spec]: {{< ref "/reference/spec#httproutespec" >}}
[svc]:https://kubernetes.io/docs/concepts/services-networking/service/
