---
title: "TLS 구성"
linkTitle: "TLS"
weight: 11
description: "Configuring TLS for downstream and upstream connections"
---

# TLS 구성

Gateway API는 TLS를 구성하는 다양한 방법을 제공한다. 이 문서는
다양한 TLS 설정을 설명하고 이를 효과적으로 사용하기 위한 일반적인
가이드라인을 제공한다.

이 문서는 Gateway API를 사용하는 가장 일반적인 TLS 구성 형태를 다루지만,
일부 구현체는 다른 또는 더 고급 형태의 TLS 구성을 허용하는
구현체별 확장을 제공할 수도 있다. 이 문서 외에도,
Gateway API와 함께 사용하는 구현체의 TLS 문서를 읽어보는 것이 좋다.

## 클라이언트/서버와 TLS

![overview](/images/tls-overview.svg)

게이트웨이(Gateway)에는 두 가지 연결이 관련된다.

- **다운스트림**(downstream): 클라이언트와 게이트웨이 간의 연결이다.
- **업스트림**(upstream): 게이트웨이와 라우트에 의해 지정된 백엔드 리소스 간의 연결이다. 이러한 백엔드 리소스는 일반적으로 서비스(Service)이다.

Gateway API에서는 다운스트림과 업스트림 연결의 TLS 구성이 독립적으로 관리된다.

다운스트림 연결의 경우, 리스너 프로토콜(Listener Protocol)에 따라 다른 TLS 모드와 라우트 타입이 지원된다.

| 리스너 프로토콜 | TLS 모드    | 지원되는 라우트 타입 |
|-------------------|-------------|---------------------|
| TLS               | Passthrough | TLSRoute            |
| TLS               | Terminate   | TLSRoute (확장)     |
| HTTPS             | Terminate   | HTTPRoute           |
| GRPC              | Terminate   | GRPCRoute           |

`Passthrough` TLS 모드에서는 클라이언트의 TLS 세션이 게이트웨이에서 종료되지 않고
암호화된 상태로 게이트웨이를 통과하므로, TLS 설정이 적용되지 않는다.

업스트림 연결의 경우, `BackendTLSPolicy`가 사용되며, 리스너 프로토콜이나 TLS 모드는
업스트림 TLS 구성에 적용되지 않는다. `HTTPRoute`의 경우, `Terminate` TLS 모드와 `BackendTLSPolicy`를
함께 사용하는 것이 지원된다.
이 두 가지를 함께 사용하면 일반적으로 게이트웨이에서 종료된 후 재암호화되는 연결로 알려진 것을 제공한다.

{{< channel-version channel="standard" version="v1.5.0" >}}

`TLSRoute` 리소스는 GA(정식 출시)되었으며 `v1.5.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하자.
{{< /channel-version >}}

`TLSRoute`의 `Terminate` 모드는 `Extended`(확장) [지원 수준]에서 사용할 수 있다.

[지원 수준]: /concepts/conformance/#2-support-levels

## 다운스트림 TLS <a name="downstream-tls"></a>

다운스트림 TLS 설정은 게이트웨이 수준에서 리스너를 사용하여 구성된다.

### 리스너와 TLS

리스너는 도메인 또는 서브도메인 단위로 TLS 설정을 노출한다.
리스너의 TLS 설정은 `hostname` 기준을 충족하는 모든 도메인에 적용된다.

다음 예시에서, 게이트웨이는 모든 요청에 대해 `default-cert` Secret 리소스에
정의된 TLS 인증서를 제공한다.
이 예시는 HTTPS 프로토콜을 참조하지만, TLS 전용 프로토콜과 TLSRoute를 함께
사용하여 동일한 기능을 사용할 수도 있다.

```yaml
listeners:
- protocol: HTTPS # Other possible value is `TLS`
  port: 443
  tls:
    mode: Terminate # If protocol is `TLS`, `Passthrough` is another possible mode
    certificateRefs:
    - kind: Secret
      group: ""
      name: default-cert
