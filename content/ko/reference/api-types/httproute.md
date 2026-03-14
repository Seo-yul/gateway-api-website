---
title: "HTTPRoute"
weight: 40
description: "Routing HTTP requests from Gateway listeners to services"
---

{{< channel-version channel="standard" version="v0.5.0" >}}

`HTTPRoute` 리소스는 GA(정식 출시)되었으며 `v0.5.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하라.
{{< /channel-version >}}

[HTTPRoute][httproute]는 **Gateway**(게이트웨이) 리스너에서 API 객체(예: Service)로의
HTTP 요청 라우팅 동작을 지정하기 위한 Gateway API 타입이다.

## 사양

HTTPRoute의 사양은 다음으로 구성된다:

- [ParentRefs][parentRef]- 이 라우트가 연결되고자 하는 게이트웨이를 정의한다.
- [Hostnames][hostname] (선택 사항)- HTTP 요청의 Host 헤더와 매칭하기 위한
  호스트네임 목록을 정의한다.
- [Rules][httprouterule]- 매칭되는 HTTP 요청에 대한 작업 규칙 목록을 정의한다.
  각 규칙은 [matches][matches], [filters][filters] (선택 사항),
  [backendRefs][backendRef] (선택 사항), [timeouts][timeouts] (선택 사항),
  [name][sectionName] (선택 사항) 필드로 구성된다.

다음은 모든 트래픽을 하나의 Service로 전송하는 HTTPRoute를 보여준다:
![httproute-basic-example](/images/httproute-basic-example.svg)

### 게이트웨이에 연결하기

각 라우트에는 연결하고자 하는 부모 리소스를 참조하는 방법이 포함되어 있다.
대부분의 경우 게이트웨이가 되지만, 구현에서 다른 타입의 부모 리소스를
지원하는 유연성도 있다.

다음 예시는 라우트가 `acme-lb` 게이트웨이에 연결하는 방법을 보여준다:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: httproute-example
spec:
  parentRefs:
  - name: acme-lb
```

대상 게이트웨이는 라우트의 네임스페이스에서 HTTPRoute의 연결을 허용해야
연결이 성공한다는 점에 유의하라.

부모 리소스의 특정 섹션에 라우트를 연결할 수도 있다.
예를 들어, `acme-lb` 게이트웨이에 다음과 같은 리스너가 포함되어 있다고
가정하자:

```yaml
  listeners:
  - name: foo
    protocol: HTTP
    port: 8080
    ...
  - name: bar
    protocol: HTTP
    port: 8090
    ...
  - name: baz
    protocol: HTTP
    port: 8090
    ...
```

`parentRefs`의 `sectionName` 필드를 사용하여 리스너 `foo`에만
라우트를 바인딩할 수 있다:

```yaml
spec:
  parentRefs:
  - name: acme-lb
    sectionName: foo
```

또는 `parentRefs`에서 `sectionName` 대신 `port` 필드를 사용하여
같은 효과를 얻을 수 있다:

```yaml
spec:
  parentRefs:
  - name: acme-lb
    port: 8080
```

포트에 바인딩하면 한 번에 여러 리스너에 연결할 수도 있다.
예를 들어, `acme-lb` 게이트웨이의 포트 `8090`에 바인딩하는 것이
이름으로 해당 리스너에 바인딩하는 것보다 더 편리할 수 있다:

```yaml
spec:
  parentRefs:
  - name: acme-lb
    sectionName: bar
  - name: acme-lb
    sectionName: baz
```

그러나 포트 번호로 라우트를 바인딩할 때, 게이트웨이 관리자는 라우트를 함께
업데이트하지 않고는 게이트웨이의 포트를 변경할 수 있는 유연성을 잃게 된다.
이 접근 방식은 포트가 변경될 수 있는 리스너가 아닌 특정 포트 번호에
라우트를 적용해야 하는 경우에만 사용해야 한다.

