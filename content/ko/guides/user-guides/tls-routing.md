---
title: "TLS 라우팅"
linkTitle: "TLS routing"
weight: 12
description: "Configuring TLS routing with TLSRoute"
---

# TLS 라우팅

[**TLSRoute**(TLS 라우트) 리소스]({{< ref "/reference/api-types/tlsroute" >}})를 사용하면 TLS
메타데이터를 기준으로 매칭하여 쿠버네티스 백엔드로 트래픽을 전달할 수 있다. 이 가이드에서는
TLSRoute가 호스트네임을 기준으로 트래픽을 매칭하고, **Gateway**(게이트웨이)에서
`Passthrough` 또는 `Terminate` TLS 모드를 사용하여 서로 다른 쿠버네티스 서비스로
트래픽을 전달하는 방법을 설명한다.

[게이트웨이][gateway]로부터 트래픽을 수신하려면 TLSRoute 리소스에
연결할 상위 게이트웨이를 참조하는 `ParentRefs`를 구성해야 한다.
다음 예제는 게이트웨이와 TLSRoute를 조합하여 `Passthrough` 및
`Terminate` 모드를 모두 사용하여 TLS 트래픽을 처리하는 방법을
보여준다(Gateway API 구현체에서 지원하는 경우):

{{< include file="examples/standard/tls-routing/gateway.yaml" >}}

TLSRoute는 [단일 호스트네임 집합][spec]에 대해 매칭할 수 있다.
`foo.example.com`과 `bar.example.com`은 서로 다른 라우팅 요구 사항을 가진
별도의 호스트이므로, 각각 `foo-route`와 `bar-route`라는 별도의
TLSRoute로 배포된다.

다음 `foo-route` TLSRoute는 `foo.example.com`에 대한 모든 트래픽을 매칭하고,
라우팅 규칙을 적용하여 구성된 백엔드로 트래픽을 전달한다.
`Passthrough` 모드로 구성된 리스너에 연결되어 있으므로,
게이트웨이는 암호화된 TCP 스트림을 백엔드로 직접 전달한다:

{{< include file="examples/standard/tls-routing/tls-route.yaml" >}}

마찬가지로, `bar-route` TLSRoute는 `bar.example.com`에 대한 트래픽을 매칭한다.
그러나 `Terminate` 모드로 구성된 리스너에 연결되어 있으므로,
게이트웨이는 먼저 리스너에 지정된 인증서를 사용하여 TLS 스트림을 종료한 다음,
결과로 생성된 비암호화 TCP 스트림을 백엔드로 전달한다:

{{< include file="examples/standard/tls-routing/tls-route-terminate.yaml" >}}

[gateway]: {{< ref "/reference/spec#gateway" >}}
[spec]: {{< ref "/reference/spec#tlsroutespec" >}}
