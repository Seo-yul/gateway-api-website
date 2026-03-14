---
title: "GRPCRoute"
weight: 30
description: "Routing gRPC requests from Gateway listeners to services"
---

{{< channel-version channel="standard" version="v1.1.0" >}}

`GRPCRoute` 리소스는 GA(정식 출시)되었으며 `v1.1.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하자.
{{< /channel-version >}}

[GRPCRoute][grpcroute]는 게이트웨이(Gateway) 리스너에서 API 객체(예: 서비스(Service))로의
gRPC 요청 라우팅 동작을 지정하기 위한 Gateway API 타입이다.

## 배경

`HTTPRoutes`를 사용하거나 사용자 정의 외부(out-of-tree) CRD를 통해 gRPC를 라우팅하는 것이
가능하지만, 장기적으로 이는 생태계의 파편화를 초래한다.

gRPC는 [업계 전반에 널리 채택된 인기 있는 RPC 프레임워크](https://grpc.io/about/#whos-using-grpc-and-why)이다.
이 프로토콜은 쿠버네티스 프로젝트 자체 내에서 다음을 포함한 많은 인터페이스의
기반으로 광범위하게 사용되고 있다.

- [CSI](https://github.com/container-storage-interface/spec/blob/5b0d4540158a260cb3347ef1c87ede8600afb9bf/spec.md),
- [CRI](https://github.com/kubernetes/cri-api/blob/49fe8b135f4556ea603b1b49470f8365b62f808e/README.md),
- [디바이스 플러그인 프레임워크](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/)

애플리케이션 계층 네트워킹 영역과 특히 쿠버네티스 프로젝트에서 gRPC의 중요성을 고려하여,
생태계가 불필요하게 파편화되는 것을 허용하지 않기로 결정하였다.

### 캡슐화된 네트워크 프로토콜

일반적으로, 캡슐화된 프로토콜을 하위 수준에서 라우팅할 수 있는 경우,
다음 기준이 충족되면 상위 계층에서 라우트 리소스를 도입하는 것이 허용된다.

- 캡슐화된 프로토콜의 사용자가 하위 계층에서 라우팅하도록 강제될 경우, 해당 생태계의 중요한 기존 기능을 놓칠 수 있다.
- 캡슐화된 프로토콜의 사용자가 하위 계층에서 라우팅하도록 강제될 경우, 저하된 사용자 경험을 겪을 수 있다.
- 캡슐화된 프로토콜이 특히 쿠버네티스 커뮤니티에서 상당한 사용자 기반을 보유하고 있다.

gRPC는 이러한 모든 기준을 충족하므로, Gateway API에 `GRPCRoute`를 포함하기로 결정하였다.

### 교차 서빙 (Cross Serving) {#cross-serving}

GRPCRoute를 지원하는 구현체는 `GRPCRoute`와 `HTTPRoute` 간의 호스트네임
고유성을 강제해야 한다. `HTTPRoute` 또는 `GRPCRoute` 타입의 라우트(A)가
리스너에 연결되어 있고, 해당 리스너에 이미 다른 타입의 또 다른 라우트(B)가
연결되어 있으며, A와 B의 호스트네임 교집합이 비어 있지 않은 경우,
구현체는 라우트 A를 거부해야 한다. 즉, 구현체는 해당
RouteParentStatus에서 'Accepted' 조건의 상태를 'False'로 설정해야 한다.

일반적으로 gRPC와 비-gRPC HTTP 트래픽에 대해 별도의 호스트네임을 사용하는 것이
권장된다. 이는 gRPC 커뮤니티의 표준 관행과 일치한다.
그러나 URI만을 구분자로 하여 동일한 호스트네임에서 HTTP와 gRPC를 모두 제공해야
하는 경우, 사용자는 gRPC와 HTTP 모두에 대해 `HTTPRoute` 리소스를 사용해야 한다.
이 경우 `GRPCRoute` 리소스의 향상된 사용자 경험(UX)을 포기해야 한다.

## GRPCRoute 사용 시점

gRPC 트래픽은 `HTTPRoute` 또는 `GRPCRoute`를 사용하여 라우팅할 수 있다. 기본적인
gRPC 로드 밸런싱의 경우, `HTTPRoute`로도 가능하므로 `GRPCRoute`가 기술적으로
필수는 아니다. 다음 가이드는 사용 사례에 적합한 리소스를 선택하는 데 도움이 된다.

### HTTPRoute가 충분한 경우

`HTTPRoute`는 gRPC 전용 기능 없이 gRPC 트래픽에 대한 기본 라우팅과 로드 밸런싱만
필요한 경우에 충분하다. 다음과 같은 경우 `HTTPRoute`를 사용한다.

- 단순한 호스트네임 또는 경로 기반 라우팅과 로드 밸런싱만 필요한 경우.
- gRPC 인식 재시도 또는 메트릭이 필요하지 않은 경우(예: gRPC 상태 코드 기반 재시도, gRPC 지향 메트릭).
- URI만을 구분자로 하여 동일한 호스트네임에서 HTTP와 gRPC를 모두 제공해야 하는 경우,
  이때는 두 트래픽 모두에 `HTTPRoute`를 사용해야 한다(위의 [교차 서빙](#cross-serving) 참조).

`HTTPRoute`에서는 gRPC 메서드를 URI 경로로 지정하여 매칭한다
(예: `/package.Service/Method`). `GRPCRoute`에서는 서비스와 메서드 필드로 매칭한다
(예: `service: com.example.User`, `method: Login`).

### GRPCRoute를 선호하는 경우

gRPC 전용 동작이 필요하거나 gRPC 사용자를 위한 더 나은 사용성이 필요한 경우
`GRPCRoute`를 선호한다. 다음과 같은 경우 `GRPCRoute`를 사용한다.

- URI 경로 대신 gRPC 서비스 및 메서드 이름으로 직접 매칭하려는 경우
  (예: `service: com.example.User`, `method: Login`).
- gRPC 상태 코드에 조건부 재시도와 같은 gRPC 인식 정책이 필요한 경우
  (예: `CANCELLED`, `RESOURCE_EXHAUSTED`).
- 구현체의 gRPC 지향 메트릭 또는 관찰 가능성이 필요한 경우.

### 컨트롤러 구현자를 위한 가이드

사용자를 대신하여 라우트를 생성하는 컨트롤러(예: Gateway API 라우트를 생성하는
상위 수준 API)는 다음을 고려해야 한다.

- 기본 gRPC 로드 밸런싱을 위해 `HTTPRoute`를 생성하는 경우, gRPC 전용 기능이
  필요한 사용자가 수동으로 라우트를 관리하지 않고도 해당 기능을 사용할 수 있도록
  `GRPCRoute` 옵션 제공을 고려한다.
- gRPC 상태 코드 기반 재시도와 같은 gRPC 전용 정책을 `HTTPRoute`에 추가하지 않는다.
  gRPC 인식 구성은 `GRPCRoute`에만 유지한다.

## 스펙

GRPCRoute의 사양은 다음으로 구성된다.

- [ParentRefs][parentRef]- 이 라우트가 연결되고자 하는 게이트웨이를 정의한다.
- [Hostnames][hostname] (선택 사항)- gRPC 요청의 Host 헤더 매칭에 사용할
  호스트네임 목록을 정의한다.
- [Rules][grpcrouterule]- 매칭되는 gRPC 요청에 대해 수행할 작업 목록을 정의한다.
  각 규칙은 [matches][matches], [filters][filters] (선택 사항),
  [backendRefs][backendRef] (선택 사항), [name][name] (선택 사항) 필드로 구성된다.

<!--- Editable SVG available at site-src/images/grpcroute-basic-example.svg -->
다음은 모든 트래픽을 하나의 서비스로 전송하는 GRPCRoute를 보여준다.
![grpcroute-basic-example](/images/grpcroute-basic-example.png)

### 게이트웨이에 연결하기

각 라우트에는 연결하고자 하는 상위 리소스를 참조하는 방법이 포함되어 있다.
대부분의 경우 이는 게이트웨이가 되지만, 구현체가 다른 유형의 상위 리소스를
지원할 수 있는 유연성이 있다.

다음 예시는 라우트가 `acme-lb` 게이트웨이에 연결되는 방법을 보여준다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: GRPCRoute
metadata:
  name: grpcroute-example
spec:
  parentRefs:
  - name: acme-lb
```