### 호스트네임(Hostnames)

호스트네임은 HTTP 요청의 Host 헤더와 매칭할 호스트네임 목록을 정의한다.
매칭이 발생하면 규칙 및 필터(선택 사항)에 따라 요청 라우팅을 수행할
HTTPRoute가 선택된다. 호스트네임은 [RFC 3986][rfc-3986]에 정의된
네트워크 호스트의 완전한 도메인 이름(FQDN)이다. RFC에서 정의한 URI의
"host" 부분과 다른 다음 사항에 유의하라:

- IP는 허용되지 않는다.
- 포트가 허용되지 않기 때문에 : 구분자는 사용되지 않는다.

수신 요청은 HTTPRoute 규칙이 평가되기 전에 호스트네임과 매칭된다.
호스트네임이 지정되지 않으면 HTTPRoute 규칙 및 필터(선택 사항)에 따라
트래픽이 라우팅된다.

다음 예시는 호스트네임 "my.example.com"을 정의한다:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: httproute-example
spec:
  hostnames:
  - my.example.com
```

### 규칙(Rules)

규칙은 조건에 따라 HTTP 요청을 매칭하고, 선택적으로 추가 처리 단계를 실행하며,
선택적으로 요청을 API 객체로 전달하는 의미를 정의한다.

#### 매칭(Matches)

매칭은 HTTP 요청을 매칭하기 위한 조건을 정의한다. 각 매칭은
독립적이며, 단일 매칭이 충족되면 이 규칙이 매칭된다.

다음 매칭 구성을 예로 살펴보자:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
...
spec:
  rules:
  - matches:
    - path:
        value: "/foo"
      headers:
      - name: "version"
        value: "2"
    - path:
        value: "/v2/foo"
```

요청이 이 규칙에 매칭되려면 다음 조건 중 하나를 충족해야 한다:

 - /foo로 시작하는 경로 **그리고** "version: 2" 헤더를 포함하는 경우
 - /v2/foo 경로 접두사인 경우

매칭이 지정되지 않으면 기본값은 "/"에 대한 접두사 경로 매칭이며,
이는 모든 HTTP 요청을 매칭하는 효과가 있다.

#### 필터(Filters) (선택 사항)

필터는 요청 또는 응답 수명 주기 동안 완료해야 하는 처리 단계를 정의한다.
필터는 게이트웨이 구현에서 수행할 수 있는 추가 처리를 표현하기 위한
확장 지점 역할을 한다. 일부 예시로는 요청 또는 응답 수정, 인증 전략 구현,
속도 제한, 트래픽 셰이핑 등이 있다.

다음 예시는 Host 헤더가 "my.filter.com"인 HTTP 요청에
"my-header: foo" 헤더를 추가한다.
```yaml
{{< include file="examples/standard/http-filter.yaml" >}}
```

API 적합성은 필터 타입에 따라 정의된다. 여러 동작의 순서 효과는
현재 지정되지 않았다. 이는 알파 단계의 피드백에 따라
향후 변경될 수 있다.

적합성 수준은 필터 타입에 따라 정의된다:

 - 모든 "core" 필터는 구현에서 반드시 지원해야 한다(MUST).
 - 구현자는 "extended" 필터를 지원하는 것이 권장된다.
 - "Implementation-specific" 필터는 구현 간 API 보장이 없다.

core 필터를 여러 번 지정하는 것은 미지정이거나 구현별 적합성을 갖는다.

모든 필터는 서로 호환될 것으로 예상되지만, URLRewrite와 RequestRedirect 필터는
결합할 수 없는 예외이다. 구현이 다른 필터 조합을 지원할 수 없는 경우,
해당 제한 사항을 명확히 문서화해야 한다. 호환되지 않거나 지원되지 않는
필터가 지정되어 `Accepted` 조건이 `False` 상태로 설정되는 경우,
구현은 이 구성 오류를 지정하기 위해 `IncompatibleFilters` 사유를 사용할 수 있다.

