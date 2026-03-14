---
title: "참여 방법"
weight: 1
---

이 페이지에는 API 관련 모든 회의록, 설계 문서 및 관련 토론에 대한 링크가
포함되어 있다. 프로젝트 내에서 공식적인 역할을 맡고자 한다면,
[기여자 래더]({{< ref "/contributing/contributor-ladder" >}})를 참조하자.

## 피드백 및 질문

일반적인 피드백, 질문 또는 아이디어를 공유하려면 [새로운 토론을 생성][gh-disc]해 주자.

[gh-disc]:https://github.com/kubernetes-sigs/gateway-api/discussions/new

## 버그 리포트

버그 리포트는 이 저장소의 [GitHub Issues][gh-issues]에 등록해야 한다.

**참고**: Gateway API 사양 자체가 아닌 특정 구현체에 해당하는 버그를 보고하는 경우,
[구현체 페이지][implementations]에서 해당 구현체의 도움을 받을 수 있는 저장소
링크를 확인하자.

[gh-issues]: https://github.com/kubernetes-sigs/gateway-api/issues/new/choose
[implementations]:../implementations.md

## 커뮤니케이션

주요 토론 및 공지사항은 [SIG-NETWORK 메일링 리스트][signetg]를 통해 전달된다.

또한 일상적인 질문과 토론을 위한 k8s.io의
[Slack 채널 (sig-network-gateway-api)][slack]도 운영하고 있다.

[signetg]: https://groups.google.com/a/kubernetes.io/g/sig-network
[slack]: https://kubernetes.slack.com/archives/CR0H13KGA

## 회의

Gateway API 커뮤니티 회의는 격주로 진행된다.
- **Week A**: 월요일 태평양 시간 오후 3시 (23:00 UTC, [시간대 변환][3pm-pst-convert])
- **Week B**: 화요일 태평양 시간 오전 8시 (16:00 UTC, [시간대 변환][8am-pst-convert])

Gateway API의 주요 회의인 만큼, 주제는 다양할 수 있으며 인그레스(ingress)와
서비스 메시(service mesh) 사용 사례를 포함한 새로운 주제와 아이디어가
논의되는 곳이다. 회의는 [Gateway API 메인테이너][maintainers]가
진행하며, 자원봉사자가 회의록을 작성한다.

