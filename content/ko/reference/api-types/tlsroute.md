---
title: "TLSRoute"
weight: 70
description: "Routing TLS requests based on SNI hostname"
---

{{< channel-version channel="standard" version="v1.5.0" >}}

`TLSRoute` 리소스는 GA(정식 출시)되었으며 `v1.5.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하자.
{{< /channel-version >}}

[TLSRoute(TLS 라우트)][tlsroute]는 클라이언트에서 API 객체(예: Service)로의
TLS 요청 라우팅 동작을 지정하기 위한 Gateway API 타입이다. TLS 핸드셰이크 중에
제공되는 [SNI(서버 이름 표시)](https://datatracker.ietf.org/doc/html/rfc6066#section-3)
호스트네임을 기반으로 특정 백엔드로 트래픽을 라우팅할 수 있다.

이 기능은 흔히 "TLS 패스스루"라고 불리며, 게이트웨이가 SNI를 통해 서버 이름을
식별하고 완전히 암호화된 상태로 통신을 백엔드에 직접 전달한다. 그러나 TLSRoute는
게이트웨이에서 트래픽을 종료한 후 암호화되지 않은 상태로 백엔드에 전달하는 것도
허용한다.

TLSRoute 지원은 다음 기능으로 표현되며, 구현체에서 보고할 수 있다.

- `TLSRoute` - 보고되는 경우, 구현체가 `Passthrough` 모드의 TLSRoute를
  지원한다. TLSRoute API를 지원한다고 주장하는 모든 구현체는 이 기능을 반드시
  보고해야 한다.
- `TLSRouteModeTerminate` - 보고되는 경우, 구현체가 `Passthrough` 모드 외에
  `Terminate` 모드의 TLSRoute를 지원한다.
- `TLSRouteModeMixed` - 보고되는 경우, 구현체가 동일한 포트에서 서로 다른
  모드(`Passthrough`와 `Terminate`)를 가진 두 개의 TLS 리스너를 지원한다.

## 배경

많은 애플리케이션 라우팅 사례가 HTTP/L7 매칭(프로토콜:호스트네임:포트:경로 튜플)을
사용하여 구현될 수 있지만, TLS를 종료하지 않고 백엔드에 직접 암호화된 통신이
필요한 특정 사례가 있다. 일반적인 예시로는 다음과 같다.

* TLS 기반이지만 HTTP 기반이 아닌 백엔드(예: Kafka 서비스 또는 TLS가 활성화된
리스너를 가진 Postgres 서비스).
* 특정 WebRTC 솔루션.
* 클라이언트 인증서를 사용한 mTLS(상호 TLS) 인증이 필요한 백엔드.

이러한 시나리오에서는 패스스루 모드를 사용하는 것이 바람직하며, 게이트웨이가
TLS 연결을 종료하지 않고 암호화된 패킷을 백엔드에 직접 전달한다.

다른 경우에는 게이트웨이에서 TLS를 종료하고 암호화되지 않은 패킷을 기본 TCP
연결로 백엔드에 전달하고 싶을 수 있다(종료 모드).

TLSRoute는 클라이언트와 게이트웨이 간의 트래픽이 암호화되어 있고 SNI 호스트네임을
포함하는 경우에 사용할 수 있으며, 이 호스트네임을 사용하여 요청에 사용할 백엔드를
결정할 수 있다.

## 사양

TLSRoute의 사양은 다음으로 구성된다.

- [ParentRefs][parentRef] - 이 라우트가 연결되고자 하는 Gateway를 정의한다.
- [Hostnames][hostname] - TLS 핸드셰이크의 SNI 호스트네임과 매칭하기 위한
  호스트네임 목록을 정의한다.
- [Rules][tlsrouterule] - 매칭되는 TLS 핸드셰이크에 대한 작업 규칙 목록을
  정의한다. TLSRoute의 경우 사용할 [backendRefs][backendRef]로 제한된다.

### Gateway에 연결하기

각 라우트에는 연결하고자 하는 부모 리소스를 참조하는 방법이 포함되어 있다.
대부분의 경우 Gateway가 되지만, 구현에서 다른 타입의 부모 리소스를
지원하는 경우도 있다.

다음 예시는 라우트가 `acme-lb` Gateway에 연결하는 방법을 보여준다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: TLSRoute
metadata:
  name: tlsroute-example
spec:
  parentRefs:
  - name: acme-lb
```

대상 Gateway는 라우트의 네임스페이스에서 TLSRoute의 연결을 허용해야
연결이 성공한다는 점에 유의하자.

TLS 리스너의 경우, `tls.mode` 필드를 정의하는 것이 필수이다. 이 필드는
두 가지 값을 허용한다.

- Passthrough - 트래픽이 암호화된 상태를 유지하면서 백엔드로 전달된다.
- Terminate - 암호화된 트래픽이 게이트웨이에서 종료되고, 암호화되지 않은
  TCP 패킷이 하나 이상의 백엔드로 전달된다.

부모 리소스의 특정 섹션에 라우트를 연결할 수도 있다.
예를 들어, `acme-lb` Gateway에 다음과 같은 리스너가 포함되어 있다고 가정하자.

