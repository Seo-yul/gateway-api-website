---
title: "HTTP 요청 미러링"
linkTitle: "HTTP request mirroring"
weight: 5
description: "Mirroring HTTP requests to multiple backends"
---

# HTTP 요청 미러링

{{< collapsible-alert color="info" title="확장 지원 기능: HTTPRouteRequestMirror" open="true" >}}
이 기능은 확장 지원의 일부이다. 지원 수준에 대한 자세한 내용은 [적합성 가이드]({{< ref "/overview/concepts/conformance" >}})를 참조하라.
{{< /collapsible-alert >}}

[HTTPRoute 리소스]({{< ref "/reference/api-types/httproute" >}})를 사용하여 요청을 여러 백엔드로
미러링할 수 있다. 이는 프로덕션 트래픽으로 새로운 서비스를 테스트하는 데
유용하다.

미러링된 요청은 이 backendRef 내의 하나의 단일 대상 엔드포인트로만 전송되며,
이 백엔드의 응답은 Gateway에 의해 반드시 무시되어야 한다(MUST).

요청 미러링은 블루-그린 배포에서 특히 유용하다. 클라이언트에 대한 응답에 어떠한
영향도 미치지 않으면서 애플리케이션 성능에 대한 영향을 평가하는 데 사용할 수
있다.

{{< include file="examples/standard/http-request-mirroring/httproute-mirroring.yaml" >}}

이 예제에서는 모든 요청이 포트 `8080`의 서비스 `foo-v1`로 전달되고, 포트 `8080`의
서비스 `foo-v2`로도 전달되지만, 응답은 서비스 `foo-v1`에서만 생성된다.
