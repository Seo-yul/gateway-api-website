---
title: "BackendTLSPolicy"
weight: 10
description: "Configuring TLS connections from Gateway to backend pods"
aliases:
  - /ko/reference/api-types/backendtlspolicy/
---

{{< channel-version channel="standard" version="v1.4.0" >}}

`BackendTLSPolicy` 리소스는 GA(정식 출시)되었으며 `v1.4.0` 부터
표준 채널의 일부이다. 릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하자.
{{< /channel-version >}}

[BackendTLSPolicy][backendtlspolicy](백엔드 TLS 정책)는 Service API 오브젝트를 통해
게이트웨이에서 백엔드 파드로의 TLS(전송 계층 보안) 연결 구성을 지정하기 위한 Gateway API 타입이다.

구현체는 HTTPRoute를 통해 연결되지 않은 서비스나
[InferencePool](https://gateway-api-inference-extension.sigs.k8s.io/reference/spec/#inferencepool)과 같은
다른 종류의 백엔드에 대한 TLS 연결 구성을 지정하기 위해 `BackendTLSPolicy` 사용을 지원할 수도 있다.

## 배경

`BackendTLSPolicy`는 게이트웨이 데이터플레인에서 백엔드로 HTTPS를 전달하기 위한 TLS 구성을
특별히 다룬다. 이를 "백엔드 TLS 종료(backend TLS termination)"라고 하며, 게이트웨이가
자체 인증서를 가진 백엔드 파드에 연결하는 방법을 알 수 있게 해준다.

**패스스루**(passthrough) 및 **엣지**(edge) 종료를 위한 TLS 구성에 사용되는 다른 API 오브젝트가 있지만,
이 API 오브젝트는 사용자가 **백엔드**(backend) TLS 종료를 구체적으로 구성할 수 있게 해준다.
게이트웨이 API에서의 TLS 구성에 대한 자세한 정보는 [TLS 구성]({{< ref "/guides/user-guides/tls" >}})을 참조하자.

![세 가지 TLS 종료 타입을 보여주는 이미지](/images/tls-termination-types.png)

백엔드 TLS 정책은 기본값이나 재정의 없이 백엔드에 접근하는 Service에 적용되는
직접 [PolicyAttachment(정책 연결)]({{< ref "/reference/policy-attachment" >}})이다.
백엔드 TLS 정책은 적용되는 Service와 동일한 네임스페이스에 위치해야 한다.
백엔드 TLS 정책과 Service는 네임스페이스 경계를 넘어 신뢰를 공유하는 복잡성을 방지하기 위해
동일한 네임스페이스에 위치해야 한다.

또한, 구현체는 HTTPRoute를 통해 연결되지 않은 서비스나 다른 종류의 백엔드에 대한
게이트웨이에서의 TLS 연결 구성을 지정하기 위해 백엔드 TLS 정책을 사용할 수 있지만,
이 동작은 선택 사항이며 모든 구현체에서 사용 가능하지 않을 수 있다.

참조된 Service를 가리키는 모든 Gateway API 라우트는 구성된 백엔드 TLS 정책을 준수해야 한다.

## 사양

[BackendTLSPolicy][backendtlspolicy](백엔드 TLS 정책)의 사양은 다음으로 구성된다:

- [TargetRefs][targetRefs] - 정책의 대상 API 오브젝트를 정의한다.
- [Validation][validation] - 호스트명, CACertificateRefs, WellKnownCACertificates를 포함한 TLS 구성을 정의한다.
- [Hostname][hostname] - 게이트웨이가 백엔드에 연결하는 데 사용하는 SNI(Server Name Indication, 서버 이름 표시)를 정의한다.
- [SubjectAltNames][subjectAltNames] - 백엔드 인증서가 일치해야 하는 하나 이상의 주체 대체 이름(Subject Alternative Name)을 지정한다. 지정된 경우, 인증서에는 일치하는 SAN이 하나 이상 있어야 한다. 이 필드는 SNI(호스트명)와 인증서 ID 검증을 분리할 수 있게 해준다. 최대 5개의 SAN이 허용된다.
- [CACertificateRefs][caCertificateRefs] - PEM 인코딩된 TLS 인증서를 포함하는 오브젝트에 대한 하나 이상의 참조를 정의하며, 게이트웨이와 백엔드 파드 간의 TLS 핸드셰이크를 수립하는 데 사용된다. CACertificateRefs 또는 WellKnownCACertificates 중 하나만 지정할 수 있으며, 둘 다 지정할 수는 없다.
- [WellKnownCACertificates][wellKnownCACertificates] - 게이트웨이와 백엔드 파드 간의 TLS 핸드셰이크에서 시스템 CA 인증서를 사용할 수 있는지 여부를 지정한다. CACertificateRefs 또는 WellKnownCACertificates 중 하나만 지정할 수 있으며, 둘 다 지정할 수는 없다.
- [Options][options] - 지원을 제공하기로 선택한 구현체를 위한 확장 TLS 구성을 가능하게 하는 키/값 쌍의 맵이다. 자세한 내용은 구현체의 문서를 확인하자.

다음 차트는 오브젝트 정의와 관계를 보여준다:
{{< mermaid >}}
flowchart LR
    backendTLSPolicy[["<b>backendTLSPolicy</b> <hr><align=left>BackendTLSPolicySpec: spec<br>PolicyStatus: status</align>"]]
    spec[["<b>spec</b><hr>PolicyTargetReferenceWithSectionName: targetRefs <br> BackendTLSPolicyValidation: tls"]]
    status[["<b>status</b><hr>[ ]PolicyAncestorStatus: ancestors"]]
    validation[["<b>tls</b><hr>LocalObjectReference: caCertificateRefs<br>wellKnownCACertificatesType: wellKnownCACertificates<br>PreciseHostname: hostname<br>[]SubjectAltName: subjectAltNames"]]
    ancestorStatus[["<b>ancestors</b><hr>AncestorRef: parentReference<br>GatewayController: controllerName<br>[]Condition: conditions"]]
    targetRefs[[<b>targetRefs</b><hr>]]
    service["<b>service</>"]
    backendTLSPolicy -->spec
    backendTLSPolicy -->status
    spec -->targetRefs & validation
    status -->ancestorStatus
    targetRefs -->service
    note[<em>choose only one<hr> caCertificateRefs OR wellKnownCACertificates</em>]
    style note fill:#fff
    validation -.- note
{{< /mermaid >}}

다음은 백엔드를 서빙하는 Service에 대해 TLS를 구성하는 백엔드 TLS 정책을 보여준다:
{{< mermaid >}}
flowchart LR
    client(["client"])
    gateway["Gateway"]
    httproute["HTTP<BR>Route"]
    service["Service"]
    pod1["Pod"]
    pod2["Pod"]
    client -.->|HTTP <br> request| gateway
    gateway --> httproute
    httproute -.->|BackendTLSPolicy|service
    service --> pod1 & pod2
{{< /mermaid >}}

### 백엔드 대상 지정

백엔드 TLS 정책은 Service에 대한 하나 이상의 TargetRefs를 통해 백엔드 파드(또는 파드 집합)를 대상으로 한다.
이 TargetRef는 이름(Name), 종류(Kind: Service), 그리고 선택적으로 네임스페이스(Namespace)와
그룹(Group)으로 Service를 지정하는 필수 오브젝트 참조이다.
TargetRefs는 HTTPRoute에 TLS가 필요한 Service를 식별한다.

{{< note >}}
**제한 사항**


- 교차 네임스페이스 인증서 참조는 허용되지 않는다.
{{< /note >}}

### BackendTLSPolicyValidation

BackendTLSPolicyValidation은 백엔드 TLS 정책의 사양이며,
호스트명(서버 이름 표시용)과 인증서를 포함한 TLS 구성을 정의한다.

#### 호스트명

호스트명은 게이트웨이가 백엔드에 연결하기 위해 사용해야 하는 SNI(서버 이름 표시)를 정의하며,
백엔드 파드가 제공하는 인증서와 일치해야 한다. 호스트명은 [RFC 3986][rfc-3986]에 정의된
네트워크 호스트의 완전한 도메인 이름(FQDN)이다. RFC에서 정의된 URI의 "host" 부분에서
다음과 같은 차이점에 주의하자:

- IP 주소는 허용되지 않는다.

또한 다음 사항에 주의하자:

{{< note >}}
**제한 사항**


- 와일드카드 호스트명은 허용되지 않는다.
{{< /note >}}

#### 주체 대체 이름 (Subject Alternative Names)

{{< channel-version channel="experimental" version="v1.2.0" >}}

이 필드는 `v1.2.0`에서 BackendTLSPolicy에 추가되었다.
{{< /channel-version >}}
subjectAltNames 필드는 게이트웨이와 백엔드 간의 기본적인 mTLS(상호 TLS) 구성과 선택적으로 SPIFFE 사용을 가능하게 한다. subjectAltNames가 지정되면, 백엔드가 제공하는 인증서에는 지정된 값 중 하나와 일치하는 주체 대체 이름이 하나 이상 있어야 한다. 이는 URI 기반 SAN이 유효한 SNI가 아닐 수 있는 SPIFFE 구현에서 특히 유용하다.
주체 대체 이름은 Hostname 또는 URI 필드 중 하나를 포함할 수 있으며, Hostname 또는 URI 중 어느 것을 선택했는지 지정하는 Type을 포함해야 한다.

{{< note >}}
**제한 사항**


- IP 주소와 와일드카드 호스트명은 허용되지 않는다. (자세한 내용은 위의 호스트명 설명을 참조하자).
- Hostname: DNS 이름 형식
- URI: URI 형식 (예: SPIFFE ID)
{{< /note >}}

#### TLS 옵션

{{< channel-version channel="experimental" version="v1.2.0" >}}

이 필드는 `v1.2.0`에서 BackendTLSPolicy에 추가되었다.
{{< /channel-version >}}
options 필드는 구현체별 TLS 구성을 지정할 수 있게 해준다. 이에는 다음이 포함될 수 있다:

- 벤더별 상호 TLS 자동화 구성
- 최소 지원 TLS 버전 제한
- 지원되는 암호 스위트 구성

자세한 내용은 구현체의 문서를 확인하자.

###
#### 인증서

BackendTLSPolicyValidation은 어떤 형태의 인증서 참조를 포함해야 하며,
백엔드 TLS에 사용할 인증서를 구성하는 두 가지 방법인 CACertificateRefs와
WellKnownCACertificates를 포함한다. BackendTLSPolicyValidation당 이 중 하나만 사용할 수 있다.

##### CACertificateRefs

CACertificateRefs는 하나 이상의 PEM 인코딩된 TLS 인증서를 참조한다.

{{< note >}}
**제한 사항**


- 교차 네임스페이스 인증서 참조는 허용되지 않는다.
{{< /note >}}

##### WellKnownCACertificates

특정 TLS 인증서가 필요하지 않은 환경에서 작업하고 있으며, Gateway API 구현체가
시스템 또는 기본 인증서 사용을 허용하는 경우(예: 개발 환경),
WellKnownCACertificates를 "System"으로 설정하여 게이트웨이가 신뢰할 수 있는 CA 인증서 세트를
사용하도록 지시할 수 있다. 각 구현체가 사용하는 시스템 인증서에는 약간의 차이가 있을 수 있다.
자세한 정보는 선택한 구현체의 문서를 참조하자.

### 정책 상태 (PolicyStatus)

Status는 백엔드 TLS 정책의 관찰된 상태를 정의하며 사용자가 구성할 수 없다.
올바른 작동을 확인하기 위해 다른 Gateway API 오브젝트와 동일한 방식으로 상태를 확인하자.
백엔드 TLS 정책의 상태는 `PolicyAncestorStatus`를 사용하여 어떤 parentReference가
해당 특정 상태를 설정했는지 알 수 있게 해준다.

[backendtlspolicy]: ../../reference/spec.md#backendtlspolicy
[validation]: ../../reference/spec.md#backendtlspolicyvalidation
[caCertificateRefs]: ../../reference/spec.md#localobjectreference
[wellKnownCACertificates]: ../../reference/spec.md#localobjectreference
[hostname]: ../../reference/spec.md#precisehostname
[rfc-3986]: https://tools.ietf.org/html/rfc3986
[targetRefs]: ../../reference/spec.md#localpolicytargetreferencewithsectionname
[subjectAltNames]: ../../reference/spec.md#subjectaltname
[options]: ../../reference/spec.md#backendtlspolicyspec
