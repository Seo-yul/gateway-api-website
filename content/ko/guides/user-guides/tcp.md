---
title: "TCP 라우팅"
linkTitle: "TCP routing"
weight: 13
description: "Configuring TCP routing with TCPRoute"
---

{{< note >}}
The `TCPRoute` resource described below is currently only included in the
"Experimental" channel of Gateway API. For more information on release
channels, refer to our [versioning guide]({{< ref "/overview/concepts/versioning" >}}).
{{< /note >}}

Gateway API는 여러 프로토콜과 함께 작동하도록 설계되었으며, [TCPRoute][tcproute]는
[TCP][tcp] 트래픽을 관리할 수 있는 라우트 중 하나이다.

이 예제에서는 하나의 게이트웨이 리소스와 두 개의 TCPRoute 리소스가 다음 규칙에 따라
트래픽을 분배한다:

- 게이트웨이의 포트 8080에 대한 모든 TCP 스트림은 `my-foo-service` 쿠버네티스
  Service의 포트 6000으로 전달된다.
- 게이트웨이의 포트 8090에 대한 모든 TCP 스트림은 `my-bar-service` 쿠버네티스
  Service의 포트 6000으로 전달된다.

이 예제에서는 두 개의 `TCP` 리스너가 [게이트웨이][gateway]에 적용되어 두 개의
별도 백엔드 `TCPRoute`로 라우팅된다. `Gateway`의 `listeners`에 설정된
`protocol`이 `TCP`임에 유의한다:

```yaml
{{< include file="examples/experimental/basic-tcp.yaml" >}}
```

위 예제에서는 `parentRefs`의 `sectionName` 필드를 사용하여 두 개의 별도 백엔드
TCP [Service][svc]에 대한 트래픽을 분리한다:

```yaml
spec:
  parentRefs:
  - name: my-tcp-gateway
    sectionName: foo
```

이는 `Gateway`의 `listeners`에 있는 `name`과 직접 대응된다:

```yaml
  listeners:
  - name: foo
    protocol: TCP
    port: 8080
  - name: bar
    protocol: TCP
    port: 8090
```

이 방식으로 각 `TCPRoute`는 `Gateway`의 서로 다른 포트에 "연결"되므로,
`my-foo-service` 서비스는 클러스터 외부에서 포트 `8080`의 트래픽을 수신하고
`my-bar-service`는 포트 `8090`의 트래픽을 수신한다.

`parentRefs`의 `port` 필드를 사용하여 라우트를 게이트웨이 리스너에 바인딩함으로써
동일한 결과를 달성할 수도 있다:

```yaml
spec:
  parentRefs:
  - name: my-tcp-gateway
    port: 8080
```

연결에 `sectionName` 대신 `port` 필드를 사용하면 게이트웨이와 관련 라우트 간의
관계가 더 밀접하게 결합된다는 단점이 있다. 자세한 내용은
[게이트웨이에 연결하기][attaching]를 참조한다.

[tcproute]: {{< ref "/reference/spec#tcproute" >}}
[tcp]:https://datatracker.ietf.org/doc/html/rfc793
[httproute]: {{< ref "/reference/spec#httproute" >}}
[gateway]: {{< ref "/reference/spec#gateway" >}}
[svc]:https://kubernetes.io/docs/concepts/services-networking/service/
[attaching]: {{< ref "/reference/api-types/httproute#attaching-to-gateways" >}}
