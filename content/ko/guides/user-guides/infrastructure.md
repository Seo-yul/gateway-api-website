---
title: "게이트웨이 인프라 레이블 및 어노테이션"
linkTitle: "Infrastructure attributes"
weight: 16
description: "Propagating labels and annotations to generated infrastructure"
---

# 게이트웨이 인프라스트럭처 레이블 및 어노테이션

{{< channel-version channel="standard" version="v1.2.0" >}}

The `infrastructure` field is GA and has been part of the Standard Channel since
`v1.2.0`. For more information on release channels, refer to our [versioning
guide]({{< ref "/overview/concepts/versioning" >}}).
{{< /channel-version >}}

Gateway API 구현체는 각 **Gateway**(게이트웨이)가 작동하는 데 필요한 기반 인프라스트럭처를 생성하는 역할을 담당한다. 예를 들어, 쿠버네티스 클러스터에서 실행되는 구현체는 주로 [Service][service]와 [Deployment][deployment]를 생성하고, 클라우드 기반 구현체는 클라우드 로드 밸런서 리소스를 생성할 수 있다. 많은 경우, 이렇게 생성된 리소스에 레이블이나 어노테이션을 전파할 수 있으면 유용하다.


게이트웨이의 [`infrastructure` 필드][infrastructure]를 사용하면 Gateway API 컨트롤러가 생성하는 인프라스트럭처에 대한 레이블과 어노테이션을 지정할 수 있다.
예를 들어, 게이트웨이 인프라스트럭처가 클러스터 내에서 실행되는 경우, 다음과 같은 게이트웨이 구성을 사용하여 Linkerd와 Istio 주입을 모두 지정할 수 있으며, 설치된 **서비스 메시**(Service Mesh)에 인프라스트럭처를 더 쉽게 통합할 수 있다.

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: meshed-gateway
  namespace: incoming
spec:
  gatewayClassName: meshed-gateway-class
  listeners:
  - name: http-listener
    protocol: HTTP
    port: 80
  infrastructure:
    labels:
      istio-injection: enabled
    annotations:
      linkerd.io/inject: enabled
```

[infrastructure]: {{< ref "/reference/spec#gatewayinfrastructure" >}}
[service]: https://kubernetes.io/docs/concepts/services-networking/service/
[deployment]: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
