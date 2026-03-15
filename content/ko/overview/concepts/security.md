---
title: "보안"
weight: 50
description: "RBAC, 네임스페이스 간 참조 및 모범 사례를 포함한 Gateway API 보안 모델"
---

## 소개
Gateway API(게이트웨이 API)는 일반적인 조직의 각 역할에 대해
세분화된 권한 부여를 가능하게 하도록 설계되었다.

## 리소스
게이트웨이 API에는 3가지 주요 API 리소스가 있다.

* **GatewayClass**(게이트웨이 클래스)는 공통 구성과 동작을 가진 게이트웨이 집합을
  정의한다.
* **Gateway**(게이트웨이)는 트래픽이 클러스터 내의 서비스로 변환될 수 있는 지점을
  요청한다.
* **Route**(라우트)는 게이트웨이를 통해 들어오는 트래픽이 서비스에 어떻게 매핑되는지 설명한다.

## 역할과 페르소나

[역할과 페르소나]에 설명된 대로 게이트웨이 API에는 3가지 주요 역할이 있다.

- **Ian** (그/그의): 인프라 제공자
- **Chihiro** (그들/그들의): 클러스터 운영자
- **Ana** (그녀/그녀의): 애플리케이션 개발자

[역할과 페르소나]: {{< ref "/overview/concepts/roles-and-personas" >}}

### RBAC

RBAC(역할 기반 접근 제어)은 쿠버네티스에서 권한 부여를 위해 사용되는 표준 방식이다.
이를 통해 사용자는 특정 범위에서 누가 어떤 리소스에 대해 작업을 수행할 수 있는지를
설정할 수 있다. RBAC를 사용하여 위에서 정의한 각 역할을 활성화할 수 있다.
대부분의 경우 모든 역할이 대부분의 리소스를 읽을 수 있도록 하는 것이 바람직하므로,
이 모델에서는 쓰기 권한에 초점을 맞추겠다.

#### 간단한 3계층 모델의 쓰기 권한
| | 게이트웨이 클래스 | 게이트웨이 | 라우트 |
|-|-|-|-|
| 인프라 제공자 | 예 | 예 | 예 |
| 클러스터 운영자 | 아니오 | 예 | 예 |
| 애플리케이션 개발자 | 아니오 | 아니오 | 예 |

#### 고급 4계층 모델의 쓰기 권한
| | 게이트웨이 클래스 | 게이트웨이 | 라우트 |
|-|-|-|-|
| 인프라 제공자 | 예 | 예 | 예 |
| 클러스터 운영자 | 때때로 | 예 | 예 |
| 애플리케이션 관리자 | 아니오 | 지정된 네임스페이스에서 | 지정된 네임스페이스에서 |
| 애플리케이션 개발자 | 아니오 | 아니오 | 지정된 네임스페이스에서 |

## 네임스페이스 경계 넘기

게이트웨이 API는 네임스페이스 경계를 넘는 새로운 방법을 제공한다. 이러한
네임스페이스 간 기능은 매우 강력하지만 의도하지 않은 노출을 방지하기 위해
신중하게 사용해야 한다. 원칙적으로 네임스페이스 경계를 넘는 것을 허용할 때마다
네임스페이스 간의 상호 합의가 필요하다. 이는 두 가지 방법으로 이루어질 수 있다.

### 1. 라우트 바인딩
라우트는 다른 네임스페이스에 있는 게이트웨이에 연결될 수 있다. 이를 위해
게이트웨이 소유자는 추가 네임스페이스에서 라우트가 바인딩되는 것을 명시적으로
허용해야 한다. 이는 게이트웨이 리스너 내에서 allowedRoutes를 다음과 같이
구성하여 수행된다.

```yaml
namespaces:
  from: Selector
  selector:
    matchExpressions:
    - key: kubernetes.io/metadata.name
      operator: In
      values:
      - foo
      - bar
```

이렇게 하면 "foo" 및 "bar" 네임스페이스의 라우트가 이 게이트웨이 리스너에
연결될 수 있다.

