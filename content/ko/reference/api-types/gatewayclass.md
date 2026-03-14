---
title: "GatewayClass"
weight: 20
description: "GatewayClass resource for infrastructure provider definitions"
---

{{< channel-version channel="standard" version="v0.5.0" >}}

`GatewayClass` 리소스는 GA(정식 출시)되었으며 `v0.5.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하라.
{{< /channel-version >}}

[GatewayClass(게이트웨이 클래스)][gatewayclass]는 인프라 제공자가 정의하는
클러스터 범위의 리소스이다. 이 리소스는 인스턴스화할 수 있는 Gateway의 클래스를
나타낸다.

> 참고: GatewayClass는 [`networking.IngressClass` 리소스][ingress-class-api]와
> 동일한 기능을 수행한다.

```yaml
kind: GatewayClass
metadata:
  name: cluster-gateway
spec:
  controllerName: "example.net/gateway-controller"
```

인프라 제공자가 사용자를 위해 하나 이상의 `GatewayClass`를 생성할 것으로
예상된다. 이를 통해 어떤 메커니즘(예: 컨트롤러)이 `Gateway`를 구현하는지를
사용자로부터 분리할 수 있다. 예를 들어, 인프라 제공자는 인터넷 대면
애플리케이션과 비공개 내부 애플리케이션을 정의하는 `Gateway`를 반영하기 위해
`internet`과 `private`이라는 두 개의 `GatewayClass`를 생성할 수 있다.

```yaml
kind: GatewayClass
metadata:
  name: internet
  ...
---
kind: GatewayClass
metadata:
  name: private
  ...
```

클래스의 사용자는 `internet`과 `private`이 *어떻게* 구현되는지 알 필요가 없다.
대신 사용자는 `Gateway`가 생성된 클래스의 결과적인 속성만 이해하면 된다.

### GatewayClass 파라미터

`Gateway` API의 제공자는 클래스 정의의 일부로 컨트롤러에 파라미터를
전달해야 할 수 있다. 이는 `GatewayClass.spec.parametersRef` 필드를 사용하여
수행된다:

```yaml
# 인터넷 대면 애플리케이션을 정의하는 Gateway를 위한 GatewayClass.
kind: GatewayClass
metadata:
  name: internet
spec:
  controllerName: "example.net/gateway-controller"
  parametersRef:
    group: example.net
    kind: Config
    name: internet-gateway-config
---
apiVersion: example.net/v1alpha1
kind: Config
metadata:
  name: internet-gateway-config
spec:
  ip-address-pool: internet-vips
  ...
```

`GatewayClass.spec.parametersRef`에 커스텀 리소스를 사용하는 것이 권장되지만,
구현에서 필요한 경우 ConfigMap을 사용할 수도 있다.

### GatewayClass 상태

`GatewayClass`는 구성된 파라미터가 유효한지 확인하기 위해 제공자에 의해
반드시 검증되어야 한다(MUST). 클래스의 유효성은 `GatewayClass.status`를 통해
사용자에게 전달된다:

```yaml
kind: GatewayClass
...
status:
  conditions:
  - type: Accepted
    status: False
    ...
```

새로운 `GatewayClass`는 `Accepted` 조건이 `False`로 설정된 상태에서 시작된다.
이 시점에서 컨트롤러는 구성을 아직 확인하지 못한 것이다. 컨트롤러가 구성을
처리하면 조건이 `True`로 설정된다:

```yaml
kind: GatewayClass
...
status:
  conditions:
  - type: Accepted
    status: True
    ...
```

`GatewayClass.spec`에 오류가 있으면, 조건은 비어 있지 않으며 오류에 대한
정보를 포함한다.

```yaml
kind: GatewayClass
...
status:
  conditions:
  - type: Accepted
    status: False
    Reason: BadFooBar
    Message: "foobar" is an FooBar.
```

### GatewayClass 컨트롤러 선택

`GatewayClass.spec.controller` 필드는 `GatewayClass`를 관리하는 컨트롤러
구현을 결정한다. 필드의 형식은 불투명하며 특정 컨트롤러에 한정된다.
주어진 controller 필드에 의해 선택되는 GatewayClass는 클러스터의 다양한
컨트롤러가 이 필드를 어떻게 해석하는지에 따라 달라진다.

컨트롤러 작성자/배포자는 충돌을 피하기 위해 자신의 관리 제어 하에 있는
도메인/경로 조합을 사용하여 선택을 고유하게 만드는 것이 권장된다(RECOMMENDED)
(예: `example.net`으로 시작하는 모든 `controller`를 관리하는 컨트롤러는
`example.net` 도메인의 소유자이다).

컨트롤러 버전 관리는 경로 부분에 컨트롤러의 버전을 인코딩하여 수행할 수 있다.
예시 체계는 다음과 같다(컨테이너 URI와 유사):

```text
example.net/gateway/v1   // 버전 1 사용
example.net/gateway/v2.1 // 버전 2.1 사용
example.net/gateway      // 기본 버전 사용
```

[gatewayclass]: ../../reference/spec.md#gatewayclass
[ingress-class-api]: https://kubernetes.io/docs/concepts/services-networking/ingress/#ingress-class
