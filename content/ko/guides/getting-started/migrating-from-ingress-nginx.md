---
title: "Ingress-NGINX 사용자를 위한 가이드"
linkTitle: "Migrating from Ingress-NGINX"
weight: 4
description: "Resources and guidance for migrating from Ingress-NGINX to Gateway API"
---

Ingress-NGINX 사용자라면 적절한 곳에 왔다. 이 페이지는 Gateway API로의 마이그레이션을 고려하거나 진행 중인 사용자를 위한 리소스 허브이다. 시간이 지남에 따라 이 페이지를 더욱 포괄적으로 구성해 나갈 계획이다. Ingress에서의 마이그레이션에 대한 일반적인 개요를 찾고 있다면, [Ingress에서 마이그레이션 가이드]({{< ref "/guides/getting-started/migrating-from-ingress" >}})를 참고하라.

마이그레이션은 복잡할 수 있으며, 목표는 원활한 전환에 필요한 정보와 도구를 제공하는 것이다.

## 자주 묻는 질문

### Gateway API는 Ingress와 어떻게 다른가?

Gateway API는 Ingress보다 역할 지향적이고 더 표현력이 풍부한 API를 제공한다. Ingress가 로드 밸런서와 라우팅 규칙의 개념을 단일 리소스로 결합하는 반면, Gateway API는 이를 분리한다:

* **Gateway:** 트래픽이 클러스터에 진입하는 위치와 방법을 정의하며, 클러스터 운영자의 업무이다.
* **HTTPRoute:** 트래픽이 서비스로 라우팅되는 방법을 정의하며, 애플리케이션 개발자의 업무이다.

이러한 분리는 더 안전한 멀티 테넌트 인프라를 가능하게 한다. 자세한 내용은 [Ingress 마이그레이션 가이드]({{< ref "/guides/getting-started/migrating-from-ingress" >}})를 확인하라.

### Ingress-NGINX 기능을 Gateway API에 어떻게 매핑하는가?

Ingress-NGINX의 많은 어노테이션 기반 기능은 Gateway API에서 대응하는 필드를 가지고 있다. 예를 들어, 트래픽 분할, 헤더 조작, TLS 설정은 모두 API에 기본적으로 포함되어 있다. 자세한 매핑은 [Gateway API HTTPRoute 문서]({{< ref "/reference/api-types/httproute" >}})와 선택한 [구현체]({{< ref "/overview/implementations" >}})의 문서를 참고하라.

### Ingress-NGINX를 제거하지 않고 Gateway API를 시도할 수 있는가?

그렇다. 그렇게 하는 것을 강력히 권장한다. 기존 Ingress-NGINX 컨트롤러와 함께 Gateway API 컨트롤러를 실행할 수 있다. 각각 다른 외부 IP 주소를 받게 되므로, 프로덕션 트래픽에 영향을 주지 않고 새로운 설정을 독립적으로 테스트하고 검증할 수 있다.

## 마이그레이션 리소스

성공적인 마이그레이션에는 신중한 계획이 필요하다. 다음 리소스가 도움이 될 것이다.

### ingress2gateway

복잡한 Ingress 규칙과 어노테이션을 수동으로 변환하는 것은 오류가 발생하기 쉽다. **[ingress2gateway](https://github.com/kubernetes-sigs/ingress2gateway)** 도구는 이 과정을 자동화하도록 설계되었다. 기존 Ingress 리소스를 읽고 해당하는 Gateway 및 HTTPRoute 리소스로 변환한다.

이 도구는 활발히 개발 중이며, 가장 널리 사용되는 Ingress-NGINX 어노테이션을 지원하기 위한 작업이 진행 중이다. 마이그레이션의 시작점으로 이 도구를 사용하는 것을 강력히 권장한다.

### 구현체 선택하기

첫 번째 단계는 요구 사항에 맞는 구현체를 선택하는 것이다. 고려해야 할 주요 요소는 다음과 같다:

* **적합성(Conformance):** [적합성 보고서]({{< ref "/overview/implementations" >}})를 확인하여 구현체가 필요한 Gateway API 기능을 지원하는지 확인하라.
* **기반 기술:** Envoy, NGINX 등의 프록시에 대한 팀의 친숙도가 선택에 영향을 줄 수 있다.
* **통합:** 클라우드 제공자나 CNI가 이미 통합된 Gateway API 구현체를 제공할 수 있다.

## 진행 중인 작업

Gateway API로의 마이그레이션을 더욱 원활하게 만들기 위해 많은 작업이 진행 중이다. ingress2gateway의 지속적인 개선과 v1.0 릴리스를 향한 작업 외에도, 다음 Gateway API 릴리스(현재 2월 목표)를 계획하고 있다. 해당 릴리스에서 다음 기능들을 GA로 졸업시키고자 한다:

* TLSRoute
* ListenerSet
* HTTPRoute CORS 필터

다른 기능이 필요하다면 알려 달라.

## 도움 안내

Gateway API 커뮤니티는 마이그레이션 경험을 최대한 원활하게 만들기 위해 노력하고 있다. `ingress2gateway` 도구는 이의 핵심 부분이며, 어노테이션 지원을 개선하기 위해 적극적으로 작업하고 있다.

질문이 있거나, 문제가 발생하거나, 기능이 부족한 경우 연락 바란다:

* **이슈 제출**: [Gateway API 저장소](https://github.com/kubernetes-sigs/gateway-api/issues)에 이슈를 제출한다.
* **커뮤니티 미팅 참여**: 사용 사례를 논의하기 위해 커뮤니티 미팅에 참여한다.
* **피드백 제공**: `ingress2gateway` 도구에 대한 피드백은 해당 [저장소](https://github.com/kubernetes-sigs/ingress2gateway/issues)에 이슈를 열어 제출한다.

여러분의 피드백은 모든 사용자를 위한 API와 마이그레이션 도구를 개선하는 데 매우 소중하다.
