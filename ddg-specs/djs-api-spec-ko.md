# DuckDuckGo `/d.js` 엔드포인트 기술 사양

기준 시각: 2026-04-19  
범위: `https://duckduckgo.com/d.js` 프로브 기반 사양

## 1) 엔드포인트 식별

- 호스트: `duckduckgo.com`
- 경로: `/d.js`
- 메서드: `GET`
- 풀 SERP 부트스트랩 호출 지점:
  - `DDG.deep.initialize('/d.js?...', true)`

부트스트랩 URL에서 관찰된 쿼리 키:

`q, t, l, s, ct, vqd, bing_market, p_ent, ex, dp, wpa, wpl, perf_id, parent_perf_id, perf_sampled, host_region, sp, dfrsp, wrap, aps`

---

## 2) 응답 외형 (관찰값)

- 상태 코드: `202` (프로브 응답)
- Content-Type: `application/x-javascript; charset=UTF-8`
- 페이로드 타입: 순수 JSON이 아닌 JavaScript 소스

주요 응답 계열:
1. `window.execDeep ... is506`
2. `window.execDeep ... isJsaChallenge`
3. `window.execDeep ... anomalyDetectionBlock`
4. `DDG.deep.initialize(...&jsa=..., false)`로 끝나는 JSA 스크립트 조각

---

## 3) 상태 전이 매트릭스 (헤더/쿠키 조건)

첫 요청은 페이지에서 추출한 전체 `/d.js?...` URL 사용.

| 케이스 | 1차 응답 | 1차 길이 | 2단계 동작 |
|---|---|---:|---|
| cookie + referer + UA | `isJsaChallenge` | 2125 | 계산된 `jsa` 후속 요청 => `anomalyDetectionBlock` (539) |
| cookie + UA | `isJsaChallenge` | 2340 | 계산된 `jsa` 후속 요청 => `anomalyDetectionBlock` (537) |
| referer + UA | `isJsaChallenge` | 2535 | 계산된 `jsa` 후속 요청 => `anomalyDetectionBlock` (541) |
| UA only | `isJsaChallenge` | 2559 | 계산된 `jsa` 후속 요청 => `anomalyDetectionBlock` (543) |
| cookie + referer, no UA | `anomalyDetectionBlock` | 538 | JSA 단계 없이 차단 응답 |

### 관찰된 전이 패턴

- 일반적으로 관찰된 흐름:
  - `isJsaChallenge` → 클라이언트 `jsa` 계산 → 후속 `/d.js?...&jsa_hash=...&jsa=...` → `anomalyDetectionBlock`
- 본 프로브에서는 UA 미포함 시 1단계부터 anomaly 응답으로 바뀌었다.

---

## 4) 파라미터 민감도 매트릭스

기준 컨텍스트: cookie + referer + UA.

| 변형 | 상태 | 길이 | 클래스 | 핵심 관찰 |
|---|---:|---:|---|---|
| baseline | 202 | 1898 | `isJsaChallenge` | 기본 챌린지 진입 |
| `vqd` 제거 | 202 | 2462 | `isJsaChallenge` | 챌린지 스크립트 유지 |
| `vqd` 변조 | 202 | 2456 | `isJsaChallenge` | 챌린지 스크립트 유지 |
| `dp` 제거 | 202 | 1377 | other/스크립트 조각 | 후속 URL이 비정상 접두(`'1&jsa_hash=...'`)로 생성됨 |
| `dp` 변조 | 202 | 2775 | `isJsaChallenge` | 챌린지 유지 |
| `s=30` | 202 | 2464 | `isJsaChallenge` | 오프셋 변경으로 우회되지 않음 |
| `s=90` | 202 | 2358 | `isJsaChallenge` | 오프셋 변경으로 우회되지 않음 |
| `wrap=0` | 202 | 2368 | other/스크립트 조각 | 샘플에서 `window.execDeep` 래퍼 없음 |
| `wrap` 제거 | 202 | 2561 | other/스크립트 조각 | 래퍼 없는 JSA 조각 관찰 |
| `host_region=use` | 202 | 2068 | `isJsaChallenge` | 챌린지 유지 |
| `host_region` 제거 | 202 | 2216 | `isJsaChallenge` | 챌린지 유지 |

---

## 5) 동일 `vqd` 페이지네이션 안정성

단일 `vqd`로 오프셋 테스트:

`s = 0, 10, 20, ... , 100`

관찰 결과:
- 전 구간 상태 코드 `202`
- 전 구간 클래스 `isJsaChallenge`
- 본문 길이는 달라지지만 응답 계열은 동일

동일 쿼리 페이지 3회 즉시 재로딩 시:
- 샘플 `vqd` 값은 3회 모두 동일했다.

---

## 6) 지역/언어 축 (`l`, `ct`, `kl`, `bing_market`)

### 6.1 풀 페이지 부트스트랩 매핑 (`?q=python&kl=...`)

| `kl` 입력 | `d.js l` | `d.js ct` | `d.js bing_market` | `d.js wpl` | 클래스 |
|---|---|---|---|---|---|
| `us-en` | `us-en` | `KR` | `en-US` | `en` | `isJsaChallenge` |
| `kr-kr` | `kr-kr` | `KR` | `ko-KR` | _(없음)_ | `isJsaChallenge` |
| `jp-jp` | `jp-jp` | `KR` | `ja-JP` | _(없음)_ | `isJsaChallenge` |
| `fr-fr` | `fr-fr` | `KR` | `fr-FR` | _(없음)_ | `isJsaChallenge` |
| `wt-wt` | `us-en` | `KR` | `en-US` | `en` | `isJsaChallenge` |

관찰 포인트:
- `kl` 값이 `l`, `bing_market`에 직접 반영된다.
- 본 환경에서는 `ct`가 `KR`로 고정 관찰되었다.
- 영어 매핑에서는 `wpl=en`이 보였고, 비영어 샘플에서는 `wpl`이 비어 있었다.

### 6.2 HTML SERP 구조 비교(`kl`별 샘플)

샘플 페이지(`us-en`, `kr-kr`, `jp-jp`, `fr-fr`, `wt-wt`)에서 DOM 구조는 동일:
- `div#links.results`
- `div.result.results_links.results_links_deep.web-result`
- `h2.result__title > a.result__a`
- `a.result__url`
- `a.result__snippet`
- `img.result__icon__img`

변화한 부분:
- 결과 순위/콘텐츠 언어/스니펫 내용
- URL 블록 근처 선택적 메타 텍스트(일부 로케일에서 관찰)

---

## 7) 프로브 범위 결론

본 프로브 구간에서 `/d.js`는 정적 결과 페이로드보다 챌린지 스크립트 엔드포인트로 동작했다:
- 스크립트 기반 분기(`is506`, `isJsaChallenge`, anomaly block)
- 파라미터/요청 컨텍스트에 따른 형태 변화
- 로케일 입력은 쿼리 매핑 필드를 변경하지만 챌린지 계열 자체는 유지됨

`/t.js` 시퀀스 및 동작은 별도 문서 참조:
- `tjs-api-spec-ko.md`
