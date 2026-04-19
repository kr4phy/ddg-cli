# DuckDuckGo 풀 SERP (`duckduckgo.com`) 기술 사양

기준 시각: 2026-04-19  
대상: `https://duckduckgo.com/?q=...` (풀 자바스크립트 SERP)

## 1) 렌더링 모델

- 풀 SERP는 자바스크립트 기반 렌더링이다.
- 초기 HTML에는 셸 컨테이너와 부트스트랩 스크립트가 포함된다.
- 초기 응답에 웹 결과 노드(`<div class="result ...">`)가 정적으로 포함되지 않는다.
- 런타임 스크립트가 `DDG.deep.initialize('/d.js?...', true)`로 딥 페치를 시작한다.

---

## 2) 초기 HTML 구조 (검증됨)

```html
<body class="body--serp">
  <input id="state_hidden" name="state_hidden" type="text" size="1">
  <div id="spacing_hidden_wrapper"><div id="spacing_hidden"></div></div>

  <div id="header_wrapper" data-testid="header" class="header-wrap ...">
    <div id="header" class="header cw">
      <div id="react-search-form"></div>
    </div>
    <div id="react-duckbar" data-testid="duckbar"></div>
  </div>

  <div id="zero_click_wrapper" class="zci-wrap">
    <div id="react-root-zci"></div>
  </div>

  <div id="vertical_wrapper" class="verticals"></div>

  <div id="web_content_wrapper" class="content-wrap" ...>
    <div data-testid="mainline" class="results--main">
      <noscript>...</noscript>
    </div>
    <div id="react-layout"></div>
  </div>
</body>
```

### 핵심 컨테이너

| 선택자 | 역할 |
|---|---|
| `#react-search-form` | 검색 폼 마운트 포인트 |
| `#react-duckbar` | 네비게이션/duckbar 마운트 포인트 |
| `#react-root-zci` | 즉시답변 마운트 포인트 |
| `#vertical_wrapper` | 버티컬 탭 마운트 포인트 |
| `#react-layout` | 메인 웹 결과 마운트 포인트 |

---

## 3) 초기 HTML 내 부트스트랩 변수 (검증됨)

인라인 스크립트에서 검색/세션 상태를 노출:

```javascript
rq = "python";
rqd = "python";
ct = "KR";
rl = "us-en";
rt = "D";
rds = 30;
rs = 0;
df = "";
vqd = "4-...";
perf_id = "...";
parent_perf_id = "...";
```

추가로 확인되는 객체:
- `backendExperimentAssignments`
- `_bootstrapBackendData`
- `window.__preloadData__`
- `window.__initialSearchFormData__`

---

## 4) 런타임 데이터 페치 부트스트랩 (초기 HTML에서 검증됨)

인라인 `nrji()` 스크립트에서:

```javascript
nrj('/t.js?...');
DDG.deep.initialize('/d.js?...&wrap=1&...', true);
DDG.ready(nrji, 1);
```

### 관찰된 `/d.js` 쿼리 필드

| 파라미터 | 의미(관찰 맥락) |
|---|---|
| `q` | 검색어 |
| `t` | 소스/타입 마커 |
| `l` | 로케일 코드 |
| `s` | 시작 오프셋 |
| `ct` | 국가 |
| `bing_market` | 백엔드 마켓 |
| `dp` | 내부 토큰 |
| `wpa`, `wpl` | 내부 컨텍스트 |
| `perf_id`, `parent_perf_id`, `perf_sampled` | 성능/세션 추적 |
| `host_region` | 호스트 리전 |
| `sp`, `dfrsp`, `wrap`, `aps` | 내부 응답/포맷 플래그 |

엔드포인트 상세 프로브는 별도 문서 참조:
- `djs-api-spec-ko.md`
- `tjs-api-spec-ko.md`

---

## 5) Noscript 폴백 (검증됨)

`data-testid="mainline"` 내부:

```html
<noscript>
  <meta http-equiv="refresh" content="0;URL=/html?q=python">
  <div class="msg msg--noscript">...</div>
</noscript>
```

자바스크립트 미사용 시 `/html`로 리다이렉트된다.

---

## 6) 봇 방어 훅 (검증됨)

문서 하단 스크립트:

```javascript
DDG.deep.anomalyDetectionBlock({...});
```

풀 SERP의 런타임 이상행위/봇 처리 경로에 해당한다.

---

## 7) 초기 HTML에 있는 것과 없는 것

### 있음
- 쿼리/상태 메타데이터
- React 마운트 컨테이너
- 부트스트랩 스크립트 및 런타임 엔드포인트(`/t.js`, `/d.js`)
- noscript 리다이렉트 경로

### 없음
- `#react-layout` 하위의 정적 웹 결과 카드 HTML
- 각 결과의 정적 title/snippet/url 노드

---

## 8) 추출 경계

### 초기 응답만으로 추출 가능
- `rq`, `vqd`, 로케일, perf id 등 상태값
- 런타임 엔드포인트 템플릿 및 쿼리 필드
- 컨테이너 ID/선택자와 페이지 셸 구조

### 자바스크립트 런타임 DOM 필요
- `#react-layout` 내 최종 렌더링된 결과 카드 노드
- 클라이언트 렌더러가 부여하는 런타임 전용 속성
- 렌더링 후 결과 순서/페이지네이션 UI 상태