대상 게이트웨이가 연결이 성공하려면 라우트의 네임스페이스에서
GRPCRoute의 연결을 허용해야 한다.

### 호스트네임 (Hostnames)

호스트네임은 gRPC 요청의 Host 헤더와 매칭할 호스트네임 목록을 정의한다.
매칭이 발생하면, 규칙과 필터(선택 사항)에 기반하여 요청 라우팅을 수행하기 위해
GRPCRoute가 선택된다. 호스트네임은 [RFC 3986][rfc-3986]에 정의된
네트워크 호스트의 완전 정규화된 도메인 이름(FQDN)이다.
RFC에 정의된 URI의 "host" 부분과 다음과 같은 차이점이 있다.

- IP는 허용되지 않는다.
- 포트가 허용되지 않으므로 `:` 구분자는 사용되지 않는다.

수신 요청은 GRPCRoute 규칙이 평가되기 전에 호스트네임과 매칭된다.
호스트네임이 지정되지 않은 경우, GRPCRoute 규칙과 필터(선택 사항)에 기반하여
트래픽이 라우팅된다.

다음 예시는 호스트네임 "my.example.com"을 정의한다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: GRPCRoute
metadata:
  name: grpcroute-example
spec:
  hostnames:
  - my.example.com
```

### 규칙 (Rules)

규칙은 조건에 기반한 gRPC 요청 매칭, 선택적인 추가 처리 단계 실행,
그리고 선택적으로 요청을 API 객체로 전달하는 의미를 정의한다.

#### 매칭 (Matches)

매칭은 gRPC 요청을 매칭하는 데 사용되는 조건을 정의한다. 각 매칭은
독립적이며, 즉 단일 매칭이 만족되면 이 규칙이 매칭된다.

다음 매칭 구성을 예시로 살펴보자.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: GRPCRoute
...
matches:
  - method:
      service: com.example.User
      method: Login
    headers:
    - name: version
      value: "2"
  - method:
      service: com.example.v2.User
      method: Login
```

