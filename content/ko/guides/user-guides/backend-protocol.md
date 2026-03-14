---
title: "백엔드 프로토콜"
linkTitle: "Backend Protocol Selection"
weight: 15
description: "Configuring backend protocol selection with Gateway API"
---

# 백엔드 프로토콜

{{< channel-version channel="standard" version="v1.2.0" >}}

This concept has been part of the Standard Channel since `v1.2.0`.
For more information on release channels, refer to our
[versioning guide]({{< ref "/overview/concepts/versioning" >}}).
{{< /channel-version >}}

모든 Gateway API 구현체가 자동 프로토콜 선택을 지원하는 것은 아니다. 일부 경우에는 명시적으로 옵트인하지 않으면 프로토콜이 비활성화된다.

**Route**(라우트)의 백엔드가 쿠버네티스 Service를 참조하는 경우, 애플리케이션 개발자는 `ServicePort`의 [`appProtocol`][appProtocol] 필드를 사용하여 프로토콜을 지정할 수 있다.

예를 들어, 다음 `store` 쿠버네티스 Service는 포트 `8080`이 HTTP/2 Prior Knowledge를 지원함을 나타낸다.


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

현재 Gateway API는 다음에 대한 적합성 테스트를 제공한다:

- `kubernetes.io/h2c` - HTTP/2 Prior Knowledge
- `kubernetes.io/ws` - WebSocket over HTTP

[appProtocol]: https://kubernetes.io/docs/concepts/services-networking/service/#application-protocol
