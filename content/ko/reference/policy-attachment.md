---
title: "메타리소스와 정책 연결"
weight: 30
description: "Metaresources and Policy Attachment in Gateway API"
---

Gateway API는 객체의 동작을 표준적인 방식으로 _보강_하는 Kubernetes 객체를
_메타리소스_로 정의한다. **ReferenceGrant**(레퍼런스 그랜트)는
이러한 일반적인 유형의 메타리소스의 예시이지만, 유일한 것은 아니다.

Gateway API는 또한 _**정책**(Policy) 연결_이라는 패턴을 정의하며, 이는
해당 객체의 spec 내에서 기술할 수 없는 추가 설정을 추가하기 위해
객체의 동작을 보강한다.

"정책 연결"은 하나의 객체에 영향을 미칠 수 있는("직접 정책 연결")
또는 계층 구조의 객체들에 영향을 미칠 수 있는("상속 정책 연결")
특정 유형의 _메타리소스_이다.

이 패턴은 EXPERIMENTAL이며, [GEP-713]({{< ref "/geps/gep-713" >}})에 기술되어 있다.
기술적 세부 사항은 해당 문서를 참조한다.