```

### 예시

#### 서로 다른 인증서를 가진 리스너

이 예시에서, 게이트웨이는 `foo.example.com`과 `bar.example.com` 도메인을
제공하도록 구성되어 있다. 이 도메인들의 인증서는 게이트웨이에서 지정된다.

{{< include file="examples/standard/tls-basic.yaml" >}}

#### 와일드카드 TLS 리스너 <a name="wildcard-tls-listeners"></a>

이 예시에서, 게이트웨이는 `*.example.com`에 대한 와일드카드 인증서와
`foo.example.com`에 대한 다른 인증서로 구성되어 있다.
구체적인 일치가 우선하므로, 게이트웨이는 `foo.example.com`에 대한 요청에는
`foo-example-com-cert`를, 다른 모든 요청에는
`wildcard-example-com-cert`를 제공한다.

{{< include file="examples/standard/wildcard-tls-gateway.yaml" >}}

#### 네임스페이스 간 인증서 참조

이 예시에서, 게이트웨이는 다른 네임스페이스의 인증서를 참조하도록 구성되어 있다.
이는 대상 네임스페이스에 생성된 ReferenceGrant에 의해 허용된다.
해당 ReferenceGrant가 없으면, 네임스페이스 간 참조는 유효하지 않다.

{{< include file="examples/standard/tls-cert-cross-namespace.yaml" >}}
### 클라이언트 인증서 검증 (프론트엔드 mTLS)
{{< channel-version channel="standard" version="v1.5.0" >}}
GatewayFrontendClientCertificateValidation 기능은 `v1.5.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하자.
{{< /channel-version >}}

Gateway API는 TLS 핸드셰이크(TLS handshake) 중에 프론트엔드 클라이언트가 게이트웨이에 제출하는 TLS 인증서를 검증하는 것을 지원한다.

서버 인증서 구성이 리스너별로 정의되는 것과 달리, 클라이언트 인증서 검증은 `spec.tls` 필드 내에서 **게이트웨이 수준**으로 구성된다. 이 설계는 HTTP/2 및 TLS 연결 병합(connection coalescing)과 관련된 보안 위험을 완화하기 위한 것으로, 하나의 리스너에 대해 설정된 연결이 동일한 포트의 다른 리스너에 재사용되어 리스너별 검증 설정을 우회할 수 있는 상황을 방지한다.

#### 구성 개요
클라이언트 검증은 게이트웨이가 클라이언트의 신원을 검증하는 방법을 지정하는 `frontendValidation` 구조체를 사용하여 정의된다.

*   **`caCertificateRefs`**: 클라이언트 인증서를 검증하기 위한 트러스트 앵커(trust anchor)로 사용되는 PEM 인코딩된 CA 인증서 번들을 포함하는 쿠버네티스 객체(일반적으로 `ConfigMap`)에 대한 참조 목록이다.
*   **`mode`**: 검증 동작을 정의한다.
    *   `AllowValidOnly` (기본값): 게이트웨이는 클라이언트가 지정된 CA 번들에 대한 검증을 통과하는 유효한 인증서를 제출하는 경우에만 연결을 수락한다.
    *   `AllowInsecureFallback`: 게이트웨이는 클라이언트 인증서가 누락되었거나 검증에 실패한 경우에도 연결을 수락한다. 이 모드는 일반적으로 인가를 백엔드에 위임하며 주의하여 사용해야 한다.

#### 게이트웨이 수준 범위 지정

검증은 게이트웨이 전체에 적용하거나 특정 포트에 대해 재정의할 수 있다.

1.  **기본 구성**: 이 구성은 포트별 재정의가 정의되지 않는 한, 게이트웨이의 모든 HTTPS 리스너에 적용된다.
2.  **포트별 구성**: 특정 포트에서 트래픽을 처리하는 모든 리스너에 대해 기본 구성을 재정의하여 세밀한 제어를 할 수 있다.

#### 예시

##### 기본 클라이언트 검증
이 예시는 기본 구성과 포트별 재정의를 사용하여 클라이언트 인증서 검증을 구성하는 방법을 보여준다.

