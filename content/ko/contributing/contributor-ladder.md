---
title: "기여자 래더"
weight: 50
---

쿠버네티스 커뮤니티 내에서는 개인이 프로젝트에서 공식적인 역할을 획득할 수 있는 방법을 정의하기 위해 기여자 래더(contributor ladder)라는 개념이 개발되었다. Gateway API 기여자 래더는 [더 넓은 쿠버네티스 커뮤니티에서 정의한 역할](https://github.com/kubernetes/community/blob/master/community-membership.md)을 대체로 따르지만, 이 커뮤니티에 고유한 몇 가지 측면이 있다.

## 목표

이 문서가 다음 목표를 향한 첫 단계를 제공하기를 바란다:

* Gateway API 커뮤니티의 장기적 건강 보장
* 새로운 기여자가 프로젝트에서 공식적인 역할과 책임을 향해 나아가도록 장려
* 리더십 역할을 향한 경로를 명확하게 정의
* 프로젝트 리더십 역할을 채울 우수한 후보를 확보하기 위한 강력한 리더십 파이프라인 개발


## 범위

다음 저장소가 이 문서의 적용 범위에 해당한다:

* [kubernetes-sigs/gateway-api](https://github.com/kubernetes-sigs/gateway-api)
* [kubernetes-sigs/ingress2gateway](https://github.com/kubernetes-sigs/ingress2gateway)
* [kubernetes-sigs/gwctl](https://github.com/kubernetes-sigs/gwctl)
* [kubernetes-sigs/blixt](https://github.com/kubernetes-sigs/blixt)

이러한 각 프로젝트 내에서 전체 프로젝트 또는 프로젝트의 하위 영역에 대한 승인자(approver)나 리뷰어(reviewer)가 될 수 있는 기회가 있다. 예를 들어, 문서, GEP, API 변경, 또는 적합성 테스트에만 집중하는 리뷰어나 승인자가 될 수 있다.

## 일반 지침

### 1. 모든 사람을 환영한다

모든 기여에 감사한다. 프로젝트에서 공식적인 역할이 없어도 Pull Request를 만들거나 리뷰하고, 이슈나 논의를 도울 수 있다. 프로젝트 내에서 공식적인 역할을 수락하는 것은 전적으로 선택 사항이다.

### 2. 이러한 역할은 지속적인 기여를 필요로 한다

위에서 정의된 역할 중 하나에 지원하는 것은 해당 역할에 합당한 수준으로 계속 기여할 의향이 있는 경우에만 해야 한다. 어떤 이유로든 위의 역할 중 하나를 계속 수행할 수 없는 경우, 사임하자. 활동 없이 프로젝트를 장기간 떠나 있는 멤버는 쿠버네티스 GitHub 조직에서 제거되며, 현재 상태를 다시 숙지한 후 조직 멤버십 프로세스를 다시 거쳐야 한다.

### 3. 합의 없이 병합하지 않는다

변경 사항이 논쟁의 여지가 있다고 판단되는 경우, PR을 병합하기 전에 다른 사람들의 추가 관점을 기다리자. PR을 병합할 수 있는 권한이 있다고 해서 병합해야 하는 것은 아니다. PR을 무기한 차단할 수는 없지만, 모든 사람이 자신의 관점을 제시할 기회를 가졌는지 확인해야 한다.

### 4. 논의를 시작한다

이러한 역할 중 하나를 향해 나아가는 데 관심이 있다면, Slack에서 Gateway API 메인테이너에게 연락하자.

## 기여자 래더

Gateway API 기여자 래더는 다음 단계로 구성된다:

1. Member
2. Reviewer
3. Approver
4. Maintainer

또한 이 래더에 깔끔하게 맞지 않는 GAMMA 전용 리더십 역할이 있다. 이러한 모든 역할은 아래에서 더 자세히 정의한다.

## Member, Reviewer, Approver

기여자 래더의 첫 번째 단계는 이미 [업스트림 쿠버네티스 커뮤니티에서 명확하게 정의](https://github.com/kubernetes/community/blob/master/community-membership.md#community-membership)되어 있다. Gateway API는 나머지 쿠버네티스 커뮤니티와 함께 이러한 지침을 따른다. Gateway API 내에서 리뷰어나 승인자가 될 수 있는 다양한 영역이 있으며, 여기에는 다음이 포함된다:

* Conformance
* Documentation
* GEPs

## 메인테이너 및 Mesh 리드

기여자 래더의 마지막 단계는 프로젝트 전체에 대한 큰 전반적인 리더십 역할을 나타낸다. 이러한 역할에 사용 가능한 자리는 제한적이다(일반적으로 각 역할에 3-4명이 이상적이다). 가능한 한, 이러한 역할에 다양한 회사가 대표되도록 노력한다.

### 메인테이너

Gateway API 메인테이너는 쿠버네티스 커뮤니티 내에서 [Subproject Owner](https://github.com/kubernetes/community/blob/master/community-membership.md#subproject-owner)로 알려져 있다. Gateway API 메인테이너가 되기 위해 가장 중요하게 기대하는 것은 다음과 같다:

* Gateway API에 대한 최소 6개월 이상의 장기적이고 지속적인 기여
* 프로젝트의 기술적 목표와 방향에 대한 깊은 이해
* 중요한 향상 제안의 성공적인 작성 및 주도
* 최소 3개월 이상 승인자 역할 수행
* 커뮤니티 미팅을 이끌 수 있는 능력

위에 설명된 모든 기대 사항 외에도, 메인테이너가 프로젝트의 기술적 방향과 목표를 설정하는 것을 기대한다. 이 역할은 프로젝트의 건강에 매우 중요하며, 메인테이너는 새로운 승인자와 리뷰어를 멘토링하고, 논의와 의사 결정을 위한 건전한 프로세스가 마련되어 있는지 확인해야 한다. 마지막으로, 메인테이너는 궁극적으로 API의 새 버전 릴리스에 대한 책임이 있다.

## Mesh 리드

**Mesh 리드**(Mesh Lead)의 개념은 업스트림 쿠버네티스 커뮤니티 래더에 완벽하게 대응하는 역할이 없다. 본질적으로 Subproject Owner이지만, Gateway API 내의 주요 이니셔티브인 GAMMA(Gateway API for Mesh Management and Administration, Gateway for Mesh라고도 함) 이니셔티브를 위한 역할이다.

Mesh 리드가 되기 위해 가장 중요하게 기대하는 것은 다음과 같다:

* **서비스 메시**(Service Mesh) 구현체에 대한 상당한 경험
* 프로젝트의 기술적 목표와 방향에 대한 깊은 이해
* GAMMA 이니셔티브에 대한 최소 6개월 이상의 장기적이고 지속적인 기여
* 커뮤니티 미팅을 이끌 수 있는 능력

위에 설명된 모든 기대 사항 외에도, Mesh 리드가 GAMMA 이니셔티브의 기술적 방향과 목표를 설정하는 것을 기대한다. 논의와 의사 결정을 위한 건전한 프로세스가 마련되어 있는지, 그리고 릴리스 목표와 마일스톤이 명확하게 정의되어 있는지 확인해야 한다.