이 규칙에 대해 요청이 매칭되려면 다음 조건 중 하나를 만족해야 한다.

 - `com.example.User.Login` 메서드 **그리고** 헤더 "version: 2"를 포함
 - `com.example.v2.User.Login` 메서드

매칭이 지정되지 않은 경우, 기본값은 모든 gRPC 요청과 매칭되는 것이다.

#### 필터 (Filters, 선택 사항)

필터는 요청 또는 응답 수명 주기 동안 완료되어야 하는 처리 단계를 정의한다.
필터는 게이트웨이 구현체에서 수행할 수 있는 추가 처리를 표현하기 위한
확장 지점으로 작동한다. 일부 예시로는 요청 또는 응답 수정,
인증 전략 구현, 속도 제한, 트래픽 셰이핑 등이 있다.

다음 예시는 Host 헤더가 "my.filter.com"인 gRPC 요청에 헤더 "my-header: foo"를
추가한다. GRPCRoute는 이와 같이 HTTPRoute와 기능이 동일한 기능에 대해
HTTPRoute 필터를 사용한다는 점에 유의하자.

```yaml
{{< include file="examples/standard/grpc-filter.yaml" >}}
```

API 적합성은 필터 타입에 기반하여 정의된다. 여러 동작의 순서에 대한 효과는
현재 지정되지 않았다. 이는 알파 단계의 피드백에 따라 향후 변경될 수 있다.

적합성 수준은 필터 타입에 의해 정의된다.

 - 모든 "core" 필터는 GRPCRoute를 지원하는 구현체에서 반드시 지원해야 한다(MUST).
 - 구현자는 "extended" 필터를 지원하는 것이 권장된다.
 - "Implementation-specific" 필터는 구현체 간에 API 보장이 없다.

core 필터를 여러 번 지정하는 것은 지정되지 않았거나 사용자 정의 적합성을 가진다.

구현체가 필터 조합을 지원할 수 없는 경우, 해당 제한 사항을 명확하게
문서화해야 한다. 호환되지 않거나 지원되지 않는 필터가 지정되어
`Accepted` 조건의 상태가 `False`로 설정되는 경우, 구현체는
이 구성 오류를 지정하기 위해 `IncompatibleFilters` 사유를 사용할 수 있다.

#### BackendRefs (선택 사항)

