---
title: "호스트네임"
weight: 100
description: "Gateway API가 라우트 연결, 트래픽 구분 및 라우팅에 호스트네임을 사용하는 방식"
---

## 소개/이 문서의 목적

이 문서는 Gateway API 사용자와 Gateway API 객체를 프로그래밍 방식으로 상호작용하는 시스템을 구축하는 통합 개발자 모두에게 Gateway API가 호스트네임을 어떻게 사용하는지, 그리고 이러한 사용에 대해 가장 중요하게 알아야 할 사항이 무엇인지를 더 잘 이해할 수 있도록 돕기 위한 것이다.

## 호스트네임은 어디서, 어떻게 구성할 수 있는가?

호스트네임은 라우트가 **호스트네임 교차**를 통해 게이트웨이 또는 리스너에 연결할 수 있는지 여부를 확인하는 데 사용되며, **라우팅 구별**을 통해 어떤 리스너와 라우트가 특정 요청을 수락해야 하는지를 선택하는 데에도 사용된다. **호스트네임 교차**와 **라우팅 구별** 모두 이 문서의 뒷부분에서 정의한다.

각 `hostname` 필드는 _정확한_ 호스트네임(예: `www.example.com`과 같은 호스트네임) 또는 _와일드카드_ 호스트네임(예: `*.example.com`과 같은 호스트네임)을 받아들일 수 있다.
정확한 호스트네임은 레이블 수에 따라 더 높거나 낮은 정밀도를 가질 수 있다 -
따라서 `www.example.com`은 `sub.domain.example.com`보다 덜 정밀하고, 와일드카드 `*.example.com`은 `www.example.com`보다 덜 정밀하다.

호스트네임의 정밀도 수준은 또한 특정 트래픽에 어떤 리스너가 매칭될지를 선택하는 과정에서 유효한 순서에 영향을 미치며, 더 정밀한 호스트네임이 덜 정밀한 호스트네임보다 우선한다.

{{< warning >}}
IP 주소는 Gateway API에서 유효한 호스트네임이 _결코_ 될 수 없다는 점에 유의하라. 다만, 이 문서 작성 시점에서 해당 필드의 유효성 검사가 IP 주소를 허용할 수 있다.
이는 버그이며 향후 수정될 예정이다. Gateway API는 이 동작에 의존하지 않기를 **강력히** 권장한다.
{{< /warning >}}

{{< note >}}
실제로 Gateway API에는 `Hostname`과 `PreciseHostname`이라는 두 가지 타입의 호스트네임이 있다. `Hostname`은 아래에 설명된 와일드카드 동작을 가지지만, `PreciseHostname`은 와일드카드를 허용하지 _않는다_. 그 외에는 두 타입이 동일하다.
{{< /note >}}

### 호스트네임 와일드카드

