---
title: "Backend Protocol"
linkTitle: "Backend Protocol Selection"
weight: 15
description: "Configuring backend protocol selection with Gateway API"
---

{{< channel-version channel="standard" version="v1.2.0" >}}
This concept has been part of the Standard Channel since `v1.2.0`.
For more information on release channels, refer to our
[versioning guide]({{< ref "/overview/concepts/versioning" >}}).
{{< /channel-version >}}

Not all implementations of Gateway API support automatic protocol selection. In some cases protocols are disabled without an explicit opt-in.

When a Route's backend references a Kubernetes Service, application developers can specify the protocol using `ServicePort` [`appProtocol`][appProtocol] field.

For example the following `store` Kubernetes Service is indicating the port `8080` supports HTTP/2 Prior Knowledge.


```yaml
apiVersion: v1
kind: Service
metadata:
  name: store
spec:
  selector:
    app: store
  ports:
  - protocol: TCP
    appProtocol: kubernetes.io/h2c
    port: 8080
    targetPort: 8080
```

Currently, Gateway API has conformance testing for:

- `kubernetes.io/h2c` - HTTP/2 Prior Knowledge
- `kubernetes.io/ws` - WebSocket over HTTP

[appProtocol]: https://kubernetes.io/overview/concepts/services-networking/service/#application-protocol
