# DuckDuckGo `/t.js` 엔드포인트 기술 사양

기준 시각: 2026-04-19  
범위: `https://duckduckgo.com/t.js` 프로브 기반 사양

## 1) 엔드포인트 식별

- 호스트: `duckduckgo.com`
- 경로: `/t.js`
- 메서드: `GET`
- 풀 SERP 부트스트랩 호출 지점:
  - `nrj('/t.js?...')`

부트스트랩 URL에서 관찰된 쿼리 키:

`q, t, l, s, ct, bing_market, p_ent, ex, dp, wpa, wpl, perf_id, parent_perf_id, perf_sampled, host_region, dfrsp, aps`

---

## 2) `/t.js` ↔ `/d.js` 시퀀스 및 역할 분리

동일 부트스트랩 함수(`nrji()`)에서 호출 순서:

1. `nrj('/t.js?...')`
2. `DDG.deep.initialize('/d.js?...', true)`

파싱한 함수 본문 내 호출 인덱스:
- `t.js` 호출 인덱스: `0`
- `d.js` 호출 인덱스: `416`
- `t.js` 선행 여부: `true`

파라미터 집합 관계:

- 두 엔드포인트 공통 키:
  - `q, t, l, s, ct, bing_market, p_ent, ex, dp, wpa, wpl, perf_id, parent_perf_id, perf_sampled, host_region, dfrsp, aps`
- `/d.js`에만 관찰된 키:
  - `vqd, sp, wrap`
- `/t.js`에만 관찰된 키:
  - _(샘플 부트스트랩 URL 기준 없음)_

본 프로브에서의 해석:
- `/t.js`가 먼저 호출되는 보조 엔드포인트.
- `/d.js`는 결과/딥 렌더 제어용 키(`vqd`, `sp`, `wrap`)를 추가로 가진다.

---

## 3) 응답 외형 (관찰값)

### 3.1 컨텍스트 변형

| 케이스 | 상태 | Content-Type | 본문 길이 |
|---|---:|---|---:|
| cookie + referer + UA | 200 | `application/x-javascript` | 0 |
| cookie + UA | 200 | `application/x-javascript` | 0 |
| UA only | 200 | `application/x-javascript` | 0 |

### 3.2 파라미터 민감도 프로브

| 변형 | 상태 | Content-Type | 본문 길이 |
|---|---:|---|---:|
| baseline | 200 | `application/x-javascript` | 0 |
| `q` 제거 | 200 | `application/x-javascript` | 0 |
| `dp` 제거 | 200 | `application/x-javascript` | 0 |
| `perf_id` 제거 | 200 | `application/x-javascript` | 0 |
| `host_region=use` | 200 | `application/x-javascript` | 0 |

본 프로브 구간 관찰:
- `/t.js`는 테스트한 컨텍스트/변형 모두에서 빈 본문 응답.
- `/t.js` 자체에서 챌린지 스크립트 계열은 관찰되지 않음.

---

## 4) 프로브 범위 결론

- `/t.js`는 풀 SERP 초기화에서 선행 호출되는 보조 엔드포인트로 관찰되었다.
- 챌린지 스크립트 기반 분기 동작은 `/d.js`에서 관찰되었다.
- 본 프로브에서는 `/t.js`가 메타데이터/텔레메트리 성격(안정적 `200`, 빈 본문)을 보였다.

딥 엔드포인트의 챌린지 전이는 아래 문서 참조:
- `djs-api-spec-ko.md`
