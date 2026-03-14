---
title: "목록"
description: "Gateway API의 다운스트림 구현체 및 통합"
weight: 1
type: docs
---

이 문서는 게이트웨이 API의 다운스트림 구현 및 통합을 추적하고
이들에 대한 상태 및 리소스 참조를 제공한다.

게이트웨이 API의 구현자와 통합자들은 자신들의 구현체에 대한 상태 정보, 지원하는 버전, 그리고
사용자가 시작하는 데 도움이 되는 문서를, 이 문서에 업데이트하는 것이
권장된다. 이 상태 정보는 몇 단락을 넘지 않아야 한다.

## 호환성(Conformance) 수준

게이트웨이 API 호환성에는 세 가지 수준이 있다.

### 호환(Conformant) 구현체

이 구현체들은 다음 항목에 대해 통과한 호환성 보고서를 최소 하나 이상 제출하였다.

  * 하나 이상의 Route 타입과 프로필(Profile) 조합에 대한 모든 코어(Core) 호환성 테스트
  * 주장하는 모든 확장(Extended) 기능

가장 최근 두(2)개의 게이트웨이 API 릴리스 중 하나에 대해 위 조건을 만족해야 한다.

즉, Mesh + HTTPRoute, Gateway + HTTPRoute, Gateway + TLSRoute,
Gateway + Mesh + HTTPRoute 등을 지원하는 것이 호환이며, 구현체가 주장하는 확장 기능도 포함된다.
단, 구현체는 _반드시_ 최소 하나의 프로필과 해당 프로필 내 하나의 Route 타입을 지원해야 하며,
해당 프로필과 Route 타입에 대한 모든 코어 호환성 테스트와 주장하는 모든 확장 기능을 통과해야 한다.

### 부분 호환(Partially Conformant) 구현체

이 구현체들은 완전한 호환을 목표로 하고 있지만 현재 달성하지 못한 상태이다.
가장 최근 세(3)개의 게이트웨이 API 릴리스 중 하나에 대해 호환(위 기준)이 되기 위한
일부 테스트를 통과한 호환성 보고서를 최소 하나 이상 제출하였다.
"부분 호환"으로 간주되기 위한 요구 사항은 향후 게이트웨이 API 릴리스에서 강화될 수 있다.

### 비활성(Stale) 구현체

이 구현체들은 현재 활발하게 개발되고 있지 않을 수 있으며,
다음 페이지 검토 시 다른 카테고리로 이동하는 호환성 보고서를 제출하지 않으면
이 페이지에서 제거될 것이다.

페이지 검토는 모든 게이트웨이 API 릴리스 이후 최소 1개월 후에 수행되며,
첫 번째 검토는 게이트웨이 API v1.3 릴리스 이후 2025년 6월 말에 수행되었다.
게이트웨이 API v1.5 검토 프로세스(2026년 중반 예정) 이후에는
비활성 구현체가 더 이상 목록에 표시되지 않는다.

## 구현 프로필(Implementation profiles)

구현체는 일반적으로 _프로필_ 이라고 하는 두 가지 범주에 해당한다.

* **Gateway** 컨트롤러는 Gateway 리소스를 조정하며, 주로 클러스터 외부에서
내부로 들어오는 북/남(north-south) 트래픽을 처리하기 위한 것이다.
* **Mesh** 컨트롤러는 HTTPRoute가 연결된 Service 리소스를 조정하며,
동일 클러스터 또는 클러스터 집합 내의 동/서(east-west) 트래픽을 처리하기 위한 것이다.

각 프로필에는 구현체가 호환(위 기준)이 되기 위해 기대되는 동작을 정의하는
호환성 테스트 세트가 연결되어 있다.

구현체는 두 프로필 모두에 해당할 수도 있다.

## 통합(Integrations)

이 페이지에는 게이트웨이 API 리소스를 활용하여 다른 기능(DNS 관리 또는 인증서 생성 등)을
수행할 수 있는 다른 소프트웨어 프로젝트인 **통합**(integrations) 도 나열되어 있다.

{{< note >}}
이 페이지에는 게이트웨이 API가 작동하는 데 필요한 기능을 제공하는
서드 파티 프로젝트에 대한 링크가 포함되어 있다. 게이트웨이 API 프로젝트 작성자는
이러한 프로젝트에 대해 책임지지 않으며, 해당 클래스 내에서 알파벳순으로 나열되어 있다.
{{< /note >}}

{{< note >}}
**구현체 간 확장 지원 기능 비교**

