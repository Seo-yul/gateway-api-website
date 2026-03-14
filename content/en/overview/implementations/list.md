---
title: "List"
description: "Downstream implementations and integrations of Gateway API"
weight: 1
type: docs
---

This document tracks downstream implementations and integrations of Gateway API
and provides status and resource references for them.

Implementors and integrators of Gateway API are encouraged to update this
document with status information about their implementations, the versions they
cover, and documentation to help users get started. This status information should
be no longer than a few paragraphs.

## Conformance levels

There are three levels of Gateway API conformance:

### Conformant implementations

These implementations have submitted at least one conformance report that has passes for:

  * All core conformance tests for at least one combination of Route type and
    Profile
  * All claimed Extended features

for one of the two (2) most recent Gateway API releases.

So, it's conformant to support Mesh + HTTPRoute, or Gateway + HTTPRoute, or
Gateway + TLSRoute, or Gateway + Mesh + HTTPRoute, plus any extended features
the implementation claims. But implementations _must_ support at least one
Profile and one Route type in that profile, and must pass all Core conformance
tests for that Profile and Route type in addition to all claimed Extended
features.

### Partially Conformant implementations

These implementations are aiming for full conformance but are not currently
achieving it. They have submitted at least one conformance report passing some
of the tests to be Conformant (as above) for one of the three (3) most recent
Gateway API releases. Note that the requirements to be considered "partially
conformant" may be tightened in a future release of Gateway API.

### Stale implementations

These implementations may not be being actively developed and will be removed
from this page on the next page review unless they submit a conformance report
moving them to one of the other categories.

Page reviews are performed at least one month after every Gateway API release,
with the first being performed after the release of Gateway API v1.3, in late
June 2025. Following the Gateway API v1.5 review process, due in mid-2026,
stale implementations will no longer be listed.

## Implementation profiles

Implementations also generally fall into two categories, which are called
_profiles_:

* **Gateway** controllers reconcile the Gateway resource and are intended to
handle north-south traffic, mainly concerned with coming from outside the
cluster to inside.
* **Mesh** controllers reconcile Service resources with HTTPRoutes attached
and are intended to handle east-west traffic, within the same cluster or
set of clusters.

Each profile has a set of conformance tests associated with it, that lay out
the expected behavior for implementations to be conformant (as above).

Implementations may also fit both profiles.

## Integrations

Also listed on this page are **integrations**, which are other software
projects that are able to make use of Gateway API resources to perform
other functions (like managing DNS or creating certificates).

{{< note >}}
This page contains links to third party projects that provide functionality
required for Gateway API to work. The Gateway API project authors aren't
responsible for these projects, which are listed alphabetically within their
class.
{{< /note >}}

{{< note >}}
**Compare extended supported features across implementations**