#### 다른 레이블의 위험성
이 셀렉터에 다른 레이블을 사용하는 것도 가능하지만, 그만큼 안전하지는 않다.
`kubernetes.io/metadata.name` 레이블은 네임스페이스의 이름으로 일관되게 설정되지만,
다른 레이블에는 동일한 보장이 없다. `env`와 같은 사용자 정의 레이블을 사용하면
클러스터 내에서 네임스페이스에 레이블을 지정할 수 있는 누구나 게이트웨이가
지원하는 네임스페이스 집합을 효과적으로 변경할 수 있게 된다.

### 2. ReferenceGrant
다른 객체 참조가 네임스페이스 경계를 넘는 것을 허용하는 경우가 있다.
여기에는 게이트웨이가 Secret(시크릿)을 참조하는 경우와 라우트가 Backend(백엔드,
일반적으로 서비스)를 참조하는 경우가 포함된다. 이러한 경우에 필요한 상호 합의는
ReferenceGrant 리소스를 통해 이루어진다. 이 리소스는 대상 네임스페이스 내에
존재하며 다른 네임스페이스에서의 참조를 허용하는 데 사용할 수 있다.

예를 들어, 다음 ReferenceGrant(레퍼런스그랜트)는 "prod" 네임스페이스의
HTTPRoute에서 ReferenceGrant와 동일한 네임스페이스에 배포된 서비스로의 참조를
허용한다.

{{< include file="examples/standard/reference-grant.yaml" >}}

ReferenceGrant에 대한 자세한 내용은 [이 리소스에 대한 상세 문서]({{< ref "/reference/api-types/referencegrant" >}})를
참조하자.

## 보안 고려 사항

게이트웨이 컨트롤러는 서로 다른 네임스페이스가 서로 다른 사용자와 고객에 의해
사용되는 멀티 테넌트 환경에서 배포될 수 있다.

보다 안전한 환경을 제공하기 위해 클러스터 관리자와 게이트웨이 소유자는
몇 가지 주의를 기울여야 한다.

### 호스트네임/도메인 하이재킹 방지

게이트웨이 API에서는 서로 다른 라우트와 ListenerSet(리스너셋)이 동일한
호스트네임을 요청할 수 있다. 게이트웨이 컨트롤러는 충돌 해결을 담당하며,
일반적으로 먼저 생성된 리소스가 충돌 관리에서 우선하는 선착순 방식으로 작동한다.

라우트 또는 리스너셋의 [호스트네임 정의]({{< ref "/reference/spec#httproutespec" >}})는
목록이며, 다음 시나리오를 고려하자.

* `Gateway`가 특정 네임스페이스 집합으로부터 라우트를 수용한다.
* `HTTPRoute` 이름이 `route1`인 라우트가 네임스페이스 `ns1`에 `creationTimestamp: 00:00:01`로
생성되고 호스트네임 `something.tld`를 정의한다.
* `HTTPRoute` 이름이 `route2`인 라우트가 네임스페이스 `ns2`에 `creationTimestamp: 00:00:30`으로
생성되고 호스트네임 `otherthing.tld`를 정의한다.

만약 `route1`의 소유자가 나중에 호스트네임 목록에 `otherthing.tld`를 추가하면,
`route1`이 더 오래되었기 때문에 `route2`에서 라우트가 하이재킹된다.

이러한 상황을 방지하기 위해 다음 조치를 취해야 한다.

* 게이트웨이에서 관리자는 호스트네임이 특정 네임스페이스 또는 네임스페이스 집합에
명확하게 위임되도록 해야 한다(SHOULD).

{{< tabs name="hostname-config" >}}
{{< tab name="올바른 구성" >}}
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: gateway
spec:
  gatewayClassName: some-class
  listeners:
  - hostname: "something.tld"
    name: listener1
    port: 80
    protocol: HTTP
    allowedRoutes:
      namespaces:
        from: Selector
        selector:
          matchLabels:
            kubernetes.io/metadata.name: ns1
  - hostname: "otherthing.tld"
    name: listener2
    port: 80
    protocol: HTTP
    allowedRoutes:
      namespaces:
        from: Selector
        selector:
          matchLabels:
            kubernetes.io/metadata.name: ns2
