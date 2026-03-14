---
title: "Gateway API 시작하기"
weight: 1
description: "Getting started with Gateway API installation and basic usage"
---

이 페이지는 새로운 사용자이든 다른 인그레스 솔루션에서 마이그레이션하든 시작하기에 가장 좋은 곳이다.

## 새로운 사용자를 위한 안내

Gateway API를 처음 사용하는 경우 다음 경로를 권장한다:

**1. Gateway API CRD 및 컨트롤러 설치**

[게이트웨이 컨트롤러를 설치](#install-gateway-controller)하면 CRD가 자동으로 설치되는 경우가 많다. 또는 [Gateway API CRD를 수동으로 설치](#install-gateway-api)할 수도 있다.

**2. 가이드 따라하기**

컨트롤러가 실행되면 간단한 예제로 시작한 다음 더 고급 기능을 탐색할 수 있다.

- [간단한 게이트웨이 배포하기]({{< ref "/guides/getting-started/simple-gateway" >}}) (시작하기에 좋은 예제)
- [HTTP 라우팅]({{< ref "/guides/user-guides/http-routing" >}})
- [HTTP 리디렉션 및 재작성]({{< ref "/guides/user-guides/http-redirect-rewrite" >}})
- [HTTP 트래픽 분할]({{< ref "/guides/user-guides/traffic-splitting" >}})
- [네임스페이스 간 라우팅]({{< ref "/guides/user-guides/multiple-ns" >}})
- [TLS 구성]({{< ref "/guides/user-guides/tls" >}})
- [TCP 라우팅]({{< ref "/guides/user-guides/tcp" >}})
- [gRPC 라우팅]({{< ref "/guides/user-guides/grpc-routing" >}})
- [ListenerSet]({{< ref "/guides/user-guides/listener-set" >}})

## Ingress에서 마이그레이션

쿠버네티스 Ingress에서 마이그레이션하려는 경우, 시작하는 데 도움이 되는 가이드가 있다.

- **[Ingress에서 마이그레이션]({{< ref "/guides/getting-started/migrating-from-ingress" >}})**: 모든 Ingress 구현에서 Gateway API로 마이그레이션하기 위한 일반 가이드이다. 주요 차이점을 다루고 수동 마이그레이션 예제를 제공한다.
- **[Ingress-NGINX 사용자를 위한 가이드]({{< ref "/guides/getting-started/migrating-from-ingress-nginx" >}})**: 가장 많이 사용되는 Ingress 구현인 Ingress-NGINX 사용자를 위한 가이드이다. 자주 묻는 질문에 대한 답변과 리소스를 제공한다.

---

## 게이트웨이 컨트롤러 설치하기 {#install-gateway-controller}

Gateway API를 지원하는 [여러 프로젝트]({{< ref "/overview/implementations" >}})가 있다.
쿠버네티스 클러스터에 게이트웨이 컨트롤러를 설치하면 위의 가이드를
따라해 볼 수 있다. 이를 통해 원하는 라우팅 구성이 실제로 Gateway 리소스
(그리고 Gateway 리소스가 나타내는 네트워크 인프라)에 의해 구현되고 있음을
확인할 수 있다. 많은 게이트웨이 컨트롤러 설정은 Gateway API 번들을
자동으로 설치하고 제거한다.

## Gateway API 설치하기 {#install-gateway-api}

{{< danger >}}
**이전 실험적 채널 릴리스에서 업그레이드**


이전 버전의 실험적 채널을 설치한 적이 있다면,
[v1.1 업그레이드 참고 사항](#v11-upgrade-notes)을 참조하라.
{{< /danger >}}

Gateway API 번들은 Gateway API 버전과 관련된 CRD 집합을 나타낸다. 각 릴리스에는
서로 다른 안정성 수준을 가진 두 개의 채널이 포함된다:

### 표준 채널 설치

표준 릴리스 채널에는 GA(정식 출시) 또는 베타로 승격된 모든 리소스가 포함되며,
GatewayClass, Gateway, HTTPRoute, ReferenceGrant가 포함됩니다.
이 채널을 설치하려면 다음 kubectl 명령을 실행한다:

```bash
kubectl apply --server-side -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.5.0/standard-install.yaml
```

이 kubectl 명령 옵션에 대해 더 알아보려면 [서버 사이드 적용 문서](https://kubernetes.io/docs/reference/using-api/server-side-apply/)를 참조하라.

### 실험적 채널 설치

실험적 릴리스 채널에는 표준 릴리스 채널의 모든 것과 일부 실험적
리소스 및 필드가 포함된다. 여기에는 TCPRoute, TLSRoute, UDPRoute가 포함된다.

향후 API 릴리스에는 실험적 리소스 및 필드에 대한 호환성을 깨는 변경이
포함될 수 있다. 예를 들어, 실험적 리소스나 필드가 향후 릴리스에서
제거될 수 있다. 실험적 채널에 대한 자세한 내용은
[버전 관리 문서]({{< ref "/overview/concepts/versioning" >}})를 참조하라.

실험적 채널을 설치하려면 다음 kubectl 명령을 실행한다:

```bash
kubectl apply --server-side -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.5.0/experimental-install.yaml
```

이 kubectl 명령 옵션에 대해 더 알아보려면 [서버 사이드 적용 문서](https://kubernetes.io/docs/reference/using-api/server-side-apply/)를 참조하라.

### v1.2 업그레이드 참고 사항
Gateway API v1.2로 업그레이드하기 전에, Gateway API 구현이 이러한 리소스의
`v1alpha2` API 버전 대신 `v1` API 버전을 지원하도록 업그레이드되었는지
확인해야 한다. YAML 매니페스트에서 `v1`을 사용하고 있더라도, 컨트롤러가
여전히 `v1alpha2`를 사용하고 있을 수 있으며 이는 업그레이드 중 실패를
야기할 수 있다.

의존하고 있는 구현이 v1으로 업그레이드되었음을 확인한 후, v1.2 CRD를
설치할 차례이다. 대부분의 경우 추가 작업 없이 작동한다.

이러한 CRD를 설치하는 데 문제가 발생한 경우, 이러한 CRD 중 하나 또는 둘 다의
`storedVersions`에 `v1alpha2`가 있을 가능성이 높다. 이 필드는
이러한 리소스 중 하나를 저장하는 데 사용된 적이 있는 API 버전을 나타내는 데
사용된다. 안타깝게도 이 필드는 자동으로 정리되지 않는다. 이 값을
확인하려면 다음 명령을 실행한다:

```
kubectl get crd grpcroutes.gateway.networking.k8s.io -ojsonpath="{.status.storedVersions}"
kubectl get crd referencegrants.gateway.networking.k8s.io -ojsonpath="{.status.storedVersions}"
```

이 중 하나라도 "v1alpha2"를 포함하는 목록을 반환하면, `storedVersions`에서
해당 버전을 수동으로 제거해야 한다.

그 전에, 모든 ReferenceGrant 및 GRPCRoute가 최신 스토리지 버전으로
업데이트되었는지 확인하는 것이 좋다:

```
crds=("GRPCRoutes" "ReferenceGrants")

for crd in "${crds[@]}"; do
  output=$(kubectl get "${crd}" -A -o json)

  echo "$output" | jq -c '.items[]' | while IFS= read -r resource; do
    namespace=$(echo "$resource" | jq -r '.metadata.namespace')
    name=$(echo "$resource" | jq -r '.metadata.name')
    kubectl patch "${crd}" "${name}" -n "${namespace}" --type='json' -p='[{"op": "replace", "path": "/metadata/annotations/migration-time", "value": "'"$(date +%Y-%m-%dT%H:%M:%S)"'" }]'
  done
done
```

이제 모든 ReferenceGrant 및 GRPCRoute 리소스가 최신 스토리지 버전을
사용하도록 업데이트되었으므로, ReferenceGrant 및 GRPCRoute CRD를
패치할 수 있다:

```
kubectl patch customresourcedefinitions referencegrants.gateway.networking.k8s.io --subresource='status' --type='merge' -p '{"status":{"storedVersions":["v1beta1"]}}'
kubectl patch customresourcedefinitions grpcroutes.gateway.networking.k8s.io --subresource='status' --type='merge' -p '{"status":{"storedVersions":["v1"]}}'
```

이 단계가 완료되면, 최신 GRPCRoute 및 ReferenceGrant로의 업그레이드가
잘 작동할 것이다.

### v1.1 업그레이드 참고 사항 {#v11-upgrade-notes}
이전 Gateway API 릴리스에서 GRPCRoute 또는 BackendTLSPolicy 실험적 채널
CRD의 이전 버전을 이미 사용하고 있다면, 이 업그레이드에 주의해야 한다.
Gateway API를 처음 설치하거나 API의 표준 채널만 사용한 경우,
이 섹션의 나머지 부분을 건너뛸 수 있다.

#### GRPCRoute
**요약:** v1alpha2 GRPCRoute를 이미 사용하고 있다면, 의존하는 구현이
GRPCRoute v1을 지원하도록 업데이트될 때까지 v1.1에서 GRPCRoute의
실험적 채널을 유지한다.

**설명:** GRPCRoute가 GA로 승격됨에 따라, 이제 표준 채널에 포함된다.
안타깝게도, 이미 GRPCRoute의 실험적 채널 버전을 사용하고 있는 사람에게는
문제가 될 수 있다. 규칙에 따라, 표준 채널의 CRD는 해당 채널에서의
버전 폐기를 피하기 위해 알파 API 버전을 노출하지 않는다. 이는
GRPCRoute의 표준 채널 버전이 v1alpha2를 제외한다는 것을 의미한다.
Gateway API의 v1.1 릴리스 이전에 빌드된 모든 GRPCRoute 구현은
v1alpha2에만 의존했으며 GRPCRoute v1을 지원하도록 업데이트해야 한다.
구현이 v1을 지원하도록 업데이트될 때까지, v1과 v1alpha2를 모두 노출하는
v1.1에 포함된 GRPCRoute의 실험적 채널 버전으로 안전하게 업그레이드할 수 있다.

**업그레이드 순서:** v1alpha2 GRPCRoute를 이미 사용하고 있다면, 다음
업그레이드 순서를 권장한다:

1. *실험적* v1.1 GRPCRoute CRD 설치
2. 모든 매니페스트를 `v1alpha2` 대신 `v1`을 사용하도록 업데이트
3. GRPCRoute `v1` API 버전을 지원하는 구현으로 업그레이드
4. *표준* 채널 v1.1 GRPCRoute CRD 설치

#### BackendTLSPolicy
**요약:** 이전에 BackendTLSPolicy를 설치한 경우, 의존하는 구현이
API의 `v1alpha3`를 지원하도록 업데이트될 때까지 기다린다. `v1alpha3`를
지원하는 구현으로 업그레이드할 때, 새 CRD를 설치하기 전에 이전
BackendTLSPolicy CRD를 제거해야 한다.

**설명:** BackendTLSPolicy는 v1.1에서 여러 중요한 필드 이름이
변경되어 v1alpha3로 버전이 올라갔다. 실험적 채널이므로,
이 변경에 대한 인플레이스 업그레이드 경로를 제공하지 않으며,
대신 의존하는 BackendTLSPolicy 구현과 CRD 업그레이드를 조율해야 한다.

**업그레이드 순서:** v1alpha2 BackendTLSPolicy를 이미 사용하고 있다면, 다음
업그레이드 순서를 권장한다:

1. 선택한 구현이 v1alpha3 지원을 릴리스할 때까지 대기
2. 이전 v1.1 이전 BackendTLSPolicy CRD 삭제 (이렇게 하면 클러스터의 모든
   BackendTLSPolicy 인스턴스도 삭제된다)
3. 새 v1.1 BackendTLSPolicy CRD 설치
4. BackendTLSPolicy v1alpha3를 지원하는 구현 버전 배포

일부 구현에서는 3단계와 4단계의 순서를 바꾸는 것을 선호할 수 있으므로,
선택한 구현의 관련 릴리스 노트를 확인하는 것이 좋다.


### 정리

작업이 끝나면, 위의 명령에서 "apply"를 "delete"로 바꾸어 Gateway API CRD를
제거하여 정리할 수 있다. 이러한 리소스가 사용 중이거나
게이트웨이 컨트롤러에 의해 설치된 경우에는 제거하지 않는다. 이렇게 하면
전체 클러스터의 Gateway API 리소스가 제거된다. 다른 사람이 사용 중일 수
있다면 이 작업을 수행하지 않는다. 이러한 리소스를 사용하는 모든 것이
중단된다.

### CRD 관리에 대한 자세한 내용
이 가이드는 Gateway API를 시작하는 방법에 대한 개략적인 개요만 제공한다.
Gateway API CRD 관리에 대한 자세한 내용은
[CRD 관리 가이드]({{< ref "/guides/crd-management" >}})를 참조하라.