Gateway API에서 와일드카드는 호스트네임의 가장 왼쪽 문자로서_만_ 지원되며, 바로 뒤에 `.`이 와야 한다(이는 와일드카드가 [RFC-9499](https://www.rfc-editor.org/rfc/rfc9499.html), [RFC-2308](https://www.rfc-editor.org/rfc/rfc2308) 등의 DNS RFC에서 정의된 완전한 DNS 레이블만 매칭함을 의미한다).

예를 들면:

- `*.example.com`은 유효한 `hostname`이다
- `f*.example.com`은 와일드카드 문자 `*`가 가장 왼쪽 문자가 아니므로 유효한 `hostname`이 아니다.
- `*oo.example.com`도 와일드카드 문자 뒤에 `.`이 오지 않으므로 유효한 `hostname`이 아니다.

추가로 - 다른 많은 시스템과 달리 - 와일드카드는 단일 DNS 레이블이 아닌 _하나 이상의_ DNS 레이블과 매칭되는 것으로 정의된다. 예를 들면:

- `*.example.com`은 `www.example.com`과 `sub.domain.example.com`에 매칭되지만, `example.com`에는 매칭되지 않는다.
- `*.com`은 `example.com`에 매칭되며, `www.example.com`에도 매칭된다.

이는 호스트네임 교차와 라우팅 구별 모두에서 중요하다.

### 사용 가능한 `hostname` 필드

#### 리스너 (게이트웨이와 ListenerSet에서 사용 가능)

게이트웨이 또는 ListenerSet에서 리스너 스탠자에는 `hostname` 필드가 포함된다. 각 리스너는 최대 하나의 `hostname`을 가질 수 있으며, 가장 왼쪽 위치에 와일드카드를 포함할 수 있다.
`hostname` 필드가 지정되지 않으면, 호스트네임 교차와 라우팅 구별 모두에서 모든 호스트네임이 매칭된다.
이 경우, 호스트네임은 사실상 특수 값 `*`이다 - 이 값이 해당 필드에 _문자 그대로_ 유효하지는 않지만, 이 문서에서 "모든 것과 매칭" 동작을 표현하는 편리한 방법이다.

ListenerSet은 게이트웨이를 소유하지 않는 사용자가 게이트웨이에 추가 리스너를 도입할 수 있는 방법을 제공하는 비교적 새로운 리소스이다. 이를 위해 게이트웨이 객체와 같은 리스너 스탠자를 포함한다.

#### 라우트 (HTTPRoute, GRPCRoute, TLSRoute)

일부 라우트에는 `hostnames` 필드가 있다. 이 필드는 호스트네임 교차와 라우팅 구별 모두에서 `OR`로 처리된다. `hostnames`는 와일드카드 호스트네임을 포함할 수 있다.

HTTPRoute와 GRPCRoute의 경우, `hostname` 필드는 선택 사항이며, 제공되지 않으면 호스트네임 교차와 라우팅 구별 모두에서 모든 호스트네임이 라우트와 매칭된다.

TLSRoute의 경우, `hostname` 필드는 선택 사항이 아니다.

## 다양한 호스트네임은 실제로 어떤 역할을 하는가?

### 라우트 연결

**라우트 연결**은 라우트와 게이트웨이가 라우트가 게이트웨이에 연결될 수 있는지 여부에 동의하는 과정이다. 라우트는 게이트웨이 또는 ListenerSet일 수 있는 `parentRef`를 지정하고, 게이트웨이와 ListenerSet은 허용되는 라우트 종류 또는 라우트가 위치할 수 있는 네임스페이스를 선택할 수 있는 `allowedRoutes`를 지정할 수 있다.

이 문서에서 가장 중요한 점은, 리스너와 일부 유형의 라우트 모두 `hostname` 필드를 포함하며, 이 두 필드가 올바르게 **교차**해야 라우트가 **Accepted** 되고 게이트웨이에 연결된다는 것이다. 게이트웨이 ParentRef에 연결되지 않은 라우트는 해당 게이트웨이에 대해 아무 동작도 하지 않는다. 따라서 이를 올바르게 설정하는 것이 중요하다!

이 과정을 **호스트네임 교차**라고 하며, 라우트 유형이 `hostname` 필드를 포함하는 한 어떤 라우트 유형에 대해서도 동일하게 작동한다.

#### 호스트네임 교차

**호스트네임 교차**에서는 리스너와 라우트 모두의 `hostname` 필드가 고려되며, 해당 호스트네임이 겹치면 교차가 성공하고, 리스너는 다른 리스너 요구사항에 따라 라우트를 허용한다.

이 교차에는 다음과 같은 규칙이 있다(예시는 이후 표를 참조하라):

* 두 호스트네임 모두 **정확한** 경우(와일드카드를 포함하지 않는 경우), 호스트네임이 정확히 일치해야 한다.
* 리스너가 와일드카드 호스트네임을 가지고 라우트가 해당 와일드카드와 매칭되는 정확한 호스트네임을 가지면, 교차한다.
* 리스너가 정확한 호스트네임을 가지고 라우트가 해당 정확한 호스트네임과 매칭되는 와일드카드 호스트네임을 가지면, 교차한다.
* 리스너와 라우트 모두 와일드카드 호스트네임을 가지면, 겹치는 한 교차한다.
* 특수 와일드카드 `*`(다른 문자 없이)는 호스트네임 교차 목적으로 다른 모든 호스트네임과 매칭된다.
* 호스트네임이 설정되지 않은 것은 `hostname` 필드가 `*`로 설정된 것과 동일하다.

호스트네임이 교차하면, 리스너와 라우트 간의 연결이 진행될 수 있다(다른 모든 요구사항도 성공적인 경우).

실제로 교차하는 호스트네임을 **교차된** 호스트네임이라고 한다.
이는 아래에서 정의되는 트래픽 및 라우팅 구별에서 중요하다.

몇 가지 예시:

| 리스너 `hostname` | 라우트 `hostname` | 교차된 `hostname` | 이유 |
|---------------------|------------------|--------------------|------------|
| `www.example.com` | `www.example.com` | `www.example.com`  | 정확한 호스트네임이 교차한다 |
| `*.example.com` | `www.example.com` | `www.example.com` | 정확한 호스트네임이 동등한 와일드카드 호스트네임과 교차한다|
| `*.example.com` | `sub.domain.example.com` | `sub.domain.example.com` | 와일드카드 호스트네임이 교차 시 여러 DNS 레이블과 매칭될 수 있다 |
| `www.example.com` | `*.example.com` | `www.example.com` | 라우트의 와일드카드 호스트네임이 리스너의 정확한 호스트네임과 매칭된다 |
| `sub.domain.example.com` | `*.example.com` | `sub.domain.example.com` | 다중 레이블 와일드카드 매칭은 반대 방향에서도 작동한다 |
| `*.example.com` | `*.example.com` | `*.example.com` | 와일드카드 호스트네임이 정확히 일치할 때 매칭된다 |
| `*.com` | `*.example.com` | `*.example.com` | 덜 구체적인 와일드카드 호스트네임이 더 구체적인 와일드카드 호스트네임과 교차한다 |
| `*` | `www.example.com` | `www.example.com` | 모든 것과 매칭이 정확한 호스트네임과 교차한다 |
| `*` | `*` | `*` | 모든 것과 매칭이 다른 모든 것과 매칭과 교차한다 |


### 트래픽 및 라우팅 구별자

호스트네임은 _트래픽_ 및 _라우팅_ 구별에도 사용된다.
다시 말해, 리스너 집합에서 리스너를 선택하는 것(트래픽 구별)이든
리스너에 연결된 라우트 집합에서 라우트를 선택하는 것(라우팅 구별)이든,
트래픽이 라우팅될 위치를 선택하는 데 사용된다.

#### 리스너 순서

게이트웨이의 리스너(게이트웨이 자체에서 정의되거나 ListenerSet으로 연결된)가 동일한 `port`와 `protocol`을 가지지만 다른 `hostname`을 가지는 경우, 게이트웨이는 여러 리스너와 매칭될 _수 있는_ 트래픽을 _가장 구체적인_ 리스너로 보내야 한다.

이 과정에서는 **교차된 호스트네임**(호스트네임 교차 계산의 결과)만이 관련됨에 유의하라.

이는 `hostname`을 고려할 때 중요한데, 와일드카드가 구체성의 계층 구조를 생성하기 때문이다. 즉, 와일드카드를 포함하는 호스트네임을 가진 리스너는 정확한 호스트네임만 포함하는 리스너보다 _덜 정밀_하고 _덜 구체적_이다.

대략적으로 말하면, 호스트네임은 다른 것보다 와일드카드를 포함하지 않는 레이블이 _더 많으면_ 더 구체적이다.

몇 가지 예시:

* `www.example.com` (3개의 구체적 레이블)은 `*.example.com` (2개의 구체적 레이블)보다 더 구체적이다.
* `*.example.com` (2개의 구체적 레이블)은 `*.com` (1개의 구체적 레이블)보다 더 구체적이다.
* `*.com` (1개의 구체적 레이블)은 `*` (0개의 구체적 레이블)보다 더 구체적이다.

수락할 리스너를 선택할 때, 매칭할 정확한 호스트네임 세부사항은 프로토콜에 따라 다르지만, 모두 일반적인 패턴을 따른다:

* 정확한 매칭
* 요청 호스트네임에 대한 가장 구체적인 와일드카드 매칭
* 요청 호스트네임에 대한 일반 와일드카드 매칭(이는 `*` 호스트네임에 해당하는 "호스트네임 없음"의 특수 케이스를 포함한다).

#### SNI 매칭

TLS를 사용하는 `protocol` 값의 경우, 교차된 호스트네임은 다음과 같은 여러 세부사항과 매칭되어야 한다:

* `tls.mode`가 `Terminate`로 설정된 경우, 교차된 호스트네임이 TLS 종료에 사용되는 인증서의 CN 또는 SAN 필드에 존재해야 한다.
* `HTTPS` 또는 `TLS` 리스너에 도착하는 TLS 요청은 매칭되는 Server Name Indicator(SNI)를 가져야 한다.

Gateway API는 `tls.mode`가 `Terminate`로 설정된 리스너에서 연결에 사용되는 인증서에 교차된 호스트네임이 해당 필드에 존재하는지를 구현체가 검증할 것을 요구하지 _않는다_는 점에 유의하라. (구현체가 그렇게 _할 수는_ 있지만, 필수는 아니다).

SNI 매칭에서 "매칭"이란 SNI 호스트네임이 [RFC-2818](https://datatracker.ietf.org/doc/html/rfc2818#section-3.1)에서 제시된 서버 ID 매칭 규칙을 사용하여 교차된 호스트네임과 매칭되어야 함을 의미한다. 또한 SNI에는 와일드카드를 포함할 수 없으며, Gateway API 용어로 정확한 호스트네임이어야 함에 유의하라.

RFC-2818에서 인용:
> 인증서에 주어진 타입의 ID가 둘 이상 존재하는 경우(예: 둘 이상의 dNSName 이름), 해당 집합 중 하나라도 일치하면 수락할 수 있다. 이름에는 와일드카드 문자 *가 포함될 수 있으며, 이는 단일 도메인 이름 구성 요소 또는 구성 요소 조각과 매칭되는 것으로 간주된다. 예를 들어, `*.a.com`은 `foo.a.com`과 매칭되지만 `bar.foo.a.com`과는 매칭되지 않는다.

인증서에서 유효한 일부 값이 Gateway API의 `hostname`으로는 유효하지 않으므로, 일부 매칭은 불가능하다 - 예를 들어, `f*.com`은 유효한 `hostname`이 아니므로 RFC-2818에서처럼 `foo.com`과 매칭할 수 없다.

또한 RFC-2818에서는 와일드카드 문자 `*`가 여러 DNS 레이블이 아닌 _단일_ DNS 레이블만 매칭한다는 점에 유의하라. 따라서 TLS 연결이 실제로 _종료_될 때, SNI 동작이 호스트네임 교차 및 리스너 선택 동작과 미묘하게 다를 _수 있다_. 그러나 많은 Gateway API 구현체는 SNI를 접미사 매칭으로 처리하는 프록시 데이터 플레인을 가지고 있으며, 이 경우 매칭 동작은 동일하다.

요약하자면: **와일드카드 SNI 매칭에 주의하라.**

추가로, IP 주소는 Gateway API의 유효한 `hostname` 값이 결코 될 수 없으므로 매칭할 수 없다는 점을 기억하라.

RFC-2818 SNI 매칭 규칙에 따른 예시:

| 교차된 호스트네임 | 요청 SNI | 매칭 |
|---|---|---|
| `www.example.com` | `www.example.com` | ✅ |
| `www.example.com` | `foo.example.com` | ❌ |
| `*.example.com` | `www.example.com` | ✅ |
| `*.example.com` | `foo.example.com` | ✅ |
| `*.example.com` | `foo.bar.example.com` | ✅* |

SNI 매칭이 관련되는 경우는 다음과 같다:

* HTTPRoute와 리스너 `protocol` `HTTPS` 또는 `TLS` 및 `tls.mode` `Terminate`.
* GRPCRoute와 리스너 `protocol` `HTTPS` 또는 `TLS` 및 `tls.mode` `Terminate`.
* TLSRoute와 리스너 `protocol` `TLS` 및 `tls.mode` `Passthrough`.
* TLSRoute와 리스너 `protocol` `TLS` 및 `tls.mode` `Terminate`.


#### `Host` 헤더 매칭

암호화되지 않은 HTTP 연결 메타데이터를 사용하는 `protocol`과 라우트 조합(즉, HTTPRoute와 GRPCRoute)의 경우,
`Host` 또는 `:authority` 헤더가 교차된 호스트네임과 매칭되어야 한다. SNI 매칭과 마찬가지로, `Host` 헤더는 Gateway API 용어로 정확한 호스트네임이어야 하므로, 여기서의 매칭은 리스너 선택 매칭과 유사하다:

| 교차된 호스트네임 | `Host` 헤더 | 매칭 |
|---|---|---|
| `www.example.com` | `www.example.com` | ✅ |
| `www.example.com` | `foo.example.com` | ❌ |
| `*.example.com` | `www.example.com` | ✅ |
| `*.example.com` | `foo.example.com` | ✅ |
| `*.example.com` | `foo.bar.example.com` | ✅ |

`Host` 헤더 매칭에서 `*`는 SNI 매칭에서의 단일 DNS 레이블이 아닌, 하나 이상의 레이블과 매칭될 수 있다는 점에 유의하라.

### 예상 매칭 예시

| 리스너 호스트네임 | TLS 모드| 라우트 유형 | 라우트 호스트네임 | 연결됨? | 교차된 호스트네임 | SNI | SNI 매칭? | Host 헤더 | Host 헤더 매칭? | 비고 |
|---|---|---|---|---|---|---|---|---|---|---|
|`www.example.com` | None | HTTPRoute | `www.example.com` | ✅ | `www.example.com` | | | `www.example.com` | ✅ ||
|`*.example.com` | None | HTTPRoute | `www.example.com` | ✅ | `www.example.com` | | | `www.example.com` | ✅ ||
|`*.example.com` | None | HTTPRoute | `*.com` | ✅ | `*.example.com` | | | `www.example.com` | ✅ ||
|`*.example.com` | None | HTTPRoute | `*.com` | ✅ | `*.example.com` | | | `foo.bar.example.com` | ✅ | Host 헤더의 와일드카드 매칭은 하나 _이상의_ DNS 레이블과 매칭된다.|
|`*.example.com` | None | HTTPRoute | `www.example.com` | ✅ | `www.example.com` | | | `example.com` | ❌ ||
|`*.example.com` | None | HTTPRoute | `www.example.com` | ✅ | `www.example.com` | | | `foo.example.com` | ❌ ||
| `*.example.com` | Terminated | HTTPRoute | `www.example.com` | ✅ | `www.example.com` |`www.example.com` | ✅ | `www.example.com` | ✅ ||
| `*.example.com` | Terminated | HTTPRoute | `foo.bar.example.com` | ✅ | `foo.bar.example.com` |`foo.bar.example.com` | ✅ | `foo.bar.example.com` | ✅ | SSL 인증서는 리스너 호스트네임이 아닌 **교차된 호스트네임**과 매칭되어야 하며, 그렇지 않으면 SNI 매칭이 실패한다. 인증서의 `*.example.com`은 SNI로서 `foo.bar.example.com`과 매칭되지 _않기_ 때문이다. |
| `*.example.com` | Terminated | HTTPRoute | `*.example.com` | ✅ | `*.example.com` |`foo.bar.example.com` | ❌ | || 인증서의 `*.example.com`은 SNI로서 `foo.bar.example.com`과 매칭되지 _않는다_. |
| `*.example.com` | Terminated | HTTPRoute | `foo.example.com` | ✅ | `foo.example.com` |`foo.example.com` | ✅ | `foo.example.com` | ✅ ||
| `www.example.com` | Passthrough | TLSRoute | `www.example.com` | ✅ | `www.example.com` | `www.example.com`  | ✅ | |||
| `*.example.com` | Passthrough | TLSRoute | `www.example.com` | ✅ | `www.example.com` | `www.example.com`  | ✅ | |||
| `*.example.com` | Passthrough | TLSRoute | `www.example.com` | ✅ | `www.example.com` | `foo.example.com`  |  ❌  | |||
| `*.example.com` | Passthrough | TLSRoute | `foo.bar.example.com` | ✅ | `foo.bar.example.com` | `www.example.com`  | ❌ | ||SNI는 교차된 호스트네임과 매칭되어야 한다.|
| `*.example.com` | Passthrough | TLSRoute | unset | ✅ | `*.example.com` | `www.example.com`  | ✅ | |||
| `*.example.com` | Passthrough | TLSRoute | unset | ✅ | `*.example.com` | `foo.bar.example.com`  | ❌ | ||인증서의 와일드카드 이름에 대한 SNI 매칭은 단일 DNS 레이블만 매칭할 수 _있다_. (이는 인증서 이름이 교차된 호스트네임과 매칭된다는 가정하에이며, 이는 필수는 아니다.)|
| `www.example.com` | Terminated | TLSRoute | `www.example.com` | ✅ | `www.example.com` | `www.example.com`  | ✅ | ||Terminated 모드에서의 TLSRoute 예시는 Passthrough 모드에서의 TLSRoute 예시와 동일하므로 생략되었다.|


GRPCRoute는 `Host` 또는 `:authority` 헤더 매칭 측면에서 HTTPRoute와 동일하게 동작한다.


## `hostname` 필드의 프로그래밍 방식 사용

Gateway API에서의 호스트네임 사용에 대한 이 모든 세부사항은 호스트네임 필드를 프로그래밍 방식으로 사용하려는 Gateway API 통합(예: TLS용 인증서 프로비저닝 또는 게이트웨이용 DNS 레코드)에 몇 가지 영향을 미친다.

위의 모든 규칙의 조합으로 인해, 하나의 절대적인 불변 조건이 있다:

**호스트네임을 지원하는 모든 프로토콜의 트래픽은 반드시 _교차된 호스트네임_에 대해 수락될 수 있어야 한다(MUST).**

즉, 게이트웨이, ListenerSet 또는 라우트의 `hostname`을 개별적으로 사용하는 것은 올바르다고 보장되지 않는다. 통합이 그렇게 해서 _대부분의_ 경우에 맞을 수 있지만, 보장되지는 않는다. 정확성을 보장하려면, 통합은 호스트네임 교차 과정의 결과인 **교차된 호스트네임**과 그것이 통합의 호스트네임 사용 방식과 어떻게 상호작용하는지를 고려해야 한다.

해당 규칙의 부정적 추론은 다음과 같다:

**_교차된 호스트네임_을 결정할 수 없는 경우, 통합은 해당 리스너 및/또는 연결된 라우트를 무시해야 한다(MUST).**

교차된 호스트네임은 특정 리스너와 연결된 모든 라우트에 대해 필요한 호스트네임의 정규 표현이므로, 교차된 호스트네임을 결정할 수 없다면 해당 리스너에 대해 어떤 작업도 올바르게 수행할 수 없다. 통합이 올바른 작업을 수행할 것을 보장할 수 있는 충분한 정보가 없는 것이다.

이 규칙은 해당 리스너가 게이트웨이 _내부_에 있든 ListenerSet 객체에 있든 적용된다.

몇 가지 일반적인 예시와 권장 사항은 아래에 있다.

### 통합 개발자를 위한 일반 참고 사항

`hostname` 필드를 프로그래밍 방식으로 사용하는 컨트롤러 작성자에게는,
게이트웨이, ListenerSet 또는 HTTPRoute의 모든 사용은 전체 소유권 트리의 컨텍스트에서만 이루어져야 한다는 Gateway API의 일반 원칙을 기억하는 것이 중요하다.

즉, 게이트웨이 - 라우트 관계를 포함하는 호스트네임 교차 계산을 수행할 때,
컨트롤러는 반드시 전체 리소스 트리와만 상호작용해야 한다(MUST).

이는 컨트롤러가 다음을 수행해야 함을 의미한다:

* 감시할 하나 이상의 GatewayClass로 구성되어야 한다. ("모든 GatewayClass"도 괜찮지만, 서로 다른 GatewayClass가 고유한 호스트네임을 가질 필요는 _없으므로_ 주의하라.)
* 해당 GatewayClass에 속하는 모든 게이트웨이 중 `Accepted` 조건이 `status: true`인 것을 찾아야 한다(이는 구현체가 처리한다).
* 해당 게이트웨이를 가리키는 모든 ListenerSet 중 `Accepted` 조건이 `status: true`인 것을 찾아야 한다.
  (게이트웨이에 연결된 모든 리스너에 걸쳐 `hostname` 필드가 고유해야 함을 기억하라.
  이 리스너들이 게이트웨이에 있든 연결된 ListenerSet에 있든 마찬가지이다.
  구현체는 중복에 대해 ListenerSet을 `Accepted` `status: false`로 설정하여 이를 처리해야 한다.)
* 해당 게이트웨이 또는 ListenerSet을 가리키는 모든 라우트 중 `Accepted` 조건이 `status: true`인 것을 찾아야 한다.
* 각 게이트웨이-라우트 쌍 또는 각 ListenerSet-라우트 쌍에 대해 호스트네임 교차 계산을 수행해야 한다.
* 교차된 호스트네임을 기반으로 항목(DNS 레코드, 인증서 또는 기타)을 생성해야 한다.
  표준 충돌 해결 규칙도 준수해야 한다. 기본적으로:
  동일한 구성이 두 곳에 존재하는 경우, 생성 시간이 가장 오래된 것이 우선한다.


### DNS 레코드 자동 프로비저닝

이 경우의 주요 규칙은 간단하다:

**게이트웨이의 모든 리스너에 표현된 모든 _교차된 호스트네임_은 해당 게이트웨이의 `status.addresses`에 있는 모든 주소로 해석되어야 한다(MUST).**

{{< note >}}
통합은 사용자가 선택적으로 `foo.example.com`의 해석을 게이트웨이에 속하지 않는 임의의 주소로 재정의할 수 있는 수단을 제공할 수 있다(MAY).
예를 들어 게이트웨이에 알려지지 않은 외부 로드 밸런서 또는 리버스 프록시에 속하는 주소이다.
이 경우, `foo.example.com`으로의 트래픽이 결국 어떤 방법으로든 게이트웨이 주소에 도달하도록 하는 것은 사용자의 책임이다.
구현체가 이를 허용하는 경우, 이 DNS 레코드는 Gateway API 적합성 테스트를 통과하지 _않을_ 것이다.
이는 Gateway API 계약의 일부를 위반하기 때문이다
(즉, status.addresses에 나열된 주소가 게이트웨이에 연결하는 데 사용되어야 하는 실제 IP라는 것).
이 때문에 이 동작은 구현체 특정이며 통합 간에 이식 가능하지 않다.
{{< /note >}}

이 요구사항을 정확히 어떻게 충족하는지는 통합에 달려 있다.

몇 가지 구성 예시와 _예시_ 처리 방법은 다음과 같다:

* HTTPRoute `hostname`, 별도의 HTTPRoute에서: `foo.example.com`, `bar.example.com`, `baz.quux.example.com`.
* 리스너 `hostname`: `*.example.com`
* 게이트웨이 주소: `192.168.0.1`, `192.168.0.2`.

이 경우, 명령적 결과는 HTTPRoute의 모든 호스트네임에 대한 쿼리가 `192.168.0.1`, `192.168.0.2`, 또는 더 가능성 높게는 둘 사이를 전환하여 해석되는 것이다.

이는 다음과 같은 설정으로 달성할 수 있다:

* `foo.example.com`, `bar.example.com`, `baz.quux.example.com`에 대한 개별 A 레코드, 각각 `192.168.0.1`과 `192.168.0.2` 두 주소를 가리킴.
* `gateway-name.example.com`에 대한 A 레코드, `192.168.0.1`과 `192.168.0.2` 모두를 가리키고, `foo.example.com`, `bar.example.com`, `baz.quux.example.com`에 대한 CNAME
* 권한 있는 DNS 서버가 지원하는 경우 `*.example.com` 와일드카드 A 레코드, `192.168.0.1`과 `192.168.0.2` 모두를 가리킴. 이 경우, `foo.example.com`, `bar.example.com`, `baz.quux.example.com`이 _아닌_ 다른 호스트네임으로의 트래픽은 실제로 게이트웨이를 서비스하는 Gateway API 구현체에 의해 거부될 것으로 예상된다.
* 특정 호스트네임 요청이 올바른 주소로 해석되는 기타 모든 방법.

몇 가지 주의할 점:

* `quux.example.com`은 포함되지 _않으며_, 와일드카드 경우와 마찬가지로 기본 Gateway API 구현체에 의해 기본 주소로 해석되더라도 거부될 것이다. Gateway API는 가능한 한 이와 같은 중간 레코드를 생성하지 않아야 한다(SHOULD NOT)고 명시한다.

게이트웨이 주소가 IP 주소가 아닌 호스트네임인 경우의 다른 예시:

* HTTPRoute `hostname`: `foo.example.com`.
* 리스너 `hostname`: `*.example.com`
* 게이트웨이 주소: `some.long.cloud-lb.com`

가장 중요한 결과는 `foo.example.com`에 대한 요청이 `some.long.cloud-lb.com`이 해석되는 동일한 IP로 도달하는 것이다.

이는 CNAME 레코드로 할 수 있다 - 이는 A 레코드 관리를 `some.long.cloud-lb.com`의 제공자에게 맡기는 이점이 있다.

또는, 컨트롤러가 `some.long.cloud-lb.com`을 IP 주소로 해석하고 별도의 A 레코드를 생성할 _수도_ 있다.
또는 사용자가 해석을 재정의하도록 할 수 있다.
그러나 이는 또한 컨트롤러가 `some.long.cloud-lb.com`의 해석을 최신 상태로 유지하거나,
사용자 재정의를 최신 상태로 유지해야 함을 의미한다.

어느 경우든, 해당 컨트롤러는 여전히 Gateway API 사양에 적합하다. 관리에 관한 추가 주의사항은 명시되어야 한다.

### TLS 인증서 자동 프로비저닝

TLS 인증서의 자동 프로비저닝은 와일드카드 매칭 정의의 미묘한 차이 때문에 DNS 프로비저닝보다 약간 더 복잡하다. 이는 실제로 와일드카드 인증서 생성에만 영향을 미친다.

이에 대한 주요 규칙은 다음과 같다:

**리스너의 모든 교차된 호스트네임은 해당 리스너에서 사용되는 생성된 인증서에 표현되어야 한다.**

가장 간단한 경우, 이는 리스너에 속하는 모든 교차된 호스트네임이 해당 리스너에 연결될 생성된 인증서의 CN 또는 SAN 필드에 나열되어야 함을 의미한다.

간단한 예시:


* HTTPRoute `hostname`, 별도의 HTTPRoute에서: `foo.example.com`, `bar.example.com`, `baz.quux.example.com`.
* 리스너 `hostname`: `*.example.com`
* 게이트웨이 주소: `192.168.0.1`, `192.168.0.2`.

리스너에 존재하는 모든 생성된 인증서는 CN 또는 SAN 필드에 `foo.example.com`, `bar.example.com`, `baz.quux.example.com` 호스트네임이 표현되어야 한다(MUST).

DNS 프로비저닝 경우와 마찬가지로, `quux.example.com`은 표현되지 않으며 포함되어서는 안 된다(SHOULD NOT).

#### 와일드카드 인증서 처리

OWASP와 같은 표준의 일부 요약과 달리, 와일드카드 인증서는 합리적으로 안전한 방법으로 사용될 _수 있다_. 그러나 관리자 개입 없이 프로그래밍 방식으로 와일드카드 인증서를 생성하는 것은 거의 좋은 아이디어가 아니므로, Gateway API의 입장은 다음과 같다:

**통합은 `hostname` 필드를 사용하여 프로그래밍 방식으로 와일드카드 인증서를 생성해서는 안 된다(MUST NOT).**

다시 말해, 교차된 호스트네임에 와일드카드 문자가 포함되어 있으면, TLS 인증서 통합은 이를 무시해야 한다(MUST).

Gateway API는 API의 설계가 와일드카드 인증서를 어떻게 관리하도록 의도하는지에 대한 가이드를 작업 중이지만, 우리가 기대하는 일반적인 접근 방식은
와일드카드 인증서가 _상당히_ 큰 보안 영향을 미치므로, 이를 위한 전용 API를 사용하지 않고는 자동으로 관리되어서는 안 된다는 것이다.

따라서, cert-manager Certificate API는 괜찮다 - 와일드카드 인증서 관리를 위해 설계되었으며,
리스너의 교차 네임스페이스 TLS 참조 및 ReferenceGrant와 결합하면
실제 클러스터에서 비교적 안전하게 사용할 수 있다(거의 확실히 OWASP의 "와일드카드 인증서에 주의하라" 기준을 충족한다).