[View a table to quickly compare supported features of projects](comparisons/v1.4/). These outline Gateway controller implementations that have passed core conformance tests, and focus on extended conformance features that they have implemented. These tables will be generated and uploaded to the site once at least 3 implementations have uploaded their conformance reports under the [conformance reports](https://github.com/kubernetes-sigs/gateway-api/tree/main/conformance/reports).
{{< /note >}}

## Gateway Controller Implementation Status {#gateways}

### Conformant
- [Agent Gateway](#agent-gateway-with-kgateway)
- [Airlock Microgateway](#airlock-microgateway)
- [Cilium](#cilium)
- [Envoy Gateway](#envoy-gateway) (GA)
- [Istio](#istio) (GA)
- [kgateway](#kgateway) (GA)
- [NGINX Gateway Fabric](#nginx-gateway-fabric) (GA)
- [Traefik Proxy](#traefik-proxy) (GA)

### Partially Conformant

- [AWS Load Balancer Controller](#aws-load-balancer-controller) (GA)
- [Azure Application Gateway for Containers](#azure-application-gateway-for-containers) (GA)
- [Contour](#contour) (GA)
- [Gloo Gateway](#gloo-gateway) (GA)
- [Google Kubernetes Engine](#google-kubernetes-engine) (GA)
- [Gravitee Kubernetes Operator](#gravitee-kubernetes-operator) (GA)
- [Kong Ingress Controller](#kong-kubernetes-ingress-controller) (GA)
- [Kong Gateway Operator](#kong-gateway-operator) (GA)
- [Kubvernor](#kubvernor)(work in progress)

### Stale

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
- [Tyk](#tyk) (work in progress)
- [WSO2 APK](#wso2-apk) (GA)

## Service Mesh Implementation Status {#meshes}

### Conformant
- [Alibaba Cloud Service Mesh](#alibaba-cloud-service-mesh) (GA)
- [Istio](#istio) (GA)
- [Linkerd](#linkerd) (GA)
- [Cilium](#cilium) (GA)

### Stale
- [Google Cloud Service Mesh](#google-cloud-service-mesh) (GA)
- [Kuma](#kuma) (GA)

## Integrations {#integrations}

- [Flagger](#flagger) (public preview)
- [cert-manager](#cert-manager) (alpha)
- [argo-rollouts](#argo-rollouts) (alpha)
- [Knative](#knative) (alpha)
- [Kuadrant](#kuadrant) (GA)
- [kruise-rollouts](#kruise-rollouts) (alpha)

## Implementations

In this section you will find specific links to blog posts, documentation and other Gateway API references for specific implementations.

### Acnodal EPIC
[EPIC](https://www.epic-gateway.org/) is an Open Source External Gateway platform designed and built with Kubernetes. It consists of the Gateway Cluster, k8s Gateway controller, a stand alone Linux Gateway controller and the Gateway Service Manager. Together they create a platform for providing Gateway services to cluster users. Each gateway consists of multiple Envoy instances running on the gateway cluster not the workload clusters. The Gateway Service Manager is a simple user management and UI that can be used to implement Gateway-as-a-Service infrastructure for public and private clusters, and integrate non-k8s endpoints.

- [Documentation](https://www.epic-gateway.org/)
- [Source Repo](https://github.com/epic-gateway)

### Agentgateway

[Agentgateway](https://agentgateway.dev/) is an open source Gateway API implementation hosted as a part of the Linux Foundation, focusing on AI use cases, including LLM consumption, LLM serving, agent-to-agent (A2A), agent-to-tool (MCP), as well as traditional TCP/HTTP traffic serving.
It is the first and only proxy designed specifically for the Kubernetes Gateway API, powered by a high performance and scalable Rust dataplane implementation.

### Airlock Microgateway

[Airlock Microgateway](https://www.airlock.com/en/secure-access-hub/components/microgateway) is a Kubernetes native WAAP (Web Application and API Protection, formerly known as WAF) solution optimized for Kubernetes environments and certified for Red Hat OpenShift.
Modern application security is embedded in the development workflow and follows DevSecOps paradigms.
Airlock Microgateway protects your applications and microservices with the tried-and-tested Airlock security features against attacks, while also providing a high degree of scalability.

#### Features
- Comprehensive WAAP (formerly known as WAF) with security features like Deny Rules to protect against known attacks (OWASP Top 10), header filtering, JSON parsing, OpenAPI specification enforcement, and GraphQL schema validation
- Identity aware proxy which makes it possible to enforce authentication using JWT authentication or OIDC, with OAuth 2.0 Token Introspection and Token Exchange for continuous validation and secure delegation across services
- Reverse proxy functionality with request routing rules, TLS termination and remote IP extraction
- Easy-to-use Grafana dashboards which provide valuable insights in allowed and blocked traffic and other metrics

#### Documentation and links
- [Product documentation](https://docs.airlock.com/microgateway/latest)
- [Gateway specific documentation](https://docs.airlock.com/microgateway/latest/?topic=MGW-00000142)
- Check our [Airlock community forum](https://forum.airlock.com/) and [support process](https://techzone.ergon.ch/support-process) for support.

### Alibaba Cloud Service Mesh

[Alibaba Cloud Service Mesh (ASM)](https://www.alibabacloud.com/help/en/asm/product-overview/what-is-asm) provides a fully managed service mesh platform that is compatible with the community Istio. It simplifies service governance, including traffic routing and split management between service calls, authentication security for inter-service communication, and mesh observability capabilities, thereby greatly reducing the workload of development and operations.

### Amazon Elastic Kubernetes Service

[Amazon Elastic Kubernetes Service (EKS)](https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html) is a managed service that you can use to run Kubernetes on AWS without needing to install, operate, and maintain your own Kubernetes control plane or nodes. EKS's implementation of the Gateway API is through [AWS Gateway API Controller](https://github.com/aws/aws-application-networking-k8s) which provisions [Amazon VPC Lattice](https://aws.amazon.com/vpc/lattice/) Resources for gateway(s), HTTPRoute(s) in EKS clusters.

### APISIX

[Apache APISIX](https://apisix.apache.org/) is a dynamic, real-time, high-performance API Gateway. APISIX provides rich traffic management features such as load balancing, dynamic upstream, canary release, circuit breaking, authentication, observability, and more.

APISIX currently supports Gateway API `v1beta1` version of the specification for its [Apache APISIX Ingress Controller](https://github.com/apache/apisix-ingress-controller).

### AWS Load Balancer Controller

[AWS Load Balancer Controller](https://github.com/kubernetes-sigs/aws-load-balancer-controller) manages AWS Elastic Load Balancers for Kubernetes clusters. The controller provisions AWS Application Load Balancers (ALB) when you create a Kubernetes Ingress and AWS Network Load Balancers (NLB) when you create a Kubernetes Service of type LoadBalancer.

Gateway API support is GA for both Layer 4 (L4) and Layer 7 (L7) routing, enabling customers to provision and manage AWS NLBs and ALBs directly from Kubernetes clusters using the extensible Gateway API.

See the [AWS Load Balancer Controller documentation](https://kubernetes-sigs.github.io/aws-load-balancer-controller/) for information on how to deploy and use the Gateway API implementation.

### Avi Kubernetes Operator

[Avi Kubernetes Operator (AKO)](https://techdocs.broadcom.com/us/en/vmware-security-load-balancing/avi-load-balancer/avi-kubernetes-operator/2-1.html) provides L4-L7 load-balancing using VMware AVI Advanced Load Balancer.

Starting with AKO version v2.1.1, Gateway API version v1.3.0 is supported. It implements v1 version of Gateway API specification supporting GatewayClass, Gateway and HTTPRoute objects.

Documentation to deploy and use AKO Gateway API can be found at [Avi Kubernetes Operator Gateway API](https://techdocs.broadcom.com/us/en/vmware-security-load-balancing/avi-load-balancer/avi-kubernetes-operator/2-1/avi-kubernetes-operator-guide-2-1/gateway-api/gateway-api-v1.html).

### Azure Application Gateway for Containers

[Application Gateway for Containers](https://aka.ms/appgwcontainers/docs) is a managed application (layer 7) load balancing solution, providing dynamic traffic management capabilities for workloads running in a Kubernetes cluster in Azure. Follow the [quickstart guide](https://learn.microsoft.com/azure/application-gateway/for-containers/quickstart-deploy-application-gateway-for-containers-alb-controller) to deploy the ALB controller and get started with Gateway API.

### Cilium

[Cilium](https://cilium.io) is an eBPF-based networking, observability and security
solution for Kubernetes and other networking environments. It includes [Cilium
Service Mesh](https://docs.cilium.io/en/stable/gettingstarted/#service-mesh), a highly efficient mesh data plane that can
be run in [sidecarless mode](https://isovalent.com/blog/post/cilium-service-mesh/) to dramatically improve
performance, and avoid the operational complexity of sidecars. Cilium also
supports the sidecar proxy model, offering choice to users.
Cilium supports Gateway API, passing conformance for v1.4.0 as of Cilium 1.19

Cilium is open source and is a CNCF Graduated project.

If you have questions about Cilium Service Mesh the #service-mesh channel on
[Cilium Slack](https://slack.cilium.io) is a good place to start. For contributing to the development
effort, check out the #development channel or join our [weekly developer meeting](https://github.com/cilium/cilium#weekly-developer-meeting).

### Contour

[Contour](https://projectcontour.io) is a CNCF open source Envoy-based ingress controller for Kubernetes.

Contour v1.31.0 implements Gateway API v1.2.1.
All Standard channel v1 API group resources (GatewayClass, Gateway, HTTPRoute, ReferenceGrant), plus most v1alpha2 API group resources (TLSRoute, TCPRoute, GRPCRoute, ReferenceGrant, and BackendTLSPolicy) are supported.
Contour's implementation passes most core extended Gateway API conformance tests included in the v1.2.1 release.

See the [Contour Gateway API Guide](https://projectcontour.io/docs/1.30/guides/gateway-api/) for information on how to deploy and use Contour's Gateway API implementation.

For help and support with Contour's implementation, [create an issue](https://github.com/projectcontour/contour/issues/new/choose) or ask for help in the [#contour channel on Kubernetes slack](https://kubernetes.slack.com/archives/C8XRH2R4J).

### Easegress

[Easegress](https://megaease.com/easegress/) is a Cloud Native traffic orchestration system.

It can function as a sophisticated modern gateway, a robust distributed cluster, a flexible traffic orchestrator, or even an accessible service mesh.

Easegress currently supports Gateway API `v1beta1` version of the specification by [GatewayController](https://github.com/megaease/easegress/blob/main/docs/04.Cloud-Native/4.2.Gateway-API.md).

### Emissary-Ingress (Ambassador API Gateway)

[Emissary-Ingress](https://www.getambassador.io/docs/edge-stack) (formerly known as Ambassador API Gateway) is an open source CNCF project that
provides an ingress controller and API gateway for Kubernetes built on top of [Envoy Proxy](https://envoyproxy.io).
See [here](https://www.getambassador.io/docs/edge-stack/latest/topics/using/gateway-api/) for more details on using the Gateway API with Emissary.

### Envoy Gateway

[Envoy Gateway](https://gateway.envoyproxy.io/) is an [Envoy](https://github.com/envoyproxy) subproject for managing Envoy-based application gateways. The supported
APIs and fields of the Gateway API are outlined [here](https://gateway.envoyproxy.io/docs/tasks/quickstart/).
Use the [quickstart](https://gateway.envoyproxy.io/docs/tasks/quickstart) to get Envoy Gateway running with Gateway API in a
few simple steps.

### Flomesh Service Mesh (FSM)

[Flomesh Service Mesh](https://github.com/flomesh-io/fsm) is a community driven lightweight service mesh for Kubernetes East-West and North-South traffic management. Flomesh uses ebpf for layer4 and pipy proxy for layer7 traffic management. Flomesh comes bundled with a load balancer, cross-cluster service registration/discovery and it supports multi-cluster networking. It supports `Ingress` (and as such is an "Ingress controller") and Gateway API.

FSM support of Gateway API is built on top [Flomesh Gateway API](https://github.com/flomesh-io/fgw) and it currently supports Kubernetes Gateway API version v0.7.1 with support for `v0.8.0` currently in progress.

- [FSM Kubernetes Gateway API compatibility matrix](https://github.com/flomesh-io/fsm/blob/main/docs/gateway-api-compatibility.md)
- [How to use Gateway API support in FSM](https://github.com/flomesh-io/fsm/blob/main/docs/tests/gateway-api/README.md)

### Gloo Gateway

[Gloo Gateway](https://docs.solo.io/gateway/latest/) by [Solo.io](https://www.solo.io) is a feature-rich, Kubernetes-native ingress controller and next-generation API gateway.
Gloo Gateway brings the full power and community support of Gateway API to its existing control-plane implementation.

The Gloo Gateway ingress controller passes all the core Gateway API conformance tests in the v1.1.0 release for the GATEWAY_HTTP conformance
profile except `HTTPRouteServiceTypes`.

### Google Cloud Service Mesh

[Google Kubernetes Engine (GKE)](https://cloud.google.com/kubernetes-engine) is a managed Kubernetes platform offered
by Google Cloud.

GKE's implementation of Gateway For Mesh (GAMMA) is through the [Cloud Service Mesh](https://cloud.google.com/products/service-mesh).

Google Cloud Service Mesh supports [Envoy-based sidecar mesh](https://cloud.google.com/service-mesh/docs/gateway/set-up-envoy-mesh) and [Proxyless-GRPC](https://cloud.google.com/service-mesh/docs/gateway/proxyless-grpc-mesh) (using GRPCRoute).

### Google Kubernetes Engine

[Google Kubernetes Engine (GKE)](https://cloud.google.com/kubernetes-engine) is a managed Kubernetes platform offered
by Google Cloud. GKE's implementation of the Gateway API is through the [GKE
Gateway controller](https://cloud.google.com/kubernetes-engine/docs/concepts/gateway-api) which provisions Google Cloud Load Balancers
for Pods in GKE clusters.

The GKE Gateway controller supports weighted traffic splitting, mirroring,
advanced routing, multi-cluster load balancing and more. See the docs to deploy
[private or public Gateways](https://cloud.google.com/kubernetes-engine/docs/how-to/deploying-gateways) and also [multi-cluster
Gateways](https://cloud.google.com/kubernetes-engine/docs/how-to/deploying-multi-cluster-gateways).

The GKE Gateway controller passes all the core Gateway API conformance tests in the
v1.4.0 release for the GATEWAY_HTTP conformance profile except `HTTPRouteHostnameIntersection`.

### Gravitee Kubernetes Operator

The [Gravitee Kubernetes Operator](https://documentation.gravitee.io/gravitee-kubernetes-operator-gko) (GKO) lets you manage [Gravitee](https://www.gravitee.io/) APIs, applications, and other assets in a Kubernetes-native and declarative way.

The Gravitee Kubernetes Operator provides partial conformance for Gateway - HTTP features in version 4.10.3. It does not support matching rules across routes. These feature will be introduced in a future release.

For support, feedback, or to engage in a discussion about the Gravitee Kubernetes Operator, please feel free to submit an [issue](https://github.com/gravitee-io/issues/issues) or visit our community [forum](https://community.gravitee.io/c/support/gravitee-kubernetes-operator-gko/26).

### HAProxy Ingress

[HAProxy Ingress](https://haproxy-ingress.github.io/) is a community driven ingress controller implementation for HAProxy.

HAProxy Ingress v0.13 partially supports the Gateway API's v1alpha1 specification. See the [controller's Gateway API documentation](https://haproxy-ingress.github.io/docs/configuration/gateway-api/) to get informed about conformance and roadmap.

### HAProxy Kubernetes Ingress Controller

HAProxy Kubernetes Ingress Controller is an open-source project maintained by HAProxy Technologies that provides fast and efficient traffic management, routing, and observability for Kubernetes. It has built-in support for the Gateway API since version 1.10. The same deployment of the ingress controller will allow you to use both the Ingress API and Gateway API. See the [documentation](https://www.haproxy.com/documentation/kubernetes-ingress/gateway-api/enable-gateway-api/) for more details. In the [GitHub repository](https://github.com/haproxytech/kubernetes-ingress/blob/master/documentation/gateway-api.md), you will also find additional information about supported API resources.

### HashiCorp Consul

[Consul](https://consul.io), by [HashiCorp](https://www.hashicorp.com), is an open source control plane for multi-cloud networking. A single Consul deployment can span bare metal, VM and container environments.

Consul service mesh works on any Kubernetes distribution, connects multiple clusters, and Consul CRDs provide a Kubernetes native workflow to manage traffic patterns and permissions in the mesh. [Consul API Gateway](https://www.consul.io/docs/api-gateway) supports Gateway API for managing North-South traffic.

Please see the [Consul API Gateway documentation](https://www.consul.io/docs/api-gateway) for current information on the supported version and features of the Gateway API.

### Istio

[Istio](https://istio.io) is an open source [service mesh](https://istio.io/latest/docs/concepts/what-is-istio/#what-is-a-service-mesh) and gateway implementation.

A minimal install of Istio can be used to provide a fully compliant
implementation of the Kubernetes Gateway API for cluster ingress traffic
control. For service mesh users, Istio also fully supports the GAMMA
initiative's Gateway API support for east-west traffic management within the mesh.

Much of Istio's documentation, including all of the [ingress tasks](https://istio.io/latest/docs/tasks/traffic-management/ingress/) and several mesh-internal traffic management tasks, already includes parallel instructions for
configuring traffic using either the Gateway API or the Istio configuration API.
Check out the [Gateway API task](https://istio.io/latest/docs/tasks/traffic-management/ingress/gateway-api/) for more information about the Gateway API implementation in Istio.

### kgateway

The [kgateway](https://kgateway.dev/docs) project is a feature-rich, Kubernetes-native ingress controller and next-generation API gateway.
It is focused on maintaining a great HTTP experience, extending features for advanced routing in scenarios such as AI and MCP gateways, and interoperating with a service mesh such as Istio in both ambient and sidecar modes.
This focus means that you can easily configure a set of Envoy instances that are reasonably distributed in a performant way across many north-south and east-west use cases.

Kgateway is generally available with its 2.0 release.

### Kong Kubernetes Ingress Controller

[Kong](https://konghq.com) is an open source API Gateway built for hybrid and multi-cloud environments.

The [Kong Kubernetes Ingress Controller (KIC)](https://github.com/kong/kubernetes-ingress-controller) can be used to configure unmanaged Gateways. See the [Gateway API Guide](https://docs.konghq.com/kubernetes-ingress-controller/latest/guides/using-gateway-api/) for usage information.

For help and support with Kong Kubernetes Ingress Controller please feel free to [create an issue](https://github.com/Kong/kubernetes-ingress-controller/issues/new) or a [discussion](https://github.com/Kong/kubernetes-ingress-controller/discussions/new). You can also ask for help in the [#kong channel on Kubernetes slack](https://kubernetes.slack.com/archives/CDCA87FRD).

### Kong Gateway Operator

[Kong](https://konghq.com) is an open source API Gateway built for hybrid and multi-cloud environments.

The [Kong Gateway operator (KGO)](https://docs.konghq.com/gateway-operator/latest/) can be used to configure managed Gateways and orchestrate instances of Kong Kubernetes Ingress Controllers.

For help and support with Kong Gateway operator please feel free to [create an issue](https://github.com/Kong/gateway-operator/issues/new) or a [discussion](https://github.com/Kong/gateway-operator/discussions/new). You can also ask for help in the [#kong channel on Kubernetes slack](https://kubernetes.slack.com/archives/CDCA87FRD).

### Kubvernor
[Kubvernor](https://github.com/kubvernor/kubvernor) is an open-source, highly experimental implementation of API controller in Rust programming language. Currently, Kubvernor supports Envoy Proxy. The project aims to be as generic as possible so Kubvernor can be used to manage/deploy different gateways (Envoy, Nginx, HAProxy, etc.).

### Kuma

[Kuma](https://kuma.io) is an open source service mesh.

Kuma implements the Gateway API specification for the Kuma built-in, Envoy-based Gateway with a beta stability guarantee. Check the [Gateway API Documentation](https://kuma.io/docs/latest/using-mesh/managing-ingress-traffic/gateway-api/) for information on how to set up a Kuma built-in gateway using the Gateway API.

Kuma 2.3 and later support the GAMMA initiative's Gateway API support for east-west traffic management within the mesh.

### Linkerd

[Linkerd](https://linkerd.io/) is the first CNCF graduated [service mesh](https://buoyant.io/service-mesh-manifesto).
It is the only major mesh not based on Envoy, instead relying on a
purpose-built Rust micro-proxy to bring security, observability, and
reliability to Kubernetes, without the complexity.

Linkerd 2.14 and later support the GAMMA initiative's Gateway API support for east-west traffic management within the mesh.

### LiteSpeed Ingress Controller

The [LiteSpeed Ingress Controller](https://litespeedtech.com/products/litespeed-web-adc/features/litespeed-ingress-controller) uses the LiteSpeed WebADC controller to operate as an Ingress Controller and Load Balancer to manage your traffic on your Kubernetes cluster. It implements the full core Gateway API including Gateway, GatewayClass, HTTPRoute and ReferenceGrant and the Gateway functions of cert-manager. Gateway is fully integrated into the LiteSpeed Ingress Controller.

- [Product documentation](https://docs.litespeedtech.com/cloud/kubernetes/).
- [Gateway specific documentation](https://docs.litespeedtech.com/cloud/kubernetes/gateway).
- Full support is available on the [LiteSpeed support web site](https://www.litespeedtech.com/support).

### LoxiLB

[kube-loxilb](https://github.com/loxilb-io/kube-loxilb) is [LoxiLB's](https://github.com/loxilb-io) implementation of Gateway API and kubernetes service load-balancer spec which includes support for load-balancer class, advanced IPAM (shared or exclusive) etc. kube-loxilb manages Gateway API resources with [LoxiLB](https://github.com/loxilb-io/loxilb) as L4 service LB and [loxilb-ingress](https://github.com/loxilb-io/loxilb-ingress) for Ingress(L7) resources.

Follow the [quickstart guide](https://docs.loxilb.io/latest/gw-api/) to get LoxiLB running with Gateway API in a few simple steps.

### NGINX Gateway Fabric

[NGINX Gateway Fabric](https://github.com/nginx/nginx-gateway-fabric) is an open-source project that provides an implementation of the Gateway API using [NGINX](https://nginx.org/) as the data plane. The goal of this project is to implement the core Gateway API to configure an HTTP or TCP/UDP load balancer, reverse-proxy, or API gateway for applications running on Kubernetes. You can find the comprehensive NGINX Gateway Fabric user documentation on the [NGINX Documentation](https://docs.nginx.com/nginx-gateway-fabric/) website.

For a list of supported Gateway API resources and features, see the [Gateway API Compatibility](https://docs.nginx.com/nginx-gateway-fabric/overview/gateway-api-compatibility/) doc.

If you have any suggestions or experience issues with NGINX Gateway Fabric, please [create an issue](https://github.com/nginx/nginx-gateway-fabric/issues/new) or a [discussion](https://github.com/nginx/nginx-gateway-fabric/discussions/new) on GitHub. You can also ask for help in the [NGINX Community Forum](https://community.nginx.org/).

### ngrok Kubernetes Operator

[ngrok Kubernetes Operator](https://github.com/ngrok/ngrok-operator) After adding preliminary support last year, the ngrok Kubernetes Operator supports the entire core Gateway API. This includes:

- Routes (HTTPRoute, TCPRoute, TLSRoute) + RouteMatches (Header, Path, +more)
- Filters: Header, Redirect, Rewrite + more
- Backends: Backend Filters + Weighted balancing
- ReferenceGrant: RBAC for multi-tenant clusters handling
- Traffic Policy as an extensionRef or annotation when the Gateway API isn't flexible enough

You can read our [docs](https://ngrok.com/docs/k8s/) for more information. If you have any feature requests or bug reports, please [create an issue](https://github.com/ngrok/ngrok-operator/issues/new/choose). You can also reach out for help on [Slack](https://ngrokcommunity.slack.com/channels/general)

### STUNner

[STUNner](https://github.com/l7mp/stunner) is an open source cloud-native WebRTC media gateway for Kubernetes. STUNner is purposed specifically to facilitate the seamless ingestion of WebRTC media streams into a Kubernetes cluster, with simplified NAT traversal and dynamic media routing. Meanwhile, STUNner provides improved security and monitoring for large-scale real-time communications services. The STUNner dataplane exposes a standards compliant TURN service to WebRTC clients, while the control plane supports a subset of the Gateway API.

STUNner currently supports version `v1alpha2` of the Gateway API specification. Check the [install guide](https://github.com/l7mp/stunner/blob/main/doc/INSTALL.md) for information on how to deploy and use STUNner for WebRTC media ingestion. Please direct all questions, comments and bug-reports related to STUNner to the [STUNner project](https://github.com/l7mp/stunner).

### Traefik Proxy

[Traefik Proxy](https://traefik.io) is an open source cloud-native application proxy.

Traefik Proxy currently supports version `v1.4.0` of the Gateway API specification, check the [Kubernetes Gateway Provider Documentation](https://doc.traefik.io/traefik/v3.6/reference/install-configuration/providers/kubernetes/kubernetes-gateway) for more information on how to deploy and use it.
Traefik Proxy's implementation passes all HTTP core and some extended conformance tests, like GRPCRoute, but also supports TCPRoute and TLSRoute features from the Experimental channel.

For help and support with Traefik Proxy, [create an issue](https://github.com/traefik/traefik/issues/new/choose) or ask for help in the [Traefik Labs Community Forum](https://community.traefik.io/c/traefik/traefik-v3/21).

### Tyk

[Tyk Gateway](https://github.com/TykTechnologies/tyk) is a cloud-native, open source, API Gateway.

The [Tyk.io](https://tyk.io) team is working towards an implementation of the Gateway API. You can track progress of this project [here](https://github.com/TykTechnologies/tyk-operator).

### WSO2 APK

[WSO2 APK](https://apk.docs.wso2.com/en/latest/) is a purpose-built API management solution tailored for Kubernetes environments, delivering seamless integration, flexibility, and scalability to organizations in managing their APIs.

WSO2 APK implements the Gateway API, encompassing Gateway and HTTPRoute functionalities. Additionally, it provides support for rate limiting, authentication/authorization, and analytics/observability through the use of Custom Resources (CRs).

For up-to-date information on the supported version and features of the Gateway API, please refer to the [APK Gateway documentation](https://apk.docs.wso2.com/en/latest/catalogs/kubernetes-crds/). If you have any questions or would like to contribute, feel free to create [issues or pull requests](https://github.com/wso2/apk). Join our [Discord channel](https://discord.com/channels/955510916064092180/1113056079501332541) to connect with us and engage in discussions.

## Integrations

In this section you will find specific links to blog posts, documentation and other Gateway API references for specific integrations.

### Flagger

[Flagger](https://flagger.app) is a progressive delivery tool that automates the release process for applications running on Kubernetes.

Flagger can be used to automate canary deployments and A/B testing using Gateway API. It supports both the `v1alpha2` and `v1beta1` spec of Gateway API. You can refer to [this tutorial](https://docs.flagger.app/tutorials/gatewayapi-progressive-delivery) to use Flagger with any implementation of Gateway API.

### cert-manager

[cert-manager](https://cert-manager.io/) is a tool to automate certificate management in cloud native environments.

cert-manager can generate TLS certificates for Gateway resources. This is configured by adding annotations to a Gateway. It currently supports the `v1` spec of Gateway API. You can refer to the [cert-manager docs](https://cert-manager.io/docs/usage/gateway/) to try it out.

### Argo rollouts

[Argo Rollouts](https://argo-rollouts.readthedocs.io/en/stable/) is a progressive delivery controller for Kubernetes. It supports several advanced deployment methods such as blue/green and canaries. Argo Rollouts supports the Gateway API via [a plugin](https://github.com/argoproj-labs/rollouts-gatewayapi-trafficrouter-plugin/).

### Knative

[Knative](https://knative.dev/) is a serverless platform built on Kubernetes. Knative Serving provides a simple API for running stateless containers with automatic management of URLs, traffic splitting between revisions, request-based autoscaling (including scale to zero), and automatic TLS provisioning. Knative Serving supports multiple HTTP routers through a plugin architecture, including a [gateway API plugin](https://github.com/knative-sandbox/net-gateway-api) which is currently in alpha as not all Knative features are supported.

### Kuadrant

[Kuadrant](https://kuadrant.io/) is an open source multi cluster Gateway API controller that integrates with and provides policies via policy attachment to other Gateway API providers.

Kuadrant supports Gateway API for defining gateways centrally and attaching policies such as DNS, TLS, Auth and Rate Limiting that apply to all of your Gateways.

Kuadrant works with both Istio and Envoy Gateway as underlying Gateway API providers, with plans to work with other gateway providers in future.

For help and support with Kuadrant's implementation please feel free to [create an issue](https://github.com/Kuadrant/kuadrant-operator/issues/new) or ask for help in the [#kuadrant channel on Kubernetes slack](https://kubernetes.slack.com/archives/C05J0D0V525).

### OpenKruise Rollouts
[OpenKruise Rollouts](https://openkruise.io/rollouts/introduction) is a plugin-n-play progressive delivery controller for Kubernetes. It supports several advanced deployment methods such as blue/green and canaries. OpenKruise Rollouts has built-in support for the Gateway API.

## Adding new entries

Implementations are free to make a PR to add their entry to this page; however,
in order to meet the requirements for being Partially Conformant or Conformant,
the implementation must have had a conformance report submission PR merged.

Part of the review process for new additions to this page is that a maintainer
will check the conformance level and verify the state.

## Page Review Policy

This page is intended to showcase actively developed and conformant implementations
of Gateway API, and so is subject to regular reviews.

These reviews are performed at least one month after every Gateway API release
(starting with the Gateway API v1.3 release).

As part of the review, a maintainer will check:

* which implementations are **Conformant** - as defined above in this document.
* which implementations are **Partially Conformant**, as defined above in this
  document.

If the maintainer performing the review finds that there are implementations
that no longer satisfy the criteria for Partially Conformant or Conformant, or
finds implementations that are in the "Stale" state, then that maintainer will:

* Inform the other maintainers and get their agreement on the list of stale and
to-be-removed implementations
* Open a draft PR with the changes to this page.
* Post on the #sig-network-gateway-api channel informing the maintainers of
implementations that are no longer at least partially conformant should contact
the Gateway API maintainers to discuss the implementation's status. This period
is called the "**right-of-reply**" period, is at least two weeks long, and functions
as a lazy consensus period.
* Any implementations that do not respond within the right-of-reply period will be
downgraded in status, either by being moved to "Stale", or being removed
from this page if they are already "Stale".

Page review timeline, starting with the v1.4 Page Review:

* Gateway API v1.4 release Page Review (at least one month after the actual
  release): a maintainer will move anyone who hasn't submitted a conformance
  report using the rules above to "Stale". They will also contact anyone who
  moves to Stale to inform them about this rule change.
  **You are here**
* Gateway API v1.5 release Page Review (at least one month after the actual
  release): A maintainer will perform the Page Review process again, removing
  any implementations that are still Stale (after a right-of-reply period).
* Gateway API v1.6 release Page Review (at least one month after the actual
  release): We will remove the Stale category, and implementation maintainers
  will need to be at least partially conformant on each review, or during the
  right-of-reply period, or be removed from the implementations page.

This means that, after the Gateway API v1.6 release, implementations cannot be
added to this page unless they have submitted at least a Partially Conformant
conformance report.