BackendRefs는 매칭된 요청이 전송되어야 하는 API 객체를 정의한다.
지정되지 않은 경우, 규칙은 전달을 수행하지 않는다. 지정되지 않았고
응답이 전송되는 필터도 지정되지 않은 경우, `UNIMPLEMENTED` 오류 코드가 반환된다.

다음 예시는 `User.Login` 메서드에 대한 gRPC 요청을 포트 `50051`의
서비스 "my-service1"로, 헤더 `magic: foo`가 포함된 `Things.DoThing` 메서드에 대한
gRPC 요청을 포트 `50051`의 서비스 "my-service2"로 전달한다.

```yaml
{{< include file="examples/standard/basic-grpc.yaml" >}}
```

다음 예시는 `weight` 필드를 사용하여 `foo.example.com`에 대한 gRPC 요청의
90%를 "foo-v1" 서비스로, 나머지 10%를 "foo-v2" 서비스로 전달한다.

```yaml
{{< include file="examples/standard/traffic-splitting/grpc-traffic-split-2.yaml" >}}
```

`weight` 및 기타 필드에 대한 추가 세부 사항은 [backendRef][backendRef] API
문서를 참조하자.

#### 이름 (Name, 선택 사항)

{{< channel-version channel="experimental" version="v1.2.0" >}}

이 개념은 `v1.2.0` 부터 실험적 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하자.
{{< /channel-version >}}

GRPCRoute 규칙에는 선택적인 `name` 필드가 포함된다. 라우트 규칙 이름의 용도는 구현체별로 다르다. 메타리소스의 `sectionName` 필드([GEP-2648]({{< ref "/geps/gep-2648" >}}#section-names))에서와 같이 다른 리소스에서 개별 라우트 규칙을 이름으로 참조하거나, 라우트 객체와 관련된 리소스의 상태 스탠자에서, 또는 구현체가 GRPCRoute 규칙에서 생성한 내부 구성 객체를 식별하는 데 사용할 수 있다.

지정된 경우, name 필드의 값은 [`SectionName`](https://github.com/kubernetes-sigs/gateway-api/blob/v1.0.0/apis/v1/shared_types.go#L607-L624) 타입을 준수해야 한다.

## 상태 (Status)

상태(Status)는 GRPCRoute의 관찰된 상태를 정의한다.

### RouteStatus

RouteStatus는 모든 라우트 타입에 걸쳐 필요한 관찰된 상태를 정의한다.

#### 부모 (Parents)

부모(Parents)는 GRPCRoute와 연관된 게이트웨이(또는 기타 상위 리소스)의
목록과 각 게이트웨이에 대한 GRPCRoute의 상태를 정의한다.
GRPCRoute가 parentRefs에 게이트웨이 참조를 추가하면,
게이트웨이를 관리하는 컨트롤러는 라우트를 처음 확인할 때 이 목록에
항목을 추가해야 하며, 라우트가 수정될 때 적절하게 항목을 업데이트해야 한다.

## 예시

다음 예시는 GRPCRoute "grpc-example"이 네임스페이스 "gw-example-ns"의
게이트웨이 "gw-example"에 의해 수락되었음을 나타낸다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: GRPCRoute
metadata:
  name: grpc-example
...
status:
  parents:
  - parentRefs:
      name: gw-example
      namespace: gw-example-ns
    conditions:
    - type: Accepted
      status: "True"
```

## 병합 (Merging)
여러 GRPCRoute를 단일 게이트웨이 리소스에 연결할 수 있다. 중요한 점은
각 요청에 대해 하나의 라우트 규칙만 매칭될 수 있다는 것이다. 충돌 해결이
병합에 어떻게 적용되는지에 대한 자세한 정보는
[API 사양][grpcrouterule]을 참조하자.

[grpcroute]: ../../reference/spec.md#grpcroute
[grpcrouterule]: ../../reference/spec.md#grpcrouterule
[hostname]: ../../reference/spec.md#hostname
[rfc-3986]: https://tools.ietf.org/html/rfc3986
[matches]: ../../reference/spec.md#grpcroutematch
[filters]: ../../reference/spec.md#grpcroutefilter
[backendRef]: ../../reference/spec.md#grpcbackendref
[parentRef]: ../../reference/spec.md#parentreference
[name]: ../../reference/spec.md#sectionname