```yaml
  listeners:
  - name: passthrough
    protocol: TLS
    port: 8883
    tls:
      mode: Passthrough
    ...
  - name: terminated
    protocol: TLS
    port: 18883
    tls:
      mode: Terminate
    ...
```

`parentRefs`의 `sectionName` 필드를 사용하여 `Passthrough` 리스너에만
라우트를 바인딩할 수 있다.

```yaml
spec:
  parentRefs:
  - name: acme-lb
    sectionName: passthrough
```

또는 `parentRefs`에서 `sectionName` 대신 `port` 필드를 사용하여
같은 효과를 얻을 수 있다.

```yaml
spec:
  parentRefs:
  - name: acme-lb
    port: 8883
```

포트에 바인딩하면 한 번에 여러 리스너에 연결할 수도 있다.
예를 들어, `acme-lb` Gateway의 포트 `8090`에 바인딩하는 것이
이름으로 해당 리스너에 바인딩하는 것보다 더 편리할 수 있다.

```yaml
spec:
  parentRefs:
  - name: acme-lb
    sectionName: bar
  - name: acme-lb
    sectionName: baz
```

그러나 포트 번호로 라우트를 바인딩할 때, Gateway 관리자는 라우트를 함께
업데이트하지 않고는 Gateway의 포트를 변경할 수 있는 유연성을 잃게 된다.
이 접근 방식은 포트가 변경될 수 있는 리스너가 아닌 특정 포트 번호에
라우트를 바인딩해야 하는 경우에만 사용해야 한다.

### 호스트네임(Hostnames)

호스트네임은 TLS 요청의 SNI 호스트네임과 매칭할 호스트네임 목록을 정의한다.
매칭이 발생하면 TLSRoute가 선택되어 규칙에 따라 요청을 라우팅한다.

SNI 사양은 호스트네임 정의에 다음과 같은 제한을 추가한다.

- 호스트네임은 반드시 완전한 도메인 이름(FQDN)이어야 한다.
- IPv4 및 IPv6 주소의 사용은 허용되지 않는다.

다음 예시는 호스트네임 "my.example.com"을 정의한다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: TLSRoute
metadata:
  name: tlsroute-example
spec:
  hostnames:
  - my.example.com
```

### 규칙(Rules)

규칙은 트래픽에 대해 수행할 작업 목록을 정의한다.

#### BackendRefs

BackendRefs는 매칭된 요청이 전송되어야 하는 API 객체를 정의한다. 최소한
하나의 backendRef가 지정되어야 한다.

다음 예시는 호스트네임이 `foo.example.com`인 TLS 요청을 포트 `443`의
"foo-svc" 서비스로 전달한다.

{{< include file="examples/standard/tls-routing/tls-route.yaml" >}}

`weight` 및 기타 필드에 대한 추가 정보는
[backendRef][backendRef] API 문서를 참조하자.

이 TLSRoute는 아래에 정의된 Gateway TLS 리스너 `tls`에 연결된다.

{{< include file="examples/standard/tls-routing/gateway.yaml" >}}

`tls` 리스너의 TLS 모드가 `Passthrough`로 구성되어 있으므로, 이 리스너를 통해
라우팅되는 트래픽은 완전히 암호화된 TCP 스트림으로 백엔드에 전송된다.

대신 `tls-terminate` 리스너가 사용된 경우, TLS 트래픽은 게이트웨이에서
종료되고 결과 TCP 스트림은 암호화되지 않은 상태로 백엔드에 전달된다.

## 상태(Status)

상태(Status)는 TLSRoute의 관찰된 상태를 정의한다.

### RouteStatus

RouteStatus는 모든 라우트 타입에서 필요한 관찰된 상태를 정의한다.

#### 부모(Parents)

부모(Parents)는 TLSRoute와 연관된 Gateway(또는 기타 부모 리소스) 목록과
각 Gateway에 대한 TLSRoute의 상태를 정의한다. TLSRoute가 parentRefs에
Gateway에 대한 참조를 추가하면, Gateway를 관리하는 컨트롤러는 라우트를
처음 발견했을 때 이 목록에 항목을 추가하고 라우트가 수정될 때
적절히 항목을 업데이트해야 한다.

다음 예시는 TLSRoute "tls-example"이 네임스페이스 "gw-example-ns"의
Gateway "gw-example"에 의해 수락되었음을 나타낸다.
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: TLSRoute
metadata:
  name: tls-example
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

## 병합(Merging)
여러 TLSRoute를 단일 Gateway 리소스에 연결할 수 있다. 중요한 점은,
각 요청에 대해 하나의 라우트 호스트네임만 매칭될 수 있다는 것이다. 병합에 대한
충돌 해결 방법에 대한 자세한 정보는 [API 사양][hostname]을 참조하자.


[tlsroute]: ../../reference/spec.md#tlsroute
[tlsrouterule]: ../../reference/spec.md#tlsrouterouterule
[hostname]: ../../reference/spec.md#hostname
[backendRef]: ../../reference/spec.md#backendref
[parentRef]: ../../reference/spec.md#parentreference
[name]: ../../reference/spec.md#sectionname
[rfc-6066]: https://tools.ietf.org/html/rfc6066
