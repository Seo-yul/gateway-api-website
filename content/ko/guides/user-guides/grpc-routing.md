---
title: "gRPC 라우팅"
weight: 14
description: "Configuring gRPC routing with GRPCRoute"
---

# gRPC 라우팅

[GRPCRoute 리소스]({{< ref "/reference/api-types/grpcroute" >}})를 사용하면 gRPC 트래픽을 매칭하고
이를 쿠버네티스 백엔드로 전달할 수 있다. 이 가이드에서는 GRPCRoute가 호스트, 헤더,
서비스 및 메서드 필드를 기반으로 트래픽을 매칭하고 서로 다른 쿠버네티스 Service로
전달하는 방법을 설명한다.

다음 다이어그램은 세 가지 서로 다른 Service에 걸친 필수 트래픽 흐름을 설명한다:

- `foo.example.com`으로 향하는 `com.Example.Login` 메서드의 트래픽은 `foo-svc`로 전달된다
- `env: canary` 헤더가 포함된 `bar.example.com`으로 향하는 트래픽은
모든 서비스와 메서드에 대해 `bar-svc-canary`로 전달된다
- 해당 헤더가 없는 `bar.example.com`으로 향하는 트래픽은 모든 서비스와 메서드에 대해
  `bar-svc`로 전달된다

<!--- Editable source available at site-src/images/grpc-routing.png -->
![gRPC Routing](/images/grpc-routing.png)

점선은 이 라우팅 동작을 구성하기 위해 배포된 `Gateway` 리소스를 나타낸다.
동일한 `prod` 게이트웨이에 라우팅 규칙을 생성하는 두 개의 `GRPCRoute` 리소스가 있다.
이는 하나 이상의 라우트가 게이트웨이에 바인딩될 수 있음을 보여주며,
충돌하지 않는 한 라우트가 `Gateway`에서 병합될 수 있다. `GRPCRoute`는 동일한
라우트 병합 의미론을 따른다. 이에 대한 자세한 내용은
[문서]({{< ref "/reference/api-types/httproute#merging" >}})를 참조한다.

[게이트웨이][gateway]로부터 트래픽을 수신하려면 `GRPCRoute` 리소스에
연결할 상위 게이트웨이를 참조하는 `ParentRefs`를 구성해야 한다. 다음 예제는
`Gateway`와 `GRPCRoute`의 조합이 gRPC 트래픽을 서비스하도록 구성되는 방법을 보여준다:

{{< include file="examples/standard/grpc-routing/gateway.yaml" >}}

`GRPCRoute`는 [단일 호스트명 집합][spec]에 대해 매칭할 수 있다.
이러한 호스트명은 GRPCRoute 내의 다른 매칭보다 먼저 매칭된다.
`foo.example.com`과 `bar.example.com`은 서로 다른 라우팅 요구사항을 가진
별개의 호스트이므로, 각각 자체 GRPCRoute인 `foo-route`와 `bar-route`로 배포된다.

다음 `foo-route`는 `foo.example.com`에 대한 모든 트래픽을 매칭하고 라우팅 규칙을
적용하여 트래픽을 올바른 백엔드로 전달한다. 매치가 하나만 지정되어 있으므로
`foo.example.com`으로의 `com.example.User.Login` 메서드 요청만 전달된다.
다른 메서드의 RPC는 이 라우트에 의해 매칭되지 않는다.

{{< include file="examples/standard/grpc-routing/foo-grpcroute.yaml" >}}

마찬가지로 `bar-route` GRPCRoute는 `bar.example.com`에 대한 RPC를 매칭한다.
이 호스트명에 대한 모든 트래픽은 라우팅 규칙에 따라 평가된다. 가장 구체적인
매치가 우선하므로, `env: canary` 헤더가 있는 트래픽은 `bar-svc-canary`로
전달되고, 헤더가 없거나 값이 `canary`가 아닌 경우 `bar-svc`로 전달된다.

{{< include file="examples/standard/grpc-routing/bar-grpcroute.yaml" >}}

[gRPC
Reflection](https://github.com/grpc/grpc/blob/v1.49.1/doc/server-reflection.md)은
대상 서비스의 protocol buffers 사본을 로컬 파일시스템에 보유하지 않고도
[`grpcurl`](https://github.com/fullstorydev/grpcurl)과 같은 대화형 클라이언트를
사용하는 데 필요하다. 이를 활성화하려면 먼저 애플리케이션 파드에서 gRPC reflection
서버가 수신 대기 중인지 확인한 다음, `GRPCRoute`에 reflection 메서드를 추가한다.
이 기능은 개발 및 스테이징 환경에서 유용할 수 있지만, 프로덕션 환경에서는
보안상의 영향을 충분히 고려한 후에만 활성화해야 한다.

{{< include file="examples/standard/grpc-routing/reflection-grpcroute.yaml" >}}

[gateway]: {{< ref "/reference/spec#gateway" >}}
[spec]: {{< ref "/reference/spec#grpcroutespec" >}}
[svc]:https://kubernetes.io/docs/concepts/services-networking/service/
