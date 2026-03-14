---
title: "도구"
weight: 60
description: "Gateway API 리소스 작업에 사용할 수 있는 도구"
---

게이트웨이 API는 리소스와의 상호작용을 쉽게 하기 위한 다양한 도구를 제공한다.

## `ingress2gateway`

기존 인그레스 리소스가 게이트웨이 API에서 어떻게 보일지 궁금한가? `ingress2gateway`는 프로바이더별 인그레스 리소스를 게이트웨이 API 리소스로 변환할 수 있는 간편한 도구다. 이 도구는 게이트웨이 API SIG-Network 하위 프로젝트에서 관리한다.

[kubernetes-sigs/ingress2gateway: 설치 가이드](https://github.com/kubernetes-sigs/ingress2gateway?tab=readme-ov-file#installation)로 시작하기!

## `gwctl`

게이트웨이 API 리소스를 관리하기 위한 커맨드라인 도구다. 정책, xRoute 등을 탐색하고 리소스와 상호작용할 수 있다.

[kubernetes-sigs/gwctl: 설치 가이드](https://github.com/kubernetes-sigs/gwctl?tab=readme-ov-file#installation)로 시작하기!

## `Headlamp`

쿠버네티스용 UI로 게이트웨이 API를 기본 지원한다. 게이트웨이 API 리소스를 맵 뷰 또는 데이터 테이블에서 확인할 수 있다. 실시간으로 업데이트되며, 쿠버네티스 리소스 간 링크를 통해 리소스 관계를 쉽게 파악할 수 있다.

[kubernetes-sigs/headlamp: 설치 가이드](https://github.com/kubernetes-sigs/headlamp?tab=readme-ov-file#quickstart)로 시작하기!

## 서드파티 도구

### `policy-machinery`

게이트웨이 API 정책 및 정책 컨트롤러 구현을 위한 프레임워크

[Kuadrant/policy-machinery](https://github.com/Kuadrant/policy-machinery/tree/main)에서 시작하기