```
{{< /tab >}}
{{< tab name="안전하지 않은 구성" >}}
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: gateway
spec:
  gatewayClassName: some-class
  listeners:
  - name: listener1
    port: 80
    protocol: HTTP
    allowedRoutes:
      namespaces:
        from: All
```
{{< /tab >}}
{{< /tabs >}}

#### 64개 이상의 리스너

게이트웨이 리소스에는 리스너 항목이 64개로 제한되어 있다. 64개 이상의 리스너가
필요한 경우, 사용자가 각 네임스페이스의 라우트에서 직접 호스트네임을 설정하도록
허용하되, `ValidatingAdmissionPolicy`와 같은 메커니즘을 활용하여 어떤 네임스페이스가
어떤 호스트네임을 사용할 수 있는지 제한하는 것을 고려해야 한다.

(아직 실험적인) `ListenerSet`을 사용하기로 선택한 경우에도, `ListenerSet`이 요청할 수
있는 호스트네임을 제한하는 유사한 메커니즘을 고려해야 한다.

#### ValidatingAdmissionPolicy 예시

[ValidatingAdmissionPolicy](https://kubernetes.io/docs/reference/access-authn-authz/validating-admission-policy/)를
사용하여 어떤 네임스페이스가 어떤 도메인을 사용할 수 있는지 제한하는 규칙을
추가할 수 있다.

{{< warning >}}
여기에 표시된 검증 정책은 **예시**이며 클러스터 관리자는 자신의 환경에 맞게
조정해야 한다! 적절한 조정 없이 이 예시를 복사/붙여넣기하지 말자.
{{< /warning >}}

여기에 예시로 제시된 정책은 다음을 수행한다.

* 네임스페이스에 존재하는 "domains" `annotation`의 쉼표로 구분된 값에서 허용된 도메인을 읽는다.
* `.spec.hostnames` 내의 모든 호스트네임이 이 어노테이션에 포함되어 있는지 검증한다.
* 항목 중 하나라도 승인되지 않은 경우, 정책이 해당 요청의 승인을 거부한다.

```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicy
metadata:
  name: httproute-hostname-policy
spec:
  failurePolicy: Fail
  matchConstraints:
    resourceRules:
    - apiGroups: ["gateway.networking.k8s.io"]
      apiVersions: ["v1", "v1beta1"]
      operations: ["CREATE", "UPDATE"]
      resources: ["httproutes"]
  variables:
  - name: allowed_hostnames_str
    expression: |
      has(namespaceObject.metadata.annotations) &&
      has(namespaceObject.metadata.annotations.domains) ?
      namespaceObject.metadata.annotations['domains'] : ''
  - name: allowed_hostnames_list
    expression: |
      variables.allowed_hostnames_str.split(',').
      map(h, h.trim()).filter(h, size(h) > 0)
  validations:
  - expression: |
      !has(object.spec.hostnames) ||
      size(object.spec.hostnames) == 0 ||
      object.spec.hostnames.all(hostname, hostname in variables.allowed_hostnames_list)
    message: "HTTPRoute validation failed. It contains unauthorized hostnames"
---
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicyBinding
metadata:
  name: httproute-hostname-binding
spec:
  policyName: httproute-hostname-policy
  validationActions: ["Deny"]
  matchResources:
    namespaceSelector: {}
```

정책이 생성되면, 클러스터 관리자는 다음과 같은 명령어로 도메인 사용을
명시적으로 허용해야 한다.
`kubectl annotate ns default domains=www.dom1.tld,www.dom2.tld`

추가적으로, DNS 레코드 생성을 제공하는 환경을 다룰 때는 관리자가 위와 동일한
제약 조건을 기반으로 DNS 생성을 제한하는 것에 대해 인지하고 있어야 한다.

### 네임스페이스 간 참조 제한

리소스 소유자는 [ReferenceGrant]({{< ref "/reference/api-types/referencegrant" >}})의 사용에
대해 인지하고 있어야 한다.

ReferenceGrant는 리소스 소유자가 자신의 리소스를 다른 네임스페이스의
게이트웨이 API 리소스에서 사용할 수 있게 해준다. 이를 수행할 수 있는
위치를 제한하는 것이 유익할 수 있다.

`ValidatingAdmissionPolicy`를 사용하여 어떤 종류의 `resource`와 어떤
`namespace`가 `ReferenceGrant`를 생성할 수 있는지 제한할 수 있다.

아래는 `referencegrants=allow` 레이블이 있는 네임스페이스에서만
`ReferenceGrant`의 사용을 제한하고, `HTTPRoute` 종류의 객체만 이름이 반드시
지정되어야 하는 `Service` 종류의 객체를 참조할 수 있도록 허용하는 **예시**이다.

```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicy
metadata:
  name: reference-grant-limit
spec:
  failurePolicy: Fail
  matchConstraints:
    resourceRules:
    - apiGroups: ["gateway.networking.k8s.io"]
      apiVersions: ["v1beta1"]
      operations: ["CREATE", "UPDATE"]
      resources: ["referencegrants"]
  variables:
  - name: allowed_grant_ns
    expression: |
      has(namespaceObject.metadata.labels) &&
      has(namespaceObject.metadata.labels.referencegrants) &&
      namespaceObject.metadata.labels['referencegrants'] == 'allow'
  - name: allowed_from_kind
    expression: |
      object.spec.from.all(f, f.kind=='HTTPRoute')
  - name: allowed_to_kind
    expression: |
      object.spec.to.all(t, t.kind == 'Service' && has(t.name) && t.name != '')
  validations:
  - expression: |
      variables.allowed_grant_ns && variables.allowed_from_kind && variables.allowed_to_kind
    message: "ReferenceGrant must be explicitly allowed on the namespace, from an HTTPRoute to a named service"
---
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicyBinding
metadata:
  name: reference-grant-limit
spec:
  policyName: reference-grant-limit
  validationActions: ["Deny"]
  matchResources:
    namespaceSelector: {}
```

ReferenceGrant 소유자는 부여되는 참조 권한이 최소화되도록 해야 한다.

* `to` 대상을 가능한 모든 방법(`group`, `kind`, **그리고** `name`)으로 지정하자.
* 특히 `name`은 선택 사항이더라도 *매우* 타당한 이유 없이 지정하지 않은 채로 두지 말자. 이름을 지정하지 않으면 포괄적인 권한을 부여하는 것이 된다.

### 역할 및 RoleBinding의 적절한 정의

새로운 게이트웨이의 생성은 권한이 부여된 작업으로 간주되어야 한다.
무분별한 게이트웨이 생성은 비용 증가, 인프라 변경(로드 밸런서 및 DNS 레코드 생성 등)을
초래할 수 있으므로, 관리자는 이를 인지하고 적절한 사용자 권한을 반영하는
[Role](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#role-and-clusterrole)과
[RoleBinding](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#rolebinding-and-clusterrolebinding)을
생성해야 한다.

또한 일반 사용자가 게이트웨이 API 상태를 수정할 수 없도록 엄격한 권한을 부여하는 것이
강력히 권장된다.

### GatewayClass의 사용과 제한

클러스터에는 서로 다른 용도의 여러 게이트웨이 클래스가 있을 수 있다. 예를 들어,
하나의 게이트웨이 클래스는 연결된 게이트웨이가 내부 로드 밸런서만 사용하도록
강제할 수 있다.

클러스터 관리자는 이러한 요구 사항을 인지하고, 승인되지 않은 사용자가
게이트웨이를 게이트웨이 클래스에 부적절하게 연결하는 것을 제한하는 검증 정책을
정의해야 한다.

[ValidatingAdmissionPolicy](https://kubernetes.io/docs/reference/access-authn-authz/validating-admission-policy/)를
사용하여 어떤 네임스페이스가 `GatewayClass`를 사용할 수 있는지 제한할 수 있다.
