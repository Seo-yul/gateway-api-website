---
title: "HTTP 트래픽 분할"
linkTitle: "HTTP traffic splitting"
weight: 4
description: "Splitting HTTP traffic between backends for canary and blue-green deployments"
---

# HTTP 트래픽 분할

[HTTPRoute 리소스]({{< ref "/reference/api-types/httproute" >}})를 사용하면 서로 다른 백엔드 간에
트래픽을 이동시키기 위한 가중치를 지정할 수 있다. 이는 롤아웃 중 트래픽 분할,
카나리 변경, 또는 긴급 상황에 유용하다.
HTTPRoute의 `spec.rules.backendRefs`는 라우트 규칙이 트래픽을 보낼 백엔드 목록을
수용한다. 이러한 백엔드의 상대적 가중치가 백엔드 간의 트래픽 분할을 정의한다.
다음 YAML 스니펫은 단일 라우트 규칙에 대해 두 개의 Service가 백엔드로
나열되는 방법을 보여준다. 이 라우트 규칙은 트래픽을 90%는 `foo-v1`로,
10%는 `foo-v2`로 분할한다.

![Traffic splitting](/images/simple-split.png)

```yaml
{{< include file="examples/standard/traffic-splitting/simple-split.yaml" >}}
```

`weight`는 (백분율이 아닌) 비례적 트래픽 분할을 나타내므로, 단일 라우트 규칙 내
모든 가중치의 합이 모든 백엔드의 분모가 된다. `weight`는 선택적 매개변수이며
지정하지 않으면 기본값은 1이다. 라우트 규칙에 단일 백엔드만 지정된 경우,
가중치가 지정되었는지 여부와 관계없이 암묵적으로 100%의 트래픽을 수신한다.

## 가이드

이 가이드에서는 Service의 두 가지 버전 배포를 보여준다. 트래픽 분할은
v1에서 v2로의 점진적인 트래픽 이동을 관리하는 데 사용된다.

이 예제에서는 다음 게이트웨이가 배포되어 있다고 가정한다:

```yaml
{{< include file="examples/standard/simple-gateway/gateway.yaml" >}}
```

## 카나리 트래픽 롤아웃

처음에는 `foo.example.com`에 대한 프로덕션 사용자 트래픽을 서비스하는
단일 버전의 Service만 존재할 수 있다. 다음 HTTPRoute에는 `foo-v1` 또는
`foo-v2`에 대한 `weight`가 지정되지 않았으므로, 각각의 라우트 규칙에 의해
매칭된 트래픽의 100%를 암묵적으로 수신한다. 카나리 라우트 규칙은
(`traffic=test` 헤더 매칭을 사용하여) 프로덕션 사용자 트래픽을 `foo-v2`로
분할하기 전에 합성 테스트 트래픽을 보내는 데 사용된다.
[라우팅 우선순위]({{< ref "/reference/spec#httprouterule" >}})는 매칭되는 호스트와
헤더(가장 구체적인 매치)를 가진 모든 트래픽이 `foo-v2`로 전송되도록 보장한다.

![Traffic splitting](/images/traffic-splitting-1.png)


```yaml
{{< include file="examples/standard/traffic-splitting/traffic-split-1.yaml" >}}
```

## 블루-그린 트래픽 롤아웃

내부 테스트를 통해 `foo-v2`로부터의 성공적인 응답이 검증된 후,
점진적이고 보다 현실적인 테스트를 위해 소량의 트래픽을 새 Service로
이동시키는 것이 바람직하다. 아래 HTTPRoute는 `foo-v2`를 가중치와 함께
백엔드로 추가한다. 가중치의 합이 100이므로 `foo-v1`은 90/100=90%의
트래픽을, `foo-v2`는 10/100=10%의 트래픽을 수신한다.

![Traffic splitting](/images/traffic-splitting-2.png)


```yaml
{{< include file="examples/standard/traffic-splitting/traffic-split-2.yaml" >}}
```

## 롤아웃 완료

마지막으로, 모든 신호가 긍정적이면 트래픽을 `foo-v2`로 완전히 이동시키고
롤아웃을 완료할 때이다. `foo-v1`의 가중치는 `0`으로 설정되어
트래픽을 수신하지 않도록 구성된다.

![Traffic splitting](/images/traffic-splitting-3.png)


```yaml
{{< include file="examples/standard/traffic-splitting/traffic-split-3.yaml" >}}
```

이 시점에서 100%의 트래픽이 `foo-v2`로 라우팅되며 롤아웃이 완료된다.
어떤 이유로든 `foo-v2`에서 오류가 발생하면, 가중치를 업데이트하여 트래픽을
`foo-v1`으로 빠르게 되돌릴 수 있다. 롤아웃이 최종적으로 확정되면
v1을 완전히 폐기할 수 있다.