#### BackendRefs (선택 사항)

BackendRefs는 매칭된 요청이 전송되어야 하는 API 객체를 정의한다.
지정되지 않으면 규칙은 전달을 수행하지 않는다. 지정되지 않고
응답을 보내는 결과가 되는 필터도 지정되지 않으면 404 오류 코드가 반환된다.

다음 예시는 경로 접두사 `/bar`에 대한 HTTP 요청을 포트 `8080`의
"my-service1" 서비스로 전달하고, 다음 네 가지 기준을 _모두_ 충족하는
HTTP 요청을

- 헤더 `magic: foo`
- 쿼리 파라미터 `great: example`
- 경로 접두사 `/some/thing`
- 메서드 `GET`

포트 `8080`의 "my-service2" 서비스로 전달한다:
```yaml
{{< include file="examples/standard/basic-http.yaml" >}}
```

다음 예시는 `weight` 필드를 사용하여 `foo.example.com`으로 향하는 HTTP 요청의
90%를 "foo-v1" Service로, 나머지 10%를 "foo-v2" Service로 전달한다:
```yaml
{{< include file="examples/standard/traffic-splitting/traffic-split-2.yaml" >}}
```

`weight` 및 기타 필드에 대한 추가 정보는
[backendRef][backendRef] API 문서를 참조하라.

#### 타임아웃(Timeouts) (선택 사항)

{{< channel-version channel="standard" version="v1.2.0" >}}

