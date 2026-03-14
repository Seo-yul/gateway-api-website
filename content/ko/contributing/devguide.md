---
title: "개발자 가이드"
weight: 10
---

## 프로젝트 관리

이 프로젝트의 할 일 목록을 관리하기 위해 GitHub 이슈와 프로젝트 대시보드를 활용하고 있다:

* [이슈 목록][gh-issues]
* [프로젝트 대시보드][gh-dashboard]

`good first issue`와 `help wanted`로 레이블된 이슈는 첫 번째 기여에 특히 좋다.

릴리스 진행 상황을 추적하기 위해 [마일스톤][gh-milestones]을 사용한다. 이러한 마일스톤은 일반적으로 해당하는 [유의적 버전(semver)][semver] 릴리스 버전 태그에 따라 레이블이 지정되며, 이는 일반적으로 해당 릴리스가 종료되고 릴리스가 완료될 때까지 순서대로 다음 릴리스에만 집중한다는 의미이다. Gateway API 메인테이너만이 마일스톤을 생성하고 이슈를 마일스톤에 연결할 수 있다.

이슈 해결의 시급성을 나타내거나, 이슈의 우선순위 지정을 위해 작성자 또는 커뮤니티의 추가 지원이 필요한지 여부를 나타내기 위해 [우선순위 레이블][prio-labels]을 사용한다. 이러한 레이블은 [PR 및 이슈 코멘트의 /priority 명령어][issue-cmds]로 설정할 수 있다. 예를 들어, `/priority important-soon`과 같이 사용한다.

[gh-issues]: https://github.com/kubernetes-sigs/gateway-api/issues
[gh-dashboard]: https://github.com/kubernetes-sigs/gateway-api/projects
[gh-milestones]: https://github.com/kubernetes-sigs/gateway-api/milestones
[semver]:https://semver.org/
[prio-labels]:https://github.com/kubernetes-sigs/gateway-api/labels?q=priority
[issue-cmds]:https://prow.k8s.io/command-help?repo=kubernetes-sigs%2Fgateway-api

## 사전 준비

Gateway API 개발을 시작하기 전에 다음 사전 준비 사항을 설치하는 것을 권장한다:

