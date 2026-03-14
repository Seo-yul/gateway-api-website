---
title: "Gateway API Graduates to Beta"
date: 2022-07-13
description: "We are excited to announce the v0.5.0 release of Gateway API. For the first time, several of our most important Gateway API resources are graduating to beta."
type: docs
---

We are excited to announce the v0.5.0 release of Gateway API. For the first
time, several of our most important Gateway API resources are graduating to
beta. Additionally, we are starting a new initiative to explore how Gateway API
can be used for mesh and introducing new experimental concepts such as URL
rewrites. We'll cover all of this and more below.

## What is Gateway API?

Gateway API is a collection of resources centered around [Gateway](/reference/api-types/gateway/)
resources (which represent the underlying network gateways / proxy servers) to enable
robust Kubernetes service networking through expressive, extensible and
role-oriented interfaces that are implemented by many vendors and have broad
industry support.

Originally conceived as a successor to the well known [Ingress](https://kubernetes.io/docs/reference/kubernetes-api/service-resources/ingress-v1/) API, the
benefits of Gateway API include (but are not limited to) explicit support for
many commonly used networking protocols (e.g. `HTTP`, `TLS`, `TCP`, `UDP`) as
well as tightly integrated support for Transport Layer Security (TLS). The
`Gateway` resource in particular enables implementations to manage the lifecycle
of network gateways as a Kubernetes API.

If you're an end-user interested in some of the benefits of Gateway API we
invite you to jump in and find an implementation that suits you. At the time of
this release there are over a dozen [implementations](/overview/implementations/) for popular API
gateways and service meshes and guides are available to start exploring quickly.

### Getting started

Gateway API is an official Kubernetes API like
[Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/).
Gateway API represents a superset of Ingress functionality, enabling more
advanced concepts. Similar to Ingress, there is no default implementation of
Gateway API built into Kubernetes. Instead, there are many different
[implementations](/overview/implementations/) available, providing significant choice in terms of underlying
technologies while providing a consistent and portable experience.

Take a look at the [API concepts documentation](/overview/concepts/api-overview/) and check out some of
the [Guides](/guides/) to start familiarizing yourself with the APIs and how they
work. When you're ready for a practical application open the [implementations
page](/overview/implementations/) and select an implementation that belongs to an existing technology
you may already be familiar with or the one your cluster provider uses as a
default (if applicable). Gateway API is a [Custom Resource Definition
(CRD)](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/) based API so you'll need to [install the CRDs](/guides/getting-started/#installing-gateway-api) onto a
cluster to use the API.

If you're specifically interested in helping to contribute to Gateway API, we
would love to have you! Please feel free to [open a new issue](https://github.com/kubernetes-sigs/gateway-api/issues/new/choose) on the
repository, or join in the [discussions](https://github.com/kubernetes-sigs/gateway-api/discussions). Also check out the [community
page](/contributing/) which includes links to the Slack channel and community meetings.

## Release highlights

### Graduation to beta

The `v0.5.0` release is particularly historic because it marks the growth in
maturity to a beta API version (`v1beta1`) release for some of the key APIs:

- [GatewayClass](/reference/api-types/gatewayclass/)
- [Gateway](/reference/api-types/gateway/)
- [HTTPRoute](/reference/api-types/httproute/)

This achievement was marked by the completion of several graduation criteria:

- API has been [widely implemented](/overview/implementations/).
- Conformance tests provide basic coverage for all resources and have multiple implementations passing tests.
- Most of the API surface is actively being used.
- Kubernetes SIG Network API reviewers have approved graduation to beta.

For more information on Gateway API versioning, refer to the [official
documentation](/overview/concepts/versioning/). To see
what's in store for future releases check out the [next steps](#next-steps)
section.

### Release channels

This release introduces the `experimental` and `standard` [release channels](/overview/concepts/versioning/#release-channels)
which enable a better balance of maintaining stability while still enabling
experimentation and iterative development.

The `standard` release channel includes:

- resources that have graduated to beta
- fields that have graduated to standard (no longer considered experimental)

The `experimental` release channel includes everything in the `standard` release
channel, plus:

- `alpha` API resources
- fields that are considered experimental and have not graduated to `standard` channel

Release channels are used internally to enable iterative development with
quick turnaround, and externally to indicate feature stability to implementors
and end-users.

For this release we've added the following experimental features:

- [Routes can attach to Gateways by specifying port numbers](/geps/gep-957/)
- [URL rewrites and path redirects](/geps/gep-726/)

### Other improvements

For an exhaustive list of changes included in the `v0.5.0` release, please see
the [v0.5.0 release notes](https://github.com/kubernetes-sigs/gateway-api/releases/tag/v0.5.0).

## Gateway API for service mesh: the GAMMA Initiative
Some service mesh projects have [already implemented support for the Gateway
API](/overview/implementations/). Significant overlap
between the Service Mesh Interface (SMI) APIs and the Gateway API has [inspired
discussion in the SMI
community](https://github.com/servicemeshinterface/smi-spec/issues/249) about
possible integration.

We are pleased to announce that the service mesh community, including
representatives from Cilium Service Mesh, Consul, Istio, Kuma, Linkerd, NGINX
Service Mesh and Open Service Mesh, is coming together to form the [GAMMA
Initiative](/overview/mesh/), a dedicated
workstream within the Gateway API subproject focused on Gateway API for Mesh
Management and Administration.

This group will deliver [enhancement
proposals](/contributing/enhancement-requests/) consisting
of resources, additions, and modifications to the Gateway API specification for
mesh and mesh-adjacent use-cases.

This work has begun with [an exploration of using Gateway API for
service-to-service
traffic](https://docs.google.com/document/d/1T_DtMQoq2tccLAtJTpo3c0ohjm25vRS35MsestSL9QU/edit#heading=h.jt37re3yi6k5)
and will continue with enhancement in areas such as authentication and
authorization policy.

## Next steps

As we continue to mature the API for production use cases, here are some of the highlights of what we'll be working on for the next Gateway API releases:

- [GRPCRoute](/geps/gep-1016/) for [gRPC](https://grpc.io/) traffic routing
- [Route delegation](https://github.com/kubernetes-sigs/gateway-api/pull/1085)
- Layer 4 API maturity: Graduating [TCPRoute](https://github.com/kubernetes-sigs/gateway-api/blob/main/apis/v1alpha2/tcproute_types.go), [UDPRoute](https://github.com/kubernetes-sigs/gateway-api/blob/main/apis/v1alpha2/udproute_types.go) and
  [TLSRoute](https://github.com/kubernetes-sigs/gateway-api/blob/main/apis/v1alpha2/tlsroute_types.go) to beta
- [GAMMA Initiative](/overview/mesh/) - Gateway API for Service Mesh

If there's something on this list you want to get involved in, or there's
something not on this list that you want to advocate be on the roadmap
please join us in the #sig-network-gateway-api channel on Kubernetes Slack or our weekly [community calls](/contributing/#meetings).
