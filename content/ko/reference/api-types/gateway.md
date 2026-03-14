---
title: "Gateway"
weight: 10
description: "Gateway resource for load balancer configuration"
---

{{< channel-version channel="standard" version="v0.5.0" >}}

`Gateway` 리소스는 GA(정식 출시)되었으며 `v0.5.0` 부터 표준 채널의 일부이다.
릴리스 채널에 대한 자세한 정보는
[버전 관리 가이드]({{< ref "/overview/concepts/versioning" >}})를 참조하라.
{{< /channel-version >}}

**Gateway**(게이트웨이)는 인프라 구성의 생명주기와 1:1로 대응된다.
사용자가 `Gateway`를 생성하면, `GatewayClass` 컨트롤러에 의해 일부 로드 밸런싱
인프라가 프로비저닝되거나 구성된다(자세한 내용은 아래 참조). `Gateway`는
이 API에서 동작을 트리거하는 리소스이다. 이 API의 다른 리소스들은
Gateway가 생성되어 리소스들을 연결하기 전까지는 구성 조각에 불과하다.

`Gateway` 스펙은 다음을 정의한다:

*   `GatewayClassName`- 이 Gateway에서 사용하는 `GatewayClass` 객체의 이름을
    정의한다.
*   `Listeners`- 호스트네임, 포트, 프로토콜, 종료, TLS 설정 및 리스너에
    연결할 수 있는 라우트를 정의한다.
*   `Addresses`- 이 게이트웨이에 요청된 네트워크 주소를 정의한다.

Gateway 스펙에 지정된 원하는 구성을 달성할 수 없는 경우,
Gateway는 상태 조건에 의해 제공되는 세부 정보와 함께 오류 상태가 된다.

### 배포 모델

`GatewayClass`에 따라, `Gateway`의 생성은 다음 작업 중 하나를 수행할 수 있다:

* 클라우드 API를 사용하여 LB 인스턴스를 생성한다.
* 소프트웨어 LB의 새 인스턴스를 생성한다(이 클러스터 또는 다른 클러스터에서).
* 이미 인스턴스화된 LB에 새로운 라우트를 처리하기 위한 구성 스탠자를 추가한다.
* SDN을 프로그래밍하여 구성을 구현한다.
* 아직 생각하지 못한 다른 것...

API는 이러한 작업 중 어떤 것이 수행될지 지정하지 않는다.

### Gateway 상태

`GatewayStatus`는 `spec`에 표현된 원하는 상태에 대한 `Gateway`의 상태를
표시하는 데 사용된다. `GatewayStatus`는 다음으로 구성된다:

- `Addresses`- Gateway에 실제로 바인딩된 IP 주소를 나열한다.
- `Listeners`- `spec`에 정의된 각 고유 리스너에 대한 상태를 제공한다.
- `Conditions`- Gateway의 현재 상태 조건을 설명한다.

`Conditions`와 `Listeners.conditions`는 모두 쿠버네티스의 다른 곳에서
사용되는 조건 패턴을 따른다. 이는 조건의 타입, 조건의 상태, 그리고
이 조건이 마지막으로 변경된 시간을 포함하는 목록이다.