* [KinD](https://kubernetes.io/docs/tasks/tools/#kind): 독립형 로컬 쿠버네티스 **클러스터**(Cluster)이다. 최소한 하나의 컨테이너 런타임이 필요하다.
* [Docker](https://docs.docker.com/engine/install/): KinD를 실행하기 위한 사전 요구 사항이다. [Podman](https://podman.io/docs/installation)과 같은 대안을 선택할 수 있지만, 그렇게 하는 것은 본인의 책임이라는 점에 유의하자.
* [BuildX](https://github.com/docker/buildx): `make verify` 실행을 위한 사전 요구 사항이다.
* [Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl): 쿠버네티스 커맨드라인 도구이다.
* [Go](https://golang.org/doc/install): 이 프로젝트의 주요 프로그래밍 언어이다. 컴파일 오류가 발생하지 않도록 이 [파일](https://github.com/kubernetes-sigs/gateway-api/blob/main/go.mod#L3)에서 최소 `Go` 버전을 확인하자.
* [Digest::SHA](https://metacpan.org/pod/Digest::SHA): 필수 의존성이다. `perl-Digest-SHA` 패키지를 설치하여 얻을 수 있다.


## 개발: 빌드, 배포, 테스트 및 검증

저장소를 복제한다:

```
mkdir -p $GOPATH/src/sigs.k8s.io
cd $GOPATH/src/sigs.k8s.io
git clone https://github.com/kubernetes-sigs/gateway-api
cd gateway-api
```

이 프로젝트는 Go 모듈을 사용하므로 $GOPATH 외부에 환경을 설정할 수도 있다.


### 코드 빌드

이 프로젝트는 빌드를 구동하기 위해 `make`를 사용한다. `make`는 이전에 생성된 코드를 정리하고, 코드 생성기를 실행하며, 코드에 대한 정적 분석을 실행하고 쿠버네티스 CRD를 생성한다. 최상위 makefile에서 전체 빌드를 시작할 수 있다:

```shell
make generate
```


#### 실험적(Experimental) 필드 추가

API에 대한 모든 추가 사항은 실험적 릴리스 채널에서 시작해야 한다. 실험적 필드는 Go 타입 정의에서 `<gateway:experimental>` 어노테이션으로 표시되어야 한다. Gateway API CRD 생성 시 이러한 필드는 실험적 CRD 세트에만 포함된다.

실험적 필드가 제거되거나 이름이 변경되면, 원래 필드 이름은 go 구조체에서 제거되어야 하며, 필드 이름이 재사용되지 않도록 톰스톤(tombstone) 주석을 남겨야 한다.

예시:

```golang
// DeprecatedField is tombstoned to show why 16 is reserved protobuf tag.
// DeprecatedField string `json:"deprecatedField,omitempty" protobuf:"bytes,16,opt,name=deprecatedField"`
```

### 코드 배포

다음 명령어를 사용하여 기존 `Kind` 클러스터에 CRD를 배포한다.

```shell
make crd
```

다음 명령어를 사용하여 CRD가 배포되었는지 확인한다.

```shell
kubectl get crds
```

### 수동 테스트

[Gateway API 구현체]({{< ref "/overview/implementations" >}})를 설치하고 변경 사항을 테스트한다. 몇 가지 [예시](/guides/)를 살펴보자.

### 검증 {#verification}

변경 사항을 제출하기 전에 저장소에서 정적 분석을 실행하자. [Prow presubmit][prow-setup] 검증이 실패하면 변경 사항이 병합되지 않는다.

```shell
make verify
```

[prow-setup]: https://github.com/kubernetes/test-infra/tree/master/config/jobs/kubernetes-sigs/gateway-api


## 개발 후: Pull Request, 문서화, 추가 테스트
### Pull Request 제출

Gateway API는 [Kubernetes](https://github.com/kubernetes/community/blob/master/contributors/guide/pull-requests.md)와 유사한 Pull Request 프로세스를 따른다. Pull Request가 자동으로 병합되려면 다음 단계를 완료해야 한다.

- [CLA 서명](https://git.k8s.io/community/CLA.md) (사전 요구 사항)
- [Pull Request 열기](https://help.github.com/articles/about-pull-requests/)
- [검증](#verification) 테스트 통과
- 리뷰어와 코드 소유자의 모든 필요한 승인 획득


### 문서화

사이트 문서는 마크다운(Markdown)으로 작성되며 [mkdocs](https://www.mkdocs.org/)로 컴파일된다. 각 PR에는 자동으로 [Netlify](https://netlify.com/) 배포 미리보기가 포함된다. 새로운 코드가 병합되면, Netlify를 통해 자동으로 [gateway-api.sigs.k8s.io]()에 배포된다. 로컬에서 문서 변경 사항을 수동으로 미리보기하려면 mkdocs를 설치하고 다음을 실행한다:

```shell
 make docs
```

올바른 버전의 mkdocs를 쉽게 사용하기 위해 컨테이너에서 문서를 빌드하고 서빙할 수 있다:

```shell
$ make build-docs
...
INFO    -  Documentation built in 6.73 seconds
$ make live-docs
...
INFO    -  [15:16:59] Serving on http://0.0.0.0:3000/
```

그런 다음 http://localhost:3000/ 에서 문서를 확인할 수 있다.

문서 작성 방법에 대한 자세한 내용은 [문서 스타일 가이드]({{< ref "/contributing/style-guide" >}})를 참조하자.

### 적합성 테스트

적합성 테스트를 개발하거나 실행하려면 [적합성 테스트 문서]({{< ref "/overview/concepts/conformance#running-tests" >}})를 참조하자.

### 새로운 도구 추가
이 프로젝트를 빌드하고 관리하는 데 사용되는 도구는 `tools` 디렉토리에 자체적으로 포함되어 있다.

새로운 도구를 추가하려면 `go get -tool -modfile tools/go.mod the.tool.repo/toolname@version`을 사용하고, `go mod tidy -modfile=tools/go.mod`로 특정 모듈을 정리한다.

새로운 도구를 실행하려면 `go tool -modfile=tools/go.mod toolname`을 사용한다.