{{< include file="examples/standard/frontend-cert-validation.yaml" >}}

## 업스트림 TLS

업스트림 TLS 설정은 대상 참조를 통해 `Service`에 연결된 `BackendTLSPolicy`를
사용하여 구성된다.

이 리소스는 게이트웨이가 백엔드에 연결할 때 사용해야 하는 SNI와
백엔드 파드(Pod)가 제공하는 인증서를 검증하는 방법을 설명하는 데 사용할 수 있다.

### TargetRef와 TLS

BackendTLSPolicy는 `TargetRefs`와 `Validation`에 대한 사양을 포함한다. TargetRef는 필수이며
HTTPRoute가 TLS를 요구하는 하나 이상의 `Service`를 식별한다. `Validation` 구성에는
필수 `Hostname`과 `CACertificateRefs` 또는 `WellKnownCACertificates` 중 하나가 포함된다.

`Hostname`은 게이트웨이가 백엔드에 연결할 때 사용해야 하는 SNI를 나타내며,
백엔드 파드가 제공하는 인증서와 일치해야 한다.

CACertificateRefs는 하나 이상의 PEM 인코딩된 TLS 인증서를 참조한다. 사용할 특정 인증서가 없는 경우,
WellKnownCACertificates를 "System"으로 설정하여 게이트웨이가 신뢰할 수 있는
CA 인증서 세트를 사용하도록 지시해야 한다. 각 구현체에서 사용되는 시스템 인증서에는
다소 차이가 있을 수 있다.
자세한 정보는 선택한 구현체의 문서를 참조하자.

{{< info >}}
**제한 사항**


- 네임스페이스 간 인증서 참조는 허용되지 않는다.
- 와일드카드 호스트네임은 허용되지 않는다.
{{< /info >}}

### 예시

#### 시스템 인증서 사용

이 예시에서, `BackendTLSPolicy`는 `dev` 서비스를 지원하는 파드가
`dev.example.com`에 대한 유효한 인증서를 제공하는 TLS 암호화 업스트림
연결에 시스템 인증서를 사용하도록 구성되어 있다.

{{< include file="examples/standard/backendtlspolicy/backendtlspolicy-system-certs.yaml" >}}

#### 명시적 CA 인증서 사용

이 예시에서, `BackendTLSPolicy`는 `auth` 서비스를 지원하는 파드가
`auth.example.com`에 대한 유효한 인증서를 제공하는 TLS 암호화 업스트림
연결에 구성 맵 `auth-cert`에 정의된 인증서를 사용하도록 구성되어 있다.

{{< include file="examples/standard/backendtlspolicy/backendtlspolicy-ca-certs.yaml" >}}
### 게이트웨이의 인증서 선택 (백엔드 mTLS)
{{< channel-version channel="standard" version="v1.5.0" >}}
GatewayBackendClientCertificate 기능은 `v1.5.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하자.
{{< /channel-version >}}

업스트림 연결을 위한 상호 TLS(mTLS, Mutual TLS)는 백엔드의 인증서를 검증하는 것 외에도 게이트웨이가 백엔드에 클라이언트 인증서를 제출할 것을 요구한다. 이를 통해 백엔드는 인가된 게이트웨이의 연결만 수락할 수 있다.

#### 게이트웨이의 클라이언트 인증서 구성
게이트웨이가 백엔드에 연결할 때 사용하는 클라이언트 인증서를 구성하려면, `Gateway` 리소스의 `tls.backend.clientCertificateRef` 필드를 사용한다.

이 구성은 해당 게이트웨이에서 관리하는 *모든* 업스트림 연결에 대해 클라이언트로서의 게이트웨이에 적용된다.

{{< include file="examples/standard/backend-tls.yaml" >}}

## 확장

게이트웨이 TLS 구성은 구현체별 기능을 위한 추가 TLS 설정을 추가할 수 있는
`options` 맵을 제공한다. 여기에 포함될 수 있는 기능의 예로는
TLS 버전 제한이나 사용할 암호화 스위트(cipher suite) 등이 있다.