[프로젝트의 지원 기능을 빠르게 비교할 수 있는 표 확인](comparisons/v1.4/). 이 표는 코어 호환성 테스트를 통과한 게이트웨이 컨트롤러 구현체를 개요로 제시하며, 구현한 확장 호환성 기능에 초점을 맞춘다. 이 표는 최소 3개의 구현체가 [호환성 보고서](https://github.com/kubernetes-sigs/gateway-api/tree/main/conformance/reports) 하위에 보고서를 업로드한 후 생성되어 사이트에 업로드된다.
{{< /note >}}

## 게이트웨이 컨트롤러 구현 상태 {#gateways}

### 호환(Conformant)
- [Agent Gateway](#agent-gateway-with-kgateway)
- [Airlock Microgateway](#airlock-microgateway)
- [Cilium](#cilium)
- [Envoy Gateway](#envoy-gateway) (GA)
- [Istio](#istio) (GA)
- [kgateway](#kgateway) (GA)
- [NGINX Gateway Fabric](#nginx-gateway-fabric) (GA)
- [Traefik Proxy](#traefik-proxy) (GA)

### 부분 호환(Partially Conformant)

- [AWS Load Balancer Controller](#aws-load-balancer-controller) (GA)
- [Azure Application Gateway for Containers](#azure-application-gateway-for-containers) (GA)
- [Contour](#contour) (GA)
- [Gloo Gateway](#gloo-gateway) (GA)
- [Google Kubernetes Engine](#google-kubernetes-engine) (GA)
- [Gravitee Kubernetes Operator](#gravitee-kubernetes-operator) (GA)
- [Kong Ingress Controller](#kong-kubernetes-ingress-controller) (GA)
- [Kong Gateway Operator](#kong-gateway-operator) (GA)
- [Kubvernor](#kubvernor)(진행 중)

### 비활성(Stale)

- [Acnodal EPIC](#acnodal-epic)
- [Amazon Elastic Kubernetes Service](#amazon-elastic-kubernetes-service) (GA)
- [Apache APISIX](#apisix) (beta)
- [Avi Kubernetes Operator](#avi-kubernetes-operator)
- [Easegress](#easegress) (GA)
- [Emissary-Ingress (Ambassador API Gateway)](#emissary-ingress-ambassador-api-gateway) (alpha)
- [Flomesh Service Mesh](#flomesh-service-mesh-fsm) (beta)
- [HAProxy Ingress](#haproxy-ingress) (alpha)
- [HAProxy Kubernetes Ingress Controller](#haproxy-kubernetes-ingress-controller) (GA)
- [HashiCorp Consul](#hashicorp-consul)
- [Kuma](#kuma) (GA)
- [LiteSpeed Ingress Controller](#litespeed-ingress-controller)
- [LoxiLB](#loxilb) (beta)
- [ngrok](#ngrok-kubernetes-operator) (preview)
- [STUNner](#stunner) (beta)
- [Tyk](#tyk) (진행 중)
- [WSO2 APK](#wso2-apk) (GA)

## 서비스 메시 구현 상태 {#meshes}

### 호환(Conformant)
- [Alibaba Cloud Service Mesh](#alibaba-cloud-service-mesh) (GA)
- [Istio](#istio) (GA)
- [Linkerd](#linkerd) (GA)
- [Cilium](#cilium) (GA)

### 비활성(Stale)
- [Google Cloud Service Mesh](#google-cloud-service-mesh) (GA)
- [Kuma](#kuma) (GA)

## 통합 {#integrations}

- [Flagger](#flagger) (public preview)
- [cert-manager](#cert-manager) (alpha)
- [argo-rollouts](#argo-rollouts) (alpha)
- [Knative](#knative) (alpha)
- [Kuadrant](#kuadrant) (GA)
- [kruise-rollouts](#kruise-rollouts) (alpha)

## 구현체

이 섹션에서는 특정 구현체들에 대한 블로그 게시글, 문서 및 기타 게이트웨이 API 참조에 대한 구체적인 링크를 찾을 수 있다.

### Acnodal EPIC
[EPIC](https://www.epic-gateway.org/)는 쿠버네티스와 함께 설계되고 구축된 오픈 소스 외부 게이트웨이 플랫폼이다. 게이트웨이 클러스터, k8s 게이트웨이 컨트롤러, 독립형 리눅스 게이트웨이 컨트롤러 및 게이트웨이 서비스 매니저로 구성된다. 이들은 함께 클러스터 사용자에게 게이트웨이 서비스를 제공하는 플랫폼을 만든다. 각 게이트웨이는 워크로드 클러스터가 아닌 게이트웨이 클러스터에서 실행되는 여러 Envoy 인스턴스로 구성된다. 게이트웨이 서비스 매니저는 공용 및 사설 클러스터를 위한 Gateway-as-a-Service 인프라를 구현하고 비-k8s 엔드포인트를 통합하는 데 사용할 수 있는 간단한 사용자 관리 및 UI이다.

- [문서](https://www.epic-gateway.org/)
- [소스 저장소](https://github.com/epic-gateway)

### Agentgateway {#agent-gateway-with-kgateway}

[Agentgateway](https://agentgateway.dev/)는 Linux Foundation의 일부로 호스팅되는 오픈 소스 게이트웨이 API 구현체로, LLM 소비, LLM 서빙, 에이전트 간([A2A](https://a2aproject.github.io/A2A/latest/)), 에이전트-도구 간([MCP](https://modelcontextprotocol.io/introduction)) 통신 및 기존 TCP/HTTP 트래픽 서빙을 포함한 AI 사용 사례에 초점을 맞추고 있다.
쿠버네티스 게이트웨이 API를 위해 특별히 설계된 최초이자 유일한 프록시로, 고성능 및 확장 가능한 Rust 데이터플레인 구현으로 구동된다.

### Airlock Microgateway

[Airlock Microgateway](https://www.airlock.com/en/secure-access-hub/components/microgateway)는 쿠버네티스 환경에 최적화되고 RedHat OpenShift 인증을 받은 쿠버네티스 네이티브 WAAP(Web Application and API Protection, 이전 WAF) 솔루션이다.
현대적인 애플리케이션 보안이 개발 워크플로에 내장되어 DevSecOps 패러다임을 따른다.
Airlock Microgateway는 검증된 Airlock 보안 기능으로 애플리케이션과 마이크로서비스를 공격으로부터 보호하며, 높은 확장성도 제공한다.

#### 기능
- 알려진 공격(OWASP Top 10)으로부터 보호하는 거부 규칙, 헤더 필터링, JSON 파싱, OpenAPI 명세 강제 적용, GraphQL 스키마 검증과 같은 보안 기능을 갖춘 포괄적인 WAAP(이전 WAF)
- JWT 인증 또는 OIDC를 사용한 인증 강제를 가능하게 하는 ID 인식 프록시(OAuth 2.0 Token Introspection 및 Token Exchange를 통한 서비스 간 지속적인 검증 및 안전한 위임 포함)
- 요청 라우팅 규칙, TLS 종료 및 원격 IP 추출을 포함한 리버스 프록시 기능
- 허용 및 차단된 트래픽과 기타 메트릭에 대한 유용한 인사이트를 제공하는 사용하기 쉬운 Grafana 대시보드

#### 문서 및 링크
- [제품 문서](https://docs.airlock.com/microgateway/latest)
- [게이트웨이 상세 문서](https://docs.airlock.com/microgateway/latest/?topic=MGW-00000142)
- 도움을 위해 [Airlock 커뮤니티 포럼](https://forum.airlock.com/)과 [지원 프로세스](https://techzone.ergon.ch/support-process)를 확인하자.

### Alibaba Cloud Service Mesh

[Alibaba Cloud Service Mesh (ASM)](https://www.alibabacloud.com/help/en/asm/product-overview/what-is-asm)는 커뮤니티 Istio와 호환되는 완전 관리형 서비스 메시 플랫폼을 제공한다. 서비스 호출 간 트래픽 라우팅 및 분할 관리, 서비스 간 통신의 인증 보안, 메시 관찰 가능성 기능을 포함한 서비스 거버넌스를 단순화하여 개발 및 운영 작업 부담을 크게 줄여준다.

### Amazon Elastic Kubernetes Service

[Amazon Elastic Kubernetes Service (EKS)](https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html)는 자체 쿠버네티스 컨트롤 플레인이나 노드를 설치, 운영, 유지 관리할 필요 없이 AWS에서 쿠버네티스를 실행하는 데 사용할 수 있는 관리형 서비스이다. EKS는 [AWS 게이트웨이 API 컨트롤러](https://github.com/aws/aws-application-networking-k8s)를 통해 게이트웨이 API를 구현하며, 이 컨트롤러는 EKS 클러스터 내 게이트웨이 및 HTTPRoute를 위해 [Amazon VPC Lattice](https://aws.amazon.com/vpc/lattice/) 리소스를 프로비저닝한다.

### APISIX

[Apache APISIX](https://apisix.apache.org/)는 동적이고 실시간이며 고성능인 API 게이트웨이이다. APISIX는 로드 밸런싱, 동적 업스트림, 카나리 릴리스, 서킷 브레이킹, 인증, 관찰 가능성 등과 같은 풍부한 트래픽 관리 기능을 제공한다.

APISIX는 현재 [Apache APISIX 인그레스 컨트롤러](https://github.com/apache/apisix-ingress-controller)에 대해 게이트웨이 API `v1beta1` 버전의 명세를 지원한다.

### AWS Load Balancer Controller

[AWS Load Balancer Controller](https://github.com/kubernetes-sigs/aws-load-balancer-controller)는 쿠버네티스 클러스터를 위한 AWS Elastic Load Balancer를 관리한다. 이 컨트롤러는 쿠버네티스 Ingress를 생성할 때 AWS Application Load Balancer(ALB)를 프로비저닝하고, LoadBalancer 타입의 쿠버네티스 서비스를 생성할 때 AWS Network Load Balancer(NLB)를 프로비저닝한다.

게이트웨이 API 지원은 레이어 4(L4) 및 레이어 7(L7) 라우팅 모두 GA 상태이며, 확장 가능한 게이트웨이 API를 사용하여 쿠버네티스 클러스터에서 직접 AWS NLB 및 ALB를 프로비저닝하고 관리할 수 있다.

게이트웨이 API 구현을 배포하고 사용하는 방법에 대한 정보는 [AWS Load Balancer Controller 문서](https://kubernetes-sigs.github.io/aws-load-balancer-controller/)를 참조하자.

### Avi Kubernetes Operator

[Avi Kubernetes Operator (AKO)](https://techdocs.broadcom.com/us/en/vmware-security-load-balancing/avi-load-balancer/avi-kubernetes-operator/2-1.html)는 VMware AVI Advanced Load Balancer를 사용하여 L4-L7 로드 밸런싱을 제공한다.

AKO 버전 v2.1.1부터 게이트웨이 API 버전 v1.3.0이 지원된다. 게이트웨이 클래스(GatewayClass), 게이트웨이(Gateway) 및 HTTPRoute 객체를 지원하는 게이트웨이 API 명세의 v1 버전을 구현한다.

AKO 게이트웨이 API를 배포하고 사용하는 문서는 [Avi 쿠버네티스 오퍼레이터 게이트웨이 API](https://techdocs.broadcom.com/us/en/vmware-security-load-balancing/avi-load-balancer/avi-kubernetes-operator/2-1/avi-kubernetes-operator-guide-2-1/gateway-api/gateway-api-v1.html)에서 찾을 수 있다.

### Azure Application Gateway for Containers

[Application Gateway for Containers](https://aka.ms/appgwcontainers/docs)는 Azure의 쿠버네티스 클러스터에서 실행되는 워크로드에 대한 동적 트래픽 관리 기능을 제공하는 관리형 애플리케이션(레이어 7) 로드 밸런싱 솔루션이다. ALB 컨트롤러를 배포하고 게이트웨이 API를 시작하려면 [빠른 시작 가이드](https://learn.microsoft.com/azure/application-gateway/for-containers/quickstart-deploy-application-gateway-for-containers-alb-controller)를 따른다.

### Cilium

[Cilium](https://cilium.io)은 쿠버네티스 및 기타 네트워킹 환경을 위한 eBPF 기반 네트워킹, 관찰 가능성 및 보안 솔루션이다.
여기에는 [Cilium Service Mesh](https://docs.cilium.io/en/stable/gettingstarted/#service-mesh)가 포함되어 있으며,
이는 높은 효율을 가진 메시 데이터 플레인으로 [사이드카 없는 모드](https://isovalent.com/blog/post/cilium-service-mesh/)에서 실행될 수 있어 성능을 크게 향상시키고,
사이드카로 인한 운영 복잡성을 피할 수 있다.
Cilium은 또한 사이드카 프록시 모델도 지원하여 사용자에게 선택권을 제공한다.
Cilium은 게이트웨이 API를 지원하며, Cilium 1.19 기준으로 v1.4.0에 대한 호환성을 통과한다.

Cilium은 오픈 소스이며 CNCF 졸업 프로젝트이다.

Cilium 서비스 메시에 대한 질문이 있다면 [Cilium Slack](https://slack.cilium.io)의 #service-mesh 채널에서 시작하는 것이 좋다.
개발 노력에 기여하려면 #development 채널을 확인하거나,
[주간 개발자 회의](https://github.com/cilium/cilium#weekly-developer-meeting)에 참여하자.

### Contour

[Contour](https://projectcontour.io)는 쿠버네티스를 위한 CNCF 오픈 소스로 Envoy 기반 인그레스 컨트롤러이다.

Contour v1.31.0은 게이트웨이 API v1.2.1을 구현한다.
모든 표준 채널 v1 API 그룹 리소스(GatewayClass, Gateway, HTTPRoute, ReferenceGrant)와 대부분의 v1alpha2 API 그룹 리소스(TLSRoute, TCPRoute, GRPCRoute, ReferenceGrant, BackendTLSPolicy)가 지원된다.
Contour의 구현은 v1.2.1 릴리스에 포함된 대부분의 코어 확장 게이트웨이 API 호환성 테스트를 통과한다.

Contour의 게이트웨이 API 구현을 배포하고 사용하는 방법에 대한 정보는 [Contour 게이트웨이 API 가이드](https://projectcontour.io/docs/1.30/guides/gateway-api/)를 확인하자.

Contour의 구현에 대한 도움과 지원을 받으려면, [이슈를 생성](https://github.com/projectcontour/contour/issues/new/choose)하거나 [쿠버네티스 slack의 #contour 채널](https://kubernetes.slack.com/archives/C8XRH2R4J)에서 도움을 요청하자.

### Easegress

[Easegress](https://megaease.com/easegress/)는 클라우드 네이티브 트래픽 오케스트레이션 시스템이다.

이 시스템은 현대적인 고급 게이트웨이, 견고한 분산 클러스터, 유연한 트래픽 오케스트레이터, 또는 접근 가능한 서비스 메시로 기능할 수 있다.

Easegress는 현재 [게이트웨이 컨트롤러](https://github.com/megaease/easegress/blob/main/docs/04.Cloud-Native/4.2.Gateway-API.md)를 통해 게이트웨이 API `v1beta1` 버전의 명세를 지원한다.

### Emissary-Ingress (Ambassador API Gateway)

[Emissary-Ingress](https://www.getambassador.io/docs/edge-stack) (이전 Ambassador API Gateway)는 [Envoy Proxy](https://envoyproxy.io) 위에 구축된
쿠버네티스용 인그레스 컨트롤러와 API 게이트웨이를 제공하는 오픈 소스 CNCF 프로젝트이다.
Emissary와 함께 게이트웨이 API를 사용하는 자세한 내용은 [여기](https://www.getambassador.io/docs/edge-stack/latest/topics/using/gateway-api/)를 참조하자.

### Envoy Gateway

[Envoy Gateway](https://gateway.envoyproxy.io/)는 Envoy 기반 애플리케이션 게이트웨이를 관리하기 위한 [Envoy](https://github.com/envoyproxy) 하위 프로젝트이다.
지원되는 게이트웨이 API의 API와 필드는 [여기](https://gateway.envoyproxy.io/docs/tasks/quickstart/)에 설명되어 있다.
몇 가지 간단한 단계로 게이트웨이 API와 함께 Envoy Gateway를 실행하려면 [빠른 시작](https://gateway.envoyproxy.io/docs/tasks/quickstart)을
사용하자.

### Flomesh Service Mesh (FSM)

[Flomesh Service Mesh](https://github.com/flomesh-io/fsm)는 쿠버네티스 동/서 및 북/남 트래픽 관리를 위한 커뮤니티 주도의 경량 서비스 메시이다. Flomesh는 레이어4 트래픽 관리를 위해 ebpf를, 레이어7 트래픽 관리에 pipy 프록시를 사용한다. Flomesh는 로드 밸런서, 크로스 클러스터 서비스 등록/발견을 내장으로 제공하며, 멀티 클러스터 네트워킹을 지원한다. `Ingress`("인그레스 컨트롤러"로서)와 게이트웨이 API를 지원한다.

FSM의 게이트웨이 API 지원은 [Flomesh 게이트웨이 API](https://github.com/flomesh-io/fgw) 위에 구축되며 현재 쿠버네티스 게이트웨이 API 버전 v0.7.1을 지원하고 `v0.8.0` 지원이 현재 진행 중이다.

- [FSM 쿠버네티스 게이트웨이 API 호환성 매트릭스](https://github.com/flomesh-io/fsm/blob/main/docs/gateway-api-compatibility.md)
- [FSM에서 게이트웨이 API 지원을 사용하는 방법](https://github.com/flomesh-io/fsm/blob/main/docs/tests/gateway-api/README.md)

### Gloo Gateway

[Solo.io](https://www.solo.io)의 [Gloo 게이트웨이](https://docs.solo.io/gateway/latest/)는 기능이 풍부한 쿠버네티스 네이티브 인그레스 컨트롤러이자 차세대 API 게이트웨이이다.
Gloo 게이트웨이는 기존 컨트롤 플레인 구현에 게이트웨이 API의 완전한 기능과 커뮤니티 지원을 제공한다.

Gloo 게이트웨이 인그레스 컨트롤러는 `HTTPRouteServiceTypes`를 제외하고
v1.1.0 릴리스의 GATEWAY_HTTP 호환성 프로필에 대한 모든 코어 게이트웨이 API 호환성 테스트를 통과한다.

### Google Cloud Service Mesh

[Google Kubernetes Engine (GKE)](https://cloud.google.com/kubernetes-engine)는 구글 클라우드에서 제공하는
관리형 쿠버네티스 플랫폼이다.

GKE의 메시를 위한 게이트웨이 (GAMMA) 구현은 [클라우드 서비스 메시](https://cloud.google.com/products/service-mesh)를 통해 이루어진다.

구글 클라우드 서비스 메시는 [Envoy 기반 사이드카 메시](https://cloud.google.com/service-mesh/docs/gateway/set-up-envoy-mesh)와 [Proxyless-GRPC](https://cloud.google.com/service-mesh/docs/gateway/proxyless-grpc-mesh) (GRPCRoute 사용)를 지원한다.

### Google Kubernetes Engine

[Google 쿠버네티스 엔진 (GKE)](https://cloud.google.com/kubernetes-engine)은 구글 클라우드에서 제공하는
관리형 쿠버네티스 플랫폼이다.
GKE의 게이트웨이 API 구현은 GKE 클러스터의 파드를 위한 구글 클라우드 로드 밸런서를 프로비저닝하는
[GKE 게이트웨이 컨트롤러](https://cloud.google.com/kubernetes-engine/docs/concepts/gateway-api)를 통해 이루어진다.

GKE 게이트웨이 컨트롤러는 가중치 트래픽 분할, 미러링, 고급 라우팅, 멀티 클러스터 로드 밸런싱 등을
지원한다.
[사설 또는 공용 게이트웨이](https://cloud.google.com/kubernetes-engine/docs/how-to/deploying-gateways) 및 [멀티 클러스터 게이트웨이](https://cloud.google.com/kubernetes-engine/docs/how-to/deploying-multi-cluster-gateways)를
배포하는 방법은 문서를 참조한다.

GKE 게이트웨이 컨트롤러는 `HTTPRouteHostnameIntersection`을 제외하고
v1.4.0 릴리스의 GATEWAY_HTTP 호환성 프로필에 대한 모든 코어 게이트웨이 API 호환성 테스트를 통과한다.

### Gravitee Kubernetes Operator

[Gravitee Kubernetes Operator](https://documentation.gravitee.io/gravitee-kubernetes-operator-gko) (GKO)를 사용하면 [Gravitee](https://www.gravitee.io/) API, 애플리케이션 및 기타 자산을 쿠버네티스 네이티브이고 선언적인 방식으로 관리할 수 있다.

Gravitee Kubernetes Operator는 버전 4.10.3에서 Gateway - HTTP 기능에 대한 부분 호환을 제공한다. 라우트 간 매칭 규칙은 지원하지 않는다. 이 기능은 향후 릴리스에서 도입될 예정이다.

지원, 피드백 또는 Gravitee Kubernetes Operator에 대한 토론 참여를 원하면 자유롭게 [이슈](https://github.com/gravitee-io/issues/issues)를 제출하거나 커뮤니티 [포럼](https://community.gravitee.io/c/support/gravitee-kubernetes-operator-gko/26)을 방문하자.

### HAProxy Ingress

[HAProxy 인그레스](https://haproxy-ingress.github.io/)는 HAProxy를 위한 커뮤니티 주도 인그레스 컨트롤러 구현이다.

HAProxy 인그레스 v0.13은 게이트웨이 API의 v1alpha1 명세를 부분적으로 지원한다. 호환성과 로드맵에 대한 정보는 [컨트롤러의 게이트웨이 API 문서](https://haproxy-ingress.github.io/docs/configuration/gateway-api/)를 참조한다.

### HAProxy Kubernetes Ingress Controller

HAProxy 쿠버네티스 인그레스 컨트롤러는 HAProxy Technologies에서 유지 관리하는 오픈 소스 프로젝트로, 쿠버네티스를 위한 빠르고 효율적인 트래픽 관리, 라우팅 및 관찰 가능성을 제공한다. 버전 1.10부터 게이트웨이 API에 대한 내장 지원을 제공한다. 동일한 인그레스 컨트롤러 배포로 인그레스 API와 게이트웨이 API를 모두 사용할 수 있다. 자세한 내용은 [문서](https://www.haproxy.com/documentation/kubernetes-ingress/gateway-api/enable-gateway-api/)를 참조하자. [GitHub 저장소](https://github.com/haproxytech/kubernetes-ingress/blob/master/documentation/gateway-api.md)에서 지원되는 API 리소스에 대한 추가 정보도 찾을 수 있다.

### HashiCorp Consul

[HashiCorp](https://www.hashicorp.com)의 [Consul](https://consul.io)은 멀티 클라우드 네트워킹을 위한 오픈 소스 컨트롤 플레인이다. 단일 Consul 배포로 베어 메탈, VM 및 컨테이너 환경에 걸쳐 확장될 수 있다.

Consul 서비스 메시는 모든 쿠버네티스 배포판에서 작동하고, 다중 클러스터를 연결을 지원하며, Consul CRD는 메시에서 트래픽 패턴과 권한을 관리하는 쿠버네티스 네이티브 워크플로를 제공한다. [Consul API 게이트웨이](https://www.consul.io/docs/api-gateway)는 북/남 트래픽 관리를 위한 게이트웨이 API를 지원한다.

게이트웨이 API의 지원되는 버전과 기능에 대한 최신 정보는 [Consul API 게이트웨이 문서](https://www.consul.io/docs/api-gateway)를 확인하길 바란다.

### Istio

[Istio](https://istio.io)는 오픈 소스 [서비스 메시](https://istio.io/latest/docs/concepts/what-is-istio/#what-is-a-service-mesh) 및 게이트웨이 구현체이다.

Istio의 최소 설치만으로 클러스터 인그레스 트래픽 제어를 위한
쿠버네티스 게이트웨이 API 완전한 적합 구현을 사용할 수 있다.
서비스 메시 사용자를 위해,
Istio는 메시 내에서 GAMMA 이니셔티브의 게이트웨이 API
동/서 트래픽 관리 지원도 완전히 지원한다.

모든 [인그레스 작업](https://istio.io/latest/docs/tasks/traffic-management/ingress/)과 여러 메시 내부 트래픽 관리 작업을 포함한 Istio 문서의 대부분은 이미
게이트웨이 API 또는 Istio 구성 API를 사용하여 트래픽을 구성하는 병렬 지침을 포함한다.
Istio의 게이트웨이 API 구현에 대한 자세한 정보는 [게이트웨이 API task](https://istio.io/latest/docs/tasks/traffic-management/ingress/gateway-api/)를 확인하자.

### kgateway

[kgateway](https://kgateway.dev/docs) 프로젝트는 기능이 풍부한 쿠버네티스 네이티브 인그레스 컨트롤러이자 차세대 API 게이트웨이이다.
우수한 HTTP 경험을 유지하는 데 중점을 두고 있으며, AI 및 MCP 게이트웨이와 같은 시나리오에서 고급 라우팅 기능을 확장하고, Istio와 같은 서비스 메쉬와 엠비언트 모드 및 사이드카 모드에서 상호 운용성을 지원한다.
이러한 초점은 많은 북/남 및 동/서 사용 사례에서 성능 효율적 방식인 합리적으로 분산된 Envoy 인스턴스 세트를 쉽게 구성할 수 있음을 의미한다.

Kgateway는 2.0 릴리스와 함께 일반적으로 사용 가능하다.

### Kong Kubernetes Ingress Controller

[Kong](https://konghq.com)은 하이브리드 및 멀티 클라우드 환경을 위해 구축된 오픈 소스 API 게이트웨이이다.

[Kong 쿠버네티스 인그레스 컨트롤러 (KIC)](https://github.com/kong/kubernetes-ingress-controller)는 비관리형 게이트웨이를 구성하는 데 사용할 수 있다. 사용 정보는 [Gateway API Guide](https://docs.konghq.com/kubernetes-ingress-controller/latest/guides/using-gateway-api/)를 확인하자.

Kong 쿠버네티스 인그레스 컨트롤러에 대한 도움과 지원을 받으려면 [이슈를 생성](https://github.com/Kong/kubernetes-ingress-controller/issues/new)하거나 [토론](https://github.com/Kong/kubernetes-ingress-controller/discussions/new)을 만들자. [쿠버네티스 slack의 #kong 채널](https://kubernetes.slack.com/archives/CDCA87FRD)에서도 도움을 요청할 수 있다.

### Kong Gateway Operator

[Kong](https://konghq.com)은 하이브리드 및 멀티 클라우드 환경을 위해 구축된 오픈 소스 API 게이트웨이이다.

[Kong 게이트웨이 오퍼레이터 (KGO)](https://docs.konghq.com/gateway-operator/latest/)는 관리형 게이트웨이를 구성하고, Kong 쿠버네티스 인그레스 컨트롤러의 인스턴스를 오케스트레이션하는 데 사용할 수 있다.

Kong 게이트웨이 오퍼레이터에 대한 도움과 지원을 받으려면 [이슈를 생성](https://github.com/Kong/gateway-operator/issues/new)하거나 [토론](https://github.com/Kong/gateway-operator/discussions/new)을 만들자. [쿠버네티스 slack의 #kong 채널](https://kubernetes.slack.com/archives/CDCA87FRD)에서도 도움을 요청할 수 있다.

### Kubvernor
[Kubvernor](https://github.com/kubvernor/kubvernor)는 Rust 프로그래밍 언어로 구현된 오픈소스이자 고도로 실험적인 API 컨트롤러이다. 현재 Kubvernor는 Envoy Proxy를 지원하며, 다양한 게이트웨이(Envoy, Nginx, HAProxy 등)를 관리/배포할 수 있도록 가능한 한 일반적인 구조를 목표로 한다.

### Kuma

[Kuma](https://kuma.io)는 오픈 소스 서비스 메시이다.

Kuma는 베타 안정성을 보장하며 Kuma 내장형, Envoy 기반 게이트웨이에 대한 게이트웨이 API 명세를 구현한다. 게이트웨이 API를 사용하여 Kuma 내장 게이트웨이를 설정하는 방법에 대한 정보는 [게이트웨이 API 문서](https://kuma.io/docs/latest/using-mesh/managing-ingress-traffic/gateway-api/)을 확인한다.

Kuma 2.3 이상은 메시 내에서 GAMMA 이니셔티브의
게이트웨이 API 동/서 트래픽 관리 지원을 지원한다.

### Linkerd

[Linkerd](https://linkerd.io/)는 최초의 CNCF 졸업 [서비스 메시](https://buoyant.io/service-mesh-manifesto)이다.
Envoy를 기반으로 하지 않은 유일한 주요 메쉬로,
대신 Rust로 특별히 설계된 마이크로 프록시를 활용해
Kubernetes에 보안, 가시성, 신뢰성을 제공하며 복잡성을 제거한다.

Linkerd 2.14 이상은 메시 내에서 GAMMA 이니셔티브의
게이트웨이 API 동/서 트래픽 관리 지원을 지원한다.

### LiteSpeed Ingress Controller

[LiteSpeed 인그레스 컨트롤러](https://litespeedtech.com/products/litespeed-web-adc/features/litespeed-ingress-controller)는 LiteSpeed WebADC 컨트롤러를 사용하여 인그레스 컨트롤러 및 로드 밸런서로 동작하며, 쿠버네티스 클러스터 내의 트래픽을 관리한다. 이 컨트롤러는 Gateway, GatewayClass, HTTPRoute, ReferenceGrant를 포함한 게이트웨이 API의 코어 기능 전체와 cert-manager의 게이트웨이 기능을 구현하고 있다. 게이트웨이는 LiteSpeed Ingress Controller에 완전히 통합되어 있다.

- [제품 문서](https://docs.litespeedtech.com/cloud/kubernetes/).
- [게이트웨이 상세 문서](https://docs.litespeedtech.com/cloud/kubernetes/gateway).
- 전체 지원은 [LiteSpeed support 웹사이트](https://www.litespeedtech.com/support)에서 제공한다.

### LoxiLB

[kube-loxilb](https://github.com/loxilb-io/kube-loxilb)는 [LoxiLB](https://github.com/loxilb-io)가 구현한 게이트웨이 API 및 쿠버네티스 서비스 로드 밸런서 명세 구현체로, 로드 밸런서 클래스, 고급 IPAM(공유 또는 전용) 등을 지원한다. kube-loxilb는 L4 서비스 로드 밸런서로서 [LoxiLB](https://github.com/loxilb-io/loxilb)를 인그레스(L7) 리소스를 위해 [loxilb-ingress](https://github.com/loxilb-io/loxilb-ingress)를 사용하여 게이트웨이 API 리소스를 관리한다.

간단한 단계로 게이트웨이 API와 함께 LoxiLB를 실행하려면 [빠른 시작 가이드](https://docs.loxilb.io/latest/gw-api/)를 참고하자.

### NGINX Gateway Fabric

[NGINX 게이트웨이 패브릭](https://github.com/nginx/nginx-gateway-fabric)은 [NGINX](https://nginx.org/)를 데이터 플레인으로 사용하는 게이트웨이 API의 구현체를 제공하는 오픈소스 프로젝트이다. 이 프로젝트의 목표는 쿠버네티스에서 실행되는 애플리케이션을 위한 HTTP 또는 TCP/UDP 로드 밸런서, 리버스 프록시 또는 API 게이트웨이를 구성하기 위해 코어 게이트웨이 API를 구현하는 것이다. [NGINX 문서](https://docs.nginx.com/nginx-gateway-fabric/) 웹사이트에서 종합적인 NGINX 게이트웨이 패브릭 사용자 문서를 찾을 수 있다.

지원되는 게이트웨이 API 리소스 및 기능 목록은 [게이트웨이 API 호환성](https://docs.nginx.com/nginx-gateway-fabric/overview/gateway-api-compatibility/) 문서를 확인하자.

NGINX 게이트웨이 패브릭에 대한 제안이 있거나 문제를 경험했다면 GitHub에서 [이슈를 생성](https://github.com/nginx/nginx-gateway-fabric/issues/new)하거나 [토론](https://github.com/nginx/nginx-gateway-fabric/discussions/new)을 부탁한다. 또한 [NGINX 커뮤니티 포럼](https://community.nginx.org/)에서 도움도 요청할 수 있다.

### ngrok Kubernetes Operator

[ngrok 쿠버네티스 오퍼레이터](https://github.com/ngrok/ngrok-operator)는 작년에 초기 지원을 추가한 이후로 게이트웨이 API의 전체 코어를 지원한다. 이것은 다음을 포함한다.

- 라우트: (HTTPRoute, TCPRoute, TLSRoute) 및 RouteMatches (Header, Path, 등)
- 필터: Header, Redirect, Rewrite 등
- 백엔드: 백엔드 Filters 및 가중치 기반 밸런싱
- ReferenceGrant: 멀티 테넌트 클러스터 처리를 위한 RBAC
- 게이트웨이 API가 충분히 유연하지 않은 경우, extensionRef 또는 어노테이션으로 트래픽 정책 설정

자세한 내용은 [docs](https://ngrok.com/docs/k8s/)를 참고하자. 기능 요청이나 버그 리포트는 [이슈 생성](https://github.com/ngrok/ngrok-operator/issues/new/choose)을 통해 제출을 부탁한다. 또한 [Slack](https://ngrokcommunity.slack.com/channels/general)에서 도움을 받을 수 있다.

### STUNner

[STUNner](https://github.com/l7mp/stunner)는 쿠버네티스용 오픈소스 클라우드 네이티브 WebRTC 미디어 게이트웨이이다. STUNner는 WebRTC 미디어 스트림을 쿠버네티스 클러스터로 원활하게 수신하기 위한 목적으로 설계되었으며, 간소화된 NAT 트래버설과 동적 미디어 라우팅을 제공한다. 동시에 STUNner는 대규모 실시간 통신 서비스에 대해 보안성과 모니터링 기능을 향상시킨다. STUNner의 데이터 플레인은 WebRTC 클라이언트를 위한 표준 규격의 TURN 서비스를 제공하며, 컨트롤 플레인은 게이트웨이 API의 일부를 지원한다.

현재 STUNner는 게이트웨이 API 명세의 `v1alpha2` 버전을 지원한다. WebRTC 미디어 수신을 위해 STUNner를 배포하고 사용하는 방법은 [설치 가이드](https://github.com/l7mp/stunner/blob/main/doc/INSTALL.md)를 확인하자. STUNner와 관련된 모든 질문, 의견 및 버그 리포트는 [STUNner 프로젝트](https://github.com/l7mp/stunner)로 보내주시기 바란다.

### Traefik Proxy

[Traefik Proxy](https://traefik.io)는 오픈 소스 클라우드 네이티브 애플리케이션 프록시이다.

Traefik 프록시는 현재 게이트웨이 API 명세의 `v1.4.0` 버전을 지원한다. 배포 및 사용 방법에 대한 자세한 정보는 [쿠버네티스 게이트웨이 프로바이더 문서](https://doc.traefik.io/traefik/v3.6/reference/install-configuration/providers/kubernetes/kubernetes-gateway)를 확인하자.
Traefik 프록시의 구현은 GRPCRoute와 같은 모든 HTTP 코어 및 일부 확장 호환성 테스트를 통과하며, 실험적 채널의 TCPRoute 및 TLSRoute 기능도 지원한다.

Traefik 프록시에 대한 도움과 지원을 받으려면, [이슈를 생성](https://github.com/traefik/traefik/issues/new/choose)하거나 [Traefik Labs 커뮤니티 포럼](https://community.traefik.io/c/traefik/traefik-v3/21)에서 도움을 요청하자.

### Tyk

[Tyk 게이트웨이](https://github.com/TykTechnologies/tyk)는 클라우드 네이티브 오픈소스 API 게이트웨이이다.

[Tyk.io](https://tyk.io) 팀은 게이트웨이 API 구현을 목표로 작업 중이며, 이 프로젝트의 진행 상황은 [여기](https://github.com/TykTechnologies/tyk-operator)에서 확인할 수 있다.

### WSO2 APK

[WSO2 APK](https://apk.docs.wso2.com/en/latest/)는 쿠버네티스 환경을 위해 특별히 설계된 API 관리 솔루션으로, API 관리를 위한 통합성, 유연성, 확장성을 조직에 제공한다.

WSO2 APK는 게이트웨이 API를 구현하며, 게이트웨이 및 HTTPRoute 기능을 포함한다. 또한, 사용자 정의 리소스(CR)를 통해 레이트 리밋팅, 인증/인가, 분석/관찰 가능성을 지원한다.

게이트웨이 API의 지원 버전과 기능에 대한 최신 정보는 [APK 게이트웨이 문서](https://apk.docs.wso2.com/en/latest/catalogs/kubernetes-crds/)를 참고하자. 질문이 있거나 기여하고 싶다면 자유롭게 [이슈 또는 풀 리퀘스트](https://github.com/wso2/apk)를 생성할 수 있다. 또한 [Discord 채널](https://discord.com/channels/955510916064092180/1113056079501332541)에서 우리와 소통하고 토론에 참여할 수 있다.

## 통합

이 섹션에서는 특정 통합을 위한 블로그 포스트, 문서 및 기타 게이트웨이 API 참조에 대한 구체적인 링크를 찾을 수 있다.

### Flagger

[Flagger](https://flagger.app)는 쿠버네티스에서 실행되는 애플리케이션의 릴리스 프로세스를 자동화하는 점진적 배포 도구이다.

Flagger는 게이트웨이 API를 사용하여 카나리 배포와 A/B 테스트를 자동화하는 데 사용할 수 있다. 게이트웨이 API의 `v1alpha2`와 `v1beta1` 명세를 모두 지원한다. 게이트웨이 API의 모든 구현과 함께 Flagger를 사용하려면 [이 튜토리얼](https://docs.flagger.app/tutorials/gatewayapi-progressive-delivery)을 참조한다.

### cert-manager

[cert-manager](https://cert-manager.io/)는 클라우드 네이티브 환경에서 인증서 관리를 자동화하기 위한 도구이다.

cert-manager는 게이트웨이 리소스를 위한 TLS 인증서를 생성할 수 있다. 이는 게이트웨이에 어노테이션을 추가하여 구성된다. 현재 게이트웨이 API의 `v1` 명세를 지원한다. 사용해보려면 [cert-manager 문서](https://cert-manager.io/docs/usage/gateway/)를 참조한다.

### Argo rollouts

[Argo Rollouts](https://argo-rollouts.readthedocs.io/en/stable/)는 쿠버네티스를 위한 점진적 배포 컨트롤러이다. 블루/그린 및 카나리와 같은 여러 고급 배포 방법을 지원한다. Argo Rollouts는 [플러그인](https://github.com/argoproj-labs/rollouts-gatewayapi-trafficrouter-plugin/)을 통해 게이트웨이 API를 지원한다.

### Knative

[Knative](https://knative.dev/)는 쿠버네티스 위에 구축된 서버리스 플랫폼이다. Knative Serving은 URL의 자동 관리, 리비전 간 트래픽 분할, 요청 기반 자동 스케일링(제로 스케일 포함), 자동 TLS 프로비저닝과 함께 상태 비저장 컨테이너를 실행하기 위한 간단한 API를 제공한다. Knative Serving은 플러그인 아키텍처를 통해 다중 HTTP 라우터를 지원하며, 이는 모든 Knative 기능이 지원되지 않아 현재 알파 단계에 있는 [게이트웨이 API 플러그인](https://github.com/knative-sandbox/net-gateway-api)을 포함한다.

### Kuadrant

[Kuadrant](https://kuadrant.io/)는 다른 게이트웨이 API 제공자와 통합되고 정책 연결을 통해 정책을 제공하는 오픈 소스 멀티 클러스터 게이트웨이 API 컨트롤러이다.

Kuadrant는 게이트웨이를 중앙에서 정의하고 모든 게이트웨이에 적용되는 DNS, TLS, 인증 및 레이트 리밋팅과 같은 정책을 연결하기 위한 게이트웨이 API를 지원한다.

Kuadrant는 Istio와 Envoy Gateway를 기본 게이트웨이 API 제공자로 지원하며, 향후 다른 게이트웨이 제공자와도 작동할 계획이다.

Kuadrant의 구현에 대한 도움과 지원을 받으려면, 자유롭게 [이슈를 생성](https://github.com/Kuadrant/kuadrant-operator/issues/new)하거나 [쿠버네티스 slack의 #kuadrant 채널](https://kubernetes.slack.com/archives/C05J0D0V525)에서 도움을 요청하자.

### OpenKruise Rollouts {#kruise-rollouts}
[OpenKruise Rollouts](https://openkruise.io/rollouts/introduction)는 쿠버네티스를 위한 플러그 앤 플레이 점진적 배포 컨트롤러이다. 블루/그린 및 카나리와 같은 여러 고급 배포 방법을 지원한다. OpenKruise Rollouts는 게이트웨이 API에 대한 내장 지원을 제공한다.

## 새 항목 추가

구현체는 자유롭게 PR을 만들어 이 페이지에 항목을 추가할 수 있다.
그러나 부분 호환 또는 호환 요구 사항을 충족하려면
구현체의 호환성 보고서 제출 PR이 병합되어 있어야 한다.

이 페이지에 새로 추가되는 항목에 대한 검토 프로세스의 일부로,
메인테이너가 호환성 수준을 확인하고 상태를 검증한다.

## 페이지 검토 정책

이 페이지는 활발히 개발되고 있으며 호환되는 게이트웨이 API 구현체를 보여주기 위한 것이며,
정기적인 검토 대상이다.

이러한 검토는 모든 게이트웨이 API 릴리스 이후 최소 1개월 후에 수행된다
(게이트웨이 API v1.3 릴리스부터 시작).

검토의 일부로, 메인테이너는 다음을 확인한다.

* 이 문서에서 위에 정의된 대로 **호환**(Conformant) 인 구현체
* 이 문서에서 위에 정의된 대로 **부분 호환**(Partially Conformant) 인 구현체

검토를 수행하는 메인테이너가 부분 호환 또는 호환 기준을 더 이상 충족하지 않는 구현체를 발견하거나
"비활성(Stale)" 상태인 구현체를 발견하면, 해당 메인테이너는 다음을 수행한다.

* 다른 메인테이너에게 알리고 비활성 및 제거 예정 구현체 목록에 대한 동의를 받는다.
* 이 페이지의 변경 사항이 포함된 드래프트 PR을 연다.
* #sig-network-gateway-api 채널에 게시하여 최소한 부분 호환이 아닌 구현체의
메인테이너가 게이트웨이 API 메인테이너에게 연락하여 구현체의 상태를 논의해야 함을 알린다.
이 기간을 "**답변 권리**(right-of-reply)" 기간이라 하며 최소 2주이고,
지연 합의(lazy consensus) 기간으로 기능한다.
* 답변 권리 기간 내에 응답하지 않는 구현체는
"비활성"으로 이동되거나, 이미 "비활성"인 경우 이 페이지에서 제거되어
상태가 하향 조정된다.

v1.4 페이지 검토부터 시작하는 페이지 검토 일정:

* 게이트웨이 API v1.4 릴리스 페이지 검토(실제 릴리스 이후 최소 1개월 후):
  메인테이너가 위 규칙에 따라 호환성 보고서를 제출하지 않은 구현체를
  "비활성"으로 이동시킨다. 또한 비활성으로 이동되는 구현체에게 이 규칙 변경에 대해 알린다.
  **현재 이 단계이다**
* 게이트웨이 API v1.5 릴리스 페이지 검토(실제 릴리스 이후 최소 1개월 후):
  메인테이너가 페이지 검토 프로세스를 다시 수행하여 여전히 비활성인
  구현체를 제거한다(답변 권리 기간 이후).
* 게이트웨이 API v1.6 릴리스 페이지 검토(실제 릴리스 이후 최소 1개월 후):
  비활성 카테고리를 제거하고, 구현체 메인테이너는
  각 검토 시 또는 답변 권리 기간 내에 최소한 부분 호환이어야 하며,
  그렇지 않으면 구현체 페이지에서 제거된다.

이는 게이트웨이 API v1.6 릴리스 이후에는 최소한 부분 호환(Partially Conformant)
호환성 보고서를 제출하지 않으면 이 페이지에 구현체를 추가할 수 없음을 의미한다.