HTTPRoute 타임아웃은 `v1.2.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하라.
{{< /channel-version >}}

HTTPRoute 규칙에는 `Timeouts` 필드가 포함된다. 지정되지 않으면 타임아웃 동작은 구현별로 다르다.

HTTPRoute 규칙에서 구성할 수 있는 타임아웃은 2가지 종류가 있다:

1. `request`는 Gateway API 구현이 클라이언트 HTTP 요청에 대한 응답을 보내는 타임아웃이다. 이 타임아웃은 전체 요청-응답 트랜잭션을 가능한 한 완전히 포괄하도록 의도되었지만, 구현은 트랜잭션이 클라이언트에 의해 시작된 직후가 아니라 전체 요청 스트림이 수신된 후 타임아웃을 시작하도록 선택할 수 있다(MAY).

2. `backendRequest`는 게이트웨이에서 백엔드로의 단일 요청에 대한 타임아웃이다. 이 타임아웃은 게이트웨이에서 요청이 처음 전송되기 시작한 시점부터 백엔드에서 전체 응답이 수신된 시점까지를 포괄한다. 이는 게이트웨이가 백엔드에 대한 연결을 재시도하는 경우 특히 유용할 수 있다.

`request` 타임아웃이 `backendRequest` 타임아웃을 포괄하므로, `backendRequest`의 값은 `request` 타임아웃의 값보다 클 수 없다.

타임아웃은 선택 사항이며, 해당 필드는 [Duration]({{< ref "/geps/gep-2257" >}}) 타입이다. 0값 타임아웃("0s")은 타임아웃을 비활성화하는 것으로 해석되어야 한다(MUST). 유효한 0이 아닌 타임아웃은 1ms 이상이어야 한다(MUST).

다음 예시는 클라이언트 요청이 완료되는 데 10초 이상 걸리면 타임아웃을 발생시키는 `request` 필드를 사용한다. 또한 게이트웨이에서 백엔드 서비스 `timeout-svc`로의 개별 요청에 대한 타임아웃을 지정하는 2초 `backendRequest`를 정의한다:

```yaml
{{< include file="examples/experimental/http-route-timeouts/timeout-example.yaml" >}}
```

추가 정보는 [timeouts][timeouts] API 문서를 참조하라.

#### 이름(Name) (선택 사항)

{{< channel-version channel="experimental" version="v1.2.0" >}}

이 개념은 `v1.2.0` 부터 실험 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하라.
{{< /channel-version >}}

HTTPRoute 규칙에는 선택적 `name` 필드가 포함된다. 라우트 규칙 이름의 활용은 구현별로 다르다. 이름은 다른 리소스에서 개별 라우트 규칙을 이름으로 참조하거나(예: 메타리소스의 `sectionName` 필드([GEP-2648]({{< ref "/geps/gep-2648" >}}#section-names))), 라우트 객체와 관련된 리소스의 상태 스탠자에서, 또는 HTTPRoute 규칙에서 구현이 생성하는 내부 구성 객체를 식별하는 데 사용할 수 있다.

지정하는 경우, name 필드의 값은 [`SectionName`](https://github.com/kubernetes-sigs/gateway-api/blob/v1.0.0/apis/v1/shared_types.go#L607-L624) 타입을 준수해야 한다.

다음 예시는 _읽기 전용_ 백엔드 서비스와 _쓰기 전용_ 백엔드 서비스 간의 트래픽 분배에 사용되는 HTTPRoute 규칙을 식별하기 위해 `name` 필드를 지정한다:

```yaml
{{< include file="examples/experimental/http-route-rule-name.yaml" >}}
```

##### 백엔드 프로토콜(Backend Protocol)

{{< channel-version channel="experimental" version="v1.0.0" >}}

이 개념은 `v1.0.0` 부터 실험 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하라.
{{< /channel-version >}}

일부 구현에서는 특정 프로토콜을 사용하여 트래픽을 라우팅하기 위해
[backendRef][backendRef]에 명시적으로 레이블을 지정해야 할 수 있다.
Kubernetes Service 백엔드의 경우 [`appProtocol`][appProtocol] 필드를
지정하여 이를 수행할 수 있다.


## 상태(Status)

상태(Status)는 HTTPRoute의 관찰된 상태를 정의한다.

### RouteStatus

RouteStatus는 모든 라우트 타입에서 필요한 관찰된 상태를 정의한다.

#### 부모(Parents)

부모(Parents)는 HTTPRoute와 연관된 게이트웨이(또는 기타 부모 리소스) 목록과
각 게이트웨이에 대한 HTTPRoute의 상태를 정의한다. HTTPRoute가 parentRefs에
게이트웨이에 대한 참조를 추가하면, 게이트웨이를 관리하는 컨트롤러는 라우트를
처음 발견했을 때 이 목록에 항목을 추가하고 라우트가 수정될 때
적절히 항목을 업데이트해야 한다.

다음 예시는 HTTPRoute "http-example"이 네임스페이스 "gw-example-ns"의
게이트웨이 "gw-example"에 의해 수락되었음을 나타낸다:
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: http-example
...
status:
  parents:
  - parentRef:
      name: gw-example
      namespace: gw-example-ns
    conditions:
    - type: Accepted
      status: "True"
```

## 병합(Merging) {#merging}
여러 HTTPRoute를 단일 게이트웨이 리소스에 연결할 수 있다. 중요한 점은,
각 요청에 대해 하나의 라우트 규칙만 매칭될 수 있다는 것이다. 병합에 대한
충돌 해결 방법에 대한 자세한 정보는 [API 사양][httprouterule]을 참조하라.


[httproute]: ../../reference/spec.md#httproute
[httprouterule]: ../../reference/spec.md#httprouterule
[hostname]: ../../reference/spec.md#hostname
[rfc-3986]: https://tools.ietf.org/html/rfc3986
[matches]: ../../reference/spec.md#httproutematch
[filters]: ../../reference/spec.md#httproutefilter
[backendRef]: ../../reference/spec.md#httpbackendref
[parentRef]: ../../reference/spec.md#parentreference
[timeouts]: ../../reference/spec.md#httproutetimeouts
[appProtocol]: https://kubernetes.io/docs/concepts/services-networking/service/#application-protocol
[sectionName]: ../../reference/spec.md#sectionname