* [Zoom 링크](https://zoom.us/j/441530404) (비밀번호는 [회의록][meeting notes] 문서에 있음)
* [캘린더에 추가](https://calendar.google.com/calendar/u/0/r?cid=88fe1l3qfn2b6r11k8um5am76c@group.calendar.google.com)

[8am-pst-convert]: http://www.thetimezoneconverter.com/?t=08:00&tz=PT%20%28Pacific%20Time%29
[3pm-pst-convert]: http://www.thetimezoneconverter.com/?t=15:00&tz=PT%20%28Pacific%20Time%29
[maintainers]:https://github.com/kubernetes-sigs/gateway-api/blob/main/OWNERS_ALIASES#L12

캘린더에는 _모든_ SIG Network 회의가 포함되어 있다
(따라서 Gateway API 회의뿐만 아니라 다른 하위 그룹 회의도 포함된다).

<iframe
  src="https://calendar.google.com/calendar/embed?src=88fe1l3qfn2b6r11k8um5am76c%40group.calendar.google.com"
  style="border: 0" width="800" height="600" frameborder="0"
  scrolling="no">
</iframe>

### 회의록 및 녹화

회의 안건과 회의록은 [회의록][meeting notes] 문서에서 관리된다.
다가오는 회의에서 논의할 주제를 자유롭게 추가하자.

모든 회의는 녹화되어
[Gateway API 회의 YouTube 재생목록][gateway-api-yt-playlist]에 자동으로 업로드된다.

[meeting notes]: https://docs.google.com/document/d/1eg-YjOHaQ7UD28htdNxBR3zufebozXKyI28cl2E11tU/edit
[gateway-api-yt-playlist]: https://www.youtube.com/playlist?list=PL69nYSiGNLP1GgO7k02ipPGZUFpSzGaHH

#### 초기 회의

일부 초기 커뮤니티 회의는 [별도의 YouTube 재생목록][early-yt-playlist]에
업로드되었으며, 이후 [SIG Network YouTube 재생목록][sig-net-yt-playlist]으로 이전되었다.

서비스 메시 사용 사례에 초점을 맞춘 초기 [GAMMA][gamma] 회의의 회의록은
별도의 [회의록 문서][gamma-meeting-notes]에서 확인할 수 있다.

[early-yt-playlist]: https://www.youtube.com/playlist?list=PL7KjrPTDcs4Xe6SZj-51WvBfufKf-la1O
[sig-net-yt-playlist]: https://www.youtube.com/playlist?list=PL69nYSiGNLP2E8vmnqo5MwPOY25sDWIxb
[gamma]: {{< ref "/overview/mesh/gamma" >}}
[gamma-meeting-notes]: https://docs.google.com/document/d/1s5hQU0CB9ehjFukRmRHQ41f1FA8GX5_1Rv6nHW6NWAA/edit#

#### 초기 설계 논의

* [Kubecon 2019 San Diego: API evolution design discussion][kubecon-2019-na-design-discussion]
* [SIG-NETWORK: Ingress Evolution Sync][sig-net-2019-11-sync]
* [Kubecon 2019 Barcelona: SIG-NETWORK discussion (general topics, includes V2)][kubecon-2019-eu-discussion]

[kubecon-2019-na-design-discussion]: https://docs.google.com/document/d/1l_SsVPLMBZ7lm_T4u7ZDBceTTUY71-iEQUPWeOdTAxM/preview
[kubecon-2019-eu-discussion]: https://docs.google.com/document/d/1n8AaDiPXyZHTosm1dscWhzpbcZklP3vd11fA6L6ajlY/preview
[sig-net-2019-11-sync]: https://docs.google.com/document/d/1AqBaxNX0uS0fb_fSpVL9c8TmaSP7RYkWO8U_SdJH67k/preview

## 발표 및 강연

| 날짜           | 제목 |    |
|----------------|-------|----|
| Mar, 2024      | Kubecon 2024 Paris: Configuring Your Service Mesh with Gateway API | [영상][2024-kubecon-video-1]|
| Mar, 2024      | Kubecon 2024 Paris: Gateway API: Beyond GA | [영상][2024-kubecon-video-2]|
| Mar, 2024      | Kubecon 2024 Paris: Tutorial: Configuring Your Service Mesh with Gateway API  | [영상][2024-kubecon-video-3]|
| Oct, 2023      | Kubecon 2023 Chicago: Gateway API: The Most Collaborative API in Kubernetes History Is GA | [영상][2023-kubecon-video-3]|
| May, 2023      | Kubecon 2023 Amsterdam: Emissary-Ingress: Self-Service APIs and the Kubernetes Gateway API | [영상][2023-kubecon-video-1]|
| May, 2023      | Kubecon 2023 Amsterdam: Exiting Ingress 201: A Primer on Extension Mechanisms in Gateway API | [영상][2023-kubecon-video-2]|
| Oct, 2022      | Kubecon 2022 Detroit: One API To Rule Them All? What the Gateway API Means For Service Meshes | [영상][2022-kubecon-video-4]|
| Oct, 2022      | Kubecon 2022 Detroit: Exiting Ingress With the Gateway API | [영상][2022-kubecon-video-3]|
| Oct, 2022      | Kubecon 2022 Detroit: Flagger, Linkerd, And Gateway API: Oh My! | [영상][2022-kubecon-video-2]|
| May, 2022      | Kubecon 2022 Valencia: Gateway API: Beta to GA | [영상][2022-kubecon-video-1]|
| May, 2021      | Kubecon 2021 Virtual: Google Cloud - Multi-cluster, Blue-green Traffic Splitting with the Gateway API | [영상][2021-kubecon-video-2]|
| May, 2021      | Kubecon 2021 Virtual: Gateway API: A New Set of Kubernetes APIs for Advanced Traffic Routing | [영상][2021-kubecon-video-1]|
| November, 2019 | Kubecon 2019 San Diego: Evolving the Kubernetes Ingress APIs to GA and Beyond | [영상][2019-kubecon-na-video]|
| November, 2019 | Kubecon 2019 San Diego: SIG-NETWORK Service/Ingress Evolution Discussion | [슬라이드][2019-kubecon-na-community-slides] |
| May, 2019      | [Kubecon 2019 Barcelona: Ingress V2 and Multicluster Services][2019-kubecon-eu] | [슬라이드][2019-kubecon-eu-slides], [영상][2019-kubecon-eu-video]|
| March, 2018    | SIG-NETWORK: Ingress user survey | [데이터][survey-data], [슬라이드][survey-slides] |

[2024-kubecon-video-1]: https://www.youtube.com/watch?v=UMGRp0fGk3o
[2024-kubecon-video-2]: https://www.youtube.com/watch?v=LITg6TvctjM
[2024-kubecon-video-3]: https://www.youtube.com/watch?v=UMGRp0fGk3o
[2023-kubecon-video-3]: https://www.youtube.com/watch?v=V3Vu_FWb4l4
[2023-kubecon-video-1]: https://www.youtube.com/watch?v=piDYmZObh_M
[2023-kubecon-video-2]: https://www.youtube.com/watch?v=7P55G8GsYRs:
[2022-kubecon-video-4]: https://www.youtube.com/watch?v=vYGP5XdP2TA
[2022-kubecon-video-3]: https://www.youtube.com/watch?v=sTQv4QOC-TI
[2022-kubecon-video-2]: https://www.youtube.com/watch?v=9Ag45POgnKw
[2022-kubecon-video-1]: https://www.youtube.com/watch?v=YPiuicxC8UU
[2021-kubecon-video-2]: https://www.youtube.com/watch?v=vs8YrjdRJJU
[2021-kubecon-video-1]: https://www.youtube.com/watch?v=lCRuzWFJBO0
[2019-kubecon-na-video]: https://www.youtube.com/watch?v=cduG0FrjdJA
[2019-kubecon-eu]: https://kccnceu19.sched.com/event/MPb6/ingress-v2-and-multicluster-services-rohit-ramkumar-bowei-du-google
[2019-kubecon-eu-slides]: https://static.sched.com/hosted_files/kccnceu19/97/%5Bwith%20speaker%20notes%5D%20Kubecon%20EU%202019_%20Ingress%20V2%20%26%20Multi-Cluster%20Services.pdf
[2019-kubecon-eu-video]: https://www.youtube.com/watch?v=Ne9UJL6irXY&t=1s
[survey-data]: https://github.com/bowei/k8s-ingress-survey-2018
[survey-slides]: https://github.com/bowei/k8s-ingress-survey-2018/blob/master/survey.pdf
[2019-kubecon-na-community-slides]: https://docs.google.com/presentation/d/1s0scrQCCFLJMVjjGXGQHoV6_4OIZkaIGjwj4wpUUJ7M

## 행동 강령

쿠버네티스 커뮤니티에 대한 참여는
[쿠버네티스 행동 강령](https://github.com/kubernetes/community/blob/master/code-of-conduct.md)에 의해 관리된다.
