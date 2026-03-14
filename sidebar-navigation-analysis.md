# Gateway API Website — Sidebar Navigation 분석

## Context

Gateway API 웹사이트(Hugo + Docsy 테마)의 각 섹션별 좌측 Sidebar Navigation 구조를 분석·정리한다.
Sidebar는 디렉토리 계층 + front matter의 `weight` 값으로 자동 생성된다.

**설정 파일**: `hugo.toml` (line 104-116)
- `sidebar_menu_compact: false` — 항상 펼침 표시
- `sidebar_menu_foldable: true` — 섹션 접기/펼치기 가능
- `ul_show: 3` — 기본 3단계까지 표시

---

## 섹션별 Sidebar 구조

### 1. Overview (`/overview/`)

```
Overview (Introduction) ← _index.md (weight: 1)
├─ Introduction (weight: 1)
├─ Concepts/ (weight: 10)
│  ├─ API Overview (10)
│  ├─ Troubleshooting (20)
│  ├─ Conformance (30)
│  ├─ Roles and Personas (40)
│  ├─ Security (50)
│  ├─ Tooling (60)
│  ├─ Use Cases (70)
│  ├─ Versioning (80)
│  ├─ Traffic Matching (90)
│  └─ Hostnames (100)
├─ Mesh/ (weight: 20)
│  ├─ Mesh Overview (1)
│  ├─ GAMMA Initiative (10)
│  └─ Service Facets (20)
├─ Implementations/ (weight: 30)
│  ├─ Implementation List (1)
│  ├─ Implementation Wizard (10)
│  └─ Comparisons/ (20)
│     ├─ v1.4 (10) → v1.3 (20) → v1.2 (30) → v1.1 (40) → v1.0 (50)
├─ FAQ (weight: 60)
└─ Glossary (weight: 70)
```

### 2. Guides (`/guides/`)

```
Guides (weight: 2)
├─ Getting Started/ (weight: 1)
│  ├─ Simple Gateway (2)
│  ├─ Migrating from Ingress (3)
│  └─ Migrating from Ingress-NGINX (4)
├─ User Guides/ (weight: 2)
│  ├─ HTTP Routing (1)
│  ├─ HTTP Redirect/Rewrite (2)
│  ├─ HTTP Header Modifier (3)
│  ├─ Traffic Splitting (4)
│  ├─ HTTP Request Mirroring (5)
│  ├─ HTTP Query Param Matching (6)
│  ├─ HTTP Method Matching (7)
│  ├─ HTTP Timeouts (8)
│  ├─ HTTP CORS (9)
│  ├─ Multiple Namespace Routing (10)
│  ├─ TLS Configuration (11)
│  ├─ TLS Routing (12)
│  ├─ TCP Routing (13)
│  ├─ gRPC Routing (14)
│  ├─ Backend Protocol (15)
│  ├─ Infrastructure (16)
│  └─ ListenerSet (17)
├─ API Design (weight: 3)
├─ CRD Management (weight: 4)
└─ Implementers (weight: 5)
```

### 3. Reference (`/reference/`)

```
Reference (weight: 3)
├─ Policy Attachment (weight: 5)
├─ API Types/ (weight: 10)
│  ├─ BackendTLSPolicy (10)
│  ├─ BackendTrafficPolicy (20)
│  ├─ Gateway (30)
│  ├─ GatewayClass (40)
│  ├─ GRPCRoute (50)
│  ├─ HTTPRoute (60)
│  ├─ ListenerSet (70)
│  ├─ ReferenceGrant (80)
│  └─ TLSRoute (90)
└─ API Specification/ (weight: 20)
   ├─ Specification (10)
   ├─ Spec Extensions (20)
   ├─ v1.4/ (14) — Spec (10), Spec Extensions (20)
   └─ v1.5/ (15) — Spec (10), Spec Extensions (20)
```

### 4. Enhancements/GEPs (`/geps/`)

```
GEPs (Gateway Enhancement Proposals) (weight: 4)
├─ Overview (_index.md) — 탭 기반 UI로 상태별 분류
│
├─ Standard Channel (~29개): GEP-91, GEP-709, GEP-718, ...
├─ Memorandum (~6개): GEP-713, GEP-917, GEP-922, ...
├─ Experimental (~5개): GEP-1494, GEP-1619, ...
├─ Implementable (~3개): GEP-3779, GEP-3793, GEP-3949
└─ Provisional (~4개): GEP-1651, GEP-2627, ...

※ 각 GEP은 /geps/gep-XXXX/index.md 형태
※ Sidebar에는 전체 GEP이 플랫하게 나열됨
```

### 5. Contributing (`/contributing/`)

```
Contributing (weight: 5)
├─ Contributing Overview (_index.md)
├─ Developer Guide (weight: 10)
├─ Style Guide (weight: 20)
├─ Enhancement Requests (weight: 30)
├─ Release Cycle (weight: 40)
└─ Contributor Ladder (weight: 50)
```

### 6. Blog (`/blog/`)

```
Blog (weight: 80)
├─ Introducing v1alpha2 (2021)
├─ Graduating to Beta (2022)
└─ Mesh Support (2023)
```

---

## Sidebar 생성 메커니즘 요약

| 항목 | 설명 |
|------|------|
| 생성 방식 | 디렉토리 계층 + `weight` front matter로 자동 생성 |
| 정렬 기준 | `weight` 값 오름차순 (낮을수록 상단) |
| 섹션 구분 | 디렉토리 = foldable 섹션, `.md` 파일 = leaf 항목 |
| 최대 깊이 | 3단계 (`ul_show: 3`) |
| 접기/펼치기 | `sidebar_menu_foldable: true` |
| 다국어 | `/content/en/`, `/content/ko/` 동일 구조 |
| toc_root | 각 메인 섹션의 `_index.md`에 `toc_root: true` 설정 |

---

## 한국어(ko) 번역 현황

한국어 콘텐츠(`/content/ko/`)는 영문(`/content/en/`)과 동일한 디렉토리 구조 및 weight 값을 유지한다.

### 번역된 타이틀 매핑 (Overview 섹션)

| English | Korean |
|---------|--------|
| Introduction | 소개 |
| Concepts | 개념 |
| API Overview | API 개요 |
| Troubleshooting | 문제 해결 및 상태 |
| Conformance | 적합성 |
| Roles and Personas | 역할 및 페르소나 |
| Security | 보안 |
| Tooling | 도구 |
| Use Cases | 사용 사례 |
| Versioning | 버전 관리 |
| Traffic Matching | 트래픽 매칭 |
| Hostnames | 호스트네임 |
| Service Mesh | 서비스 메시 |
| GAMMA Initiative | GAMMA 이니셔티브 |
| Service Facets | 서비스 패싯 |
| Implementations | 구현체 |
| List | 목록 |
| Matching Wizard | 매칭 위자드 |
| Comparisons | 비교 |
| FAQ | FAQ |
| Glossary | 용어집 |

### 미번역 섹션

Guides, Reference, GEPs, Contributing 섹션은 타이틀이 영문 그대로 유지되어 있다.
