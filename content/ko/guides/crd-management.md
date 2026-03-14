---
title: "CRD 관리"
weight: 4
description: "Managing Gateway API Custom Resource Definitions"
---

Gateway API는 CRD로 구축되어 있다. 이는 여러 가지 중요한 이점을 제공하는데, 특히 Gateway API의 각 릴리스가 최근 5개의 쿠버네티스 마이너 버전을 지원한다는 점이 주목할 만하다. 이는 이 API의 최신 버전을 사용하기 위해 쿠버네티스 클러스터를 업그레이드할 필요가 없을 가능성이 높다는 것을 의미한다.

그러나 이러한 추가적인 유연성은 혼란의 여지를 남기기도 한다. 이 가이드는 Gateway API CRD 관리와 관련된 가장 일반적인 질문에 답하는 것을 목표로 한다.

## 누가 CRD를 관리해야 하는가?

궁극적으로 CRD는 높은 권한을 가진 클러스터 범위의 리소스이다. 이는 클러스터 관리자 또는 클러스터 제공자가 클러스터 내 CRD 관리를 담당해야 한다는 것을 의미한다.

실질적으로 다음 중 어느 접근 방식이든 합리적이다:

* 클러스터 관리자가 CRD를 설치한다
* 클러스터 프로비저닝 도구 또는 제공자가 CRD를 설치하고 관리한다

일부 구현체는 설치를 단순화하기 위해 CRD를 번들로 제공하고자 할 수 있다. 이는 다음을 절대 하지 않는 한 허용된다:

1. 인식되지 않거나 더 새로운 버전의 Gateway API CRD를 덮어쓰는 것.
1. 다른 릴리스 채널의 Gateway API CRD를 덮어쓰는 것.
1. Gateway API CRD를 제거하는 것.

[Issue #2678](https://github.com/kubernetes-sigs/gateway-api/issues/2678)에서 구현체가 이를 달성하기 위해 사용할 수 있는 가능한 접근 방식 중 하나를 탐구하고 있다.

## 새 버전으로 업그레이드

Gateway API는 두 가지 [릴리스 채널]({{< ref "/overview/concepts/versioning" >}})로 CRD를 릴리스한다.
Standard 채널 CRD를 유지하면 CRD 업그레이드가 더 간단하고 안전해진다.

### 전반적인 가이드라인

1. 되돌리는 것을 피한다. 새로운 버전의 CRD는 새로운 필드와 기능을 추가할 수 있다.
   이전 버전의 CRD로 롤백하면 해당 구성이 손실될 수 있다.
1. 업그레이드 전에 릴리스 노트를 읽는다. 일부 경우, 업그레이드 전에 따라야 할
   가이드라인이 포함되어 있을 수 있다.
1. [Gateway API 버전 관리 정책]({{< ref "/overview/concepts/versioning" >}})을 이해하여 무엇이 변경될 수 있는지 파악한다.
1. 여러 Gateway API 마이너 버전을 한 번에 업그레이드하는 것이 보통 안전하지만,
   가장 안전하고 널리 테스트된 경로는 한 번에 하나의 마이너 버전씩 업그레이드하는 것이다.

### 검증 웹훅

검증 웹훅은 이전 버전의 Gateway API에 포함되어 있었다. v1.0부터 해당 웹훅은 CRD 내에 직접 포함된 CEL 검증을 위해 공식적으로 폐기되었다. Gateway API v1.1에서는 웹훅이 완전히 제거된다. 이는 검증 웹훅이 더 이상 새로운 Gateway API 버전으로 업그레이드할 때 고려 사항이 아니라는 것을 의미한다.

### API 버전 제거

{{< note >}}
This is an advanced use case that is currently only applicable to users that
have been using Gateway API since v0.5.0 within the same cluster.
{{< /note >}}

Gateway API 릴리스에서 더 새롭거나 더 안정적인 API 버전이 있는 CRD의 v1alpha2와 같은 알파 API 버전을 제거할 수 있다. Standard 채널 내에서 API 버전의 제거는 최소 4개의 마이너 릴리스에 걸쳐 진행된다:

1. 더 새로운 API 버전이 스토리지 버전으로 구성된다.
1. 버전이 폐기된다(릴리스 노트에 기록되고 폐기된 API 버전 사용 시 폐기 경고를 통해 알린다).
1. 버전이 더 이상 서빙되지 않지만, API 버전 간 자동 변환을 위해 CRD에 여전히 포함된다.
1. 버전이 더 이상 CRD에 포함되지 않는다.

이 과정을 거친 CRD(스토리지 버전 마이그레이션 포함)를 사용하고 있었다면, 일부 리소스가 이전(폐기된) 스토리지 버전에 고정되어 있을 수 있다. CRD 스토리지 버전이 업데이트되면, 해당 CRD를 사용하는 개별 리소스가 다시 저장될 때에만 적용된다.

예를 들어, Gateway API v0.5.0 CRD를 사용하여 "foo" GatewayClass를 생성했다면, 해당 GatewayClass의 스토리지 버전은 v1alpha2이다. 해당 "foo" GatewayClass가 수정되거나 업데이트되지 않은 채로 남아 있었다면, Gateway API v1.0.0 CRD로 업그레이드할 수 없다. 리소스 중 하나가 여전히 v1alpha2를 스토리지 버전으로 사용하고 있었고, 이는 더 이상 CRD에 포함되지 않기 때문이다(위의 4단계).

업그레이드하려면 이전 스토리지 버전을 사용하는 모든 GatewayClass를 업데이트하는 작업을 수행해야 한다. 예를 들어, 각 GatewayClass에 빈 kubectl patch를 보내면 이 효과를 얻을 수 있다. 다행히도 이를 자동화할 수 있는 도구가 있다 -
[kube-storage-version-migrator](https://github.com/kubernetes-sigs/kube-storage-version-migrator)는 리소스가 최신 스토리지 버전을 사용하도록 자동으로 업데이트한다.

### Experimental 채널

이름에서 알 수 있듯이, Experimental 채널은 Standard 채널과 동일한 안정성 보장을 제공하지 않는다. 마이너 릴리스에서 Experimental 채널 CRD에 대해 다음이 가능하다:

* 기존 API 필드 또는 리소스에 대한 호환성을 깨는 변경
* 사전 폐기 없이 API 필드 또는 리소스 제거

실제로 이는 새로운 Experimental 버전으로의 일부 업그레이드 시 Experimental CRD를 제거하고 다시 설치해야 할 수 있다는 것을 의미한다. 그런 경우가 발생하면 릴리스 노트에서 명확하게 안내된다.
