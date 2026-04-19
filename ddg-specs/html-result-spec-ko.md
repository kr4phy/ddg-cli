# DuckDuckGo HTML 결과 페이지 구조 분석

기준 시각: 2026-04-19  
대상: `https://html.duckduckgo.com/html/` 검색 결과

## 1) 전체 페이지 구조

### 기본 레이아웃

```
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" ...>
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <!-- 메타데이터 및 스타일시트 -->
  </head>
  <body class="body--html">
    <a name="top" id="top"></a>
    
    <!-- 숨김 상태 폼 -->
    <form action="/html/" method="post">
      <input type="text" name="state_hidden" id="state_hidden" />
    </form>

    <div>
      <div class="site-wrapper-border"></div>

      <!-- 헤더 및 검색 폼 -->
      <div id="header" class="header cw header--html">
        <!-- 로고 및 검색 폼 -->
        <form name="x" class="header__form" action="/html/" method="post">
          <div class="search search--header">
            <input name="q" type="text" value="..." />
            <input name="b" type="submit" />
          </div>
          <div class="frm__select">
            <select name="kl"><!-- 지역 --></select>
          </div>
          <div class="frm__select frm__select--last">
            <select name="df"><!-- 시간 필터 --></select>
          </div>
        </form>
      </div>

      <!-- 결과 컨테이너 -->
      <div>
        <div class="serp__results">
          <div id="links" class="results">
            <!-- 개별 결과 -->
          </div>
        </div>
      </div>
    </div>
  </body>
</html>
```

---

## 2) 검색 폼 구조

### 헤더 폼 컨테이너
- **선택자**: `form.header__form[action="/html/"][method="post"]`
- **목적**: 검색어, 지역, 시간 필터를 포함한 주 검색 폼

### 폼 요소

| 요소 | 선택자 | 타입 | 속성 | 설명 |
|------|--------|------|------|------|
| 검색어 입력 | `input[name="q"][type="text"]` | text | `class="search__input"`, `id="search_form_input_homepage"`, `autocomplete="off"` | 현재 검색어가 value에 유지됨 |
| 제출 버튼 | `input[name="b"][type="submit"]` | submit | `class="search__button search__button--html"`, `id="search_button_homepage"` | value 속성이 비어있을 수 있음 |
| 지역 선택 | `select[name="kl"]` | select | `class="frm__select"` (부모) | 지역 코드 드롭다운 |
| 시간 필터 | `select[name="df"]` | select | `class="frm__select frm__select--last"` (부모) | 옵션: "" (모든 시간), d, w, m, y |

### 폼 래퍼 클래스
- `.search` - 검색 입력 래퍼
- `.search--header` - 헤더 전용 검색
- `.frm__select` - 필터 선택 컨테이너 (지역)
- `.frm__select--last` - 마지막 필터 선택 컨테이너 (시간)

---

## 3) 결과 컨테이너 구조

### 주 결과 래퍼
- **선택자**: `div.serp__results > div#links.results`
- **구조**: 클래스 `results`, id `links`를 가진 컨테이너 div
- **콘텐츠**: 개별 결과 `<div>` 요소들의 시퀀스

---

## 4) 단일 결과 블록 구조

### 컨테이너
- **선택자**: `div.result.results_links.results_links_deep.web-result`
- **클래스**:
  - `result` - 기본 결과 스타일
  - `results_links` - 링크 기반 결과
  - `results_links_deep` - 깊은 링크 결과
  - `web-result` - 웹 검색 결과

### 결과 본문
- **선택자**: `div.links_main.links_deep.result__body`
- **목적**: 모든 결과 콘텐츠 컨테이너
- **주의**: 주석 "This is the visible part" 포함

### 단일 결과 블록 HTML 패턴

```html
<div class="result results_links results_links_deep web-result">
  <div class="links_main links_deep result__body">
    
    <!-- 제목/링크 -->
    <h2 class="result__title">
      <a rel="nofollow" class="result__a" href="//duckduckgo.com/l/?uddg=...&rut=...">
        Welcome to Python.org
      </a>
    </h2>

    <!-- URL 및 파비콘 섹션 -->
    <div class="result__extras">
      <div class="result__extras__url">
        <span class="result__icon">
          <a rel="nofollow" href="//duckduckgo.com/l/?uddg=...&rut=...">
            <img class="result__icon__img" width="16" height="16" 
                 src="//external-content.duckduckgo.com/ip3/www.python.org.ico" />
          </a>
        </span>
        <a class="result__url" href="//duckduckgo.com/l/?uddg=...&rut=...">
          www.python.org
        </a>
      </div>
    </div>

    <!-- 스니펫/설명 -->
    <a class="result__snippet" href="//duckduckgo.com/l/?uddg=...&rut=...">
      Python knows the usual control flow statements...
    </a>

    <div class="clear"></div>
  </div>
</div>
```

---

## 5) 결과 요소 상세

### 결과 제목
- **선택자**: `h2.result__title > a.result__a`
- **속성**:
  - `rel="nofollow"`: 항상 존재
  - `href`: DDG 리다이렉트 URL (형식: `//duckduckgo.com/l/?uddg=<encoded>&rut=<hash>`)
  - `class="result__a"`: 제목 링크 스타일
- **콘텐츠**: 링크 텍스트 (결과 제목)

### 결과 링크 (리다이렉트 URL)
- **형식**: `//duckduckgo.com/l/?uddg=<URL_ENCODED>&rut=<TRACKING_HASH>`
- **uddg 파라미터**: 인코딩된 대상 URL 포함
- **추출**: `uddg` 파라미터 값을 URL-디코딩

### 파비콘/결과 아이콘
- **선택자**: `span.result__icon > img.result__icon__img`
- **속성**:
  - `src`: `//external-content.duckduckgo.com/ip3/<domain>.ico`
  - `width="16"`, `height="16"`
  - `alt=""`: 빈 alt 속성
- **목적**: 사이트 파비콘/아이콘
- **형식**: 도메인에서 추출 (예: `www.python.org` → `.../ip3/www.python.org.ico`)

### 표시 URL
- **선택자**: `a.result__url`
- **속성**:
  - `href`: 제목과 동일한 리다이렉트 URL
- **콘텐츠**: 표시 URL (예: `www.python.org` 또는 `www.python.org/path`)
- **목적**: 사용자 참고용 도메인/경로 표시

### 결과 스니펫/설명
- **선택자**: `a.result__snippet`
- **속성**:
  - `href`: 제목/도메인과 동일한 리다이렉트 URL
- **콘텐츠**: 쿼리 키워드를 감싼 `<b>` 태그가 포함된 텍스트
- **형식**: 설명 텍스트, HTML 엔티티 인코딩된 특수 문자
- **주의**: 클릭 가능한 링크 요소 (단순 텍스트 span 아님)

---

## 6) CSS 클래스 참고

| 클래스 | 요소 | 목적 |
|--------|------|------|
| `body--html` | `<body>` | HTML 버전 본문 스타일 |
| `header` | `<div>` | 헤더 컨테이너 |
| `header--html` | `<div>` | HTML 특화 헤더 스타일 |
| `cw` | `<div>` | 컨테이너 너비/레이아웃 클래스 |
| `header__logo-wrap` | `<a>` | 로고 링크 래퍼 |
| `header__form` | `<form>` | 헤더 검색 폼 |
| `search` | `<div>` | 검색 입력 래퍼 |
| `search--header` | `<div>` | 헤더 검색 스타일 |
| `search__input` | `<input>` | 검색 입력 필드 |
| `search__button` | `<input>` | 제출 버튼 |
| `search__button--html` | `<input>` | HTML 버전 제출 버튼 |
| `frm__select` | `<div>` | 선택 컨테이너 (지역 필터) |
| `frm__select--last` | `<div>` | 마지막 선택 컨테이너 (시간 필터) |
| `serp__results` | `<div>` | 결과 페이지 컨테이너 |
| `result` | `<div>` | 단일 결과 래퍼 |
| `results_links` | `<div>` | 링크 기반 결과 |
| `results_links_deep` | `<div>` | 깊은 링크 결과 |
| `web-result` | `<div>` | 웹 검색 결과 타입 |
| `links_main` | `<div>` | 주 결과 콘텐츠 |
| `links_deep` | `<div>` | 깊은 링크 콘텐츠 |
| `result__body` | `<div>` | 결과 콘텐츠 본문 |
| `result__title` | `<h2>` | 결과 제목 제목 |
| `result__a` | `<a>` | 제목 링크 스타일 |
| `result__extras` | `<div>` | 추가 정보 컨테이너 |
| `result__extras__url` | `<div>` | URL 추가 섹션 |
| `result__icon` | `<span>` | 파비콘 컨테이너 |
| `result__icon__img` | `<img>` | 파비콘 이미지 |
| `result__url` | `<a>` | 표시 URL 링크 |
| `result__snippet` | `<a>` | 설명/스니펫 텍스트 링크 |
| `clear` | `<div>` | CSS float 해제 |

---

## 7) Lite 버전과의 주요 차이점

| 특징 | Lite (`lite.duckduckgo.com`) | HTML (`html.duckduckgo.com`) |
|------|---|---|
| 페이지 구조 | 테이블 기반 레이아웃 | div 기반 의미론적 클래스 |
| 결과 컨테이너 | `<table border="0">` | `<div class="serp__results">` |
| 결과 행 | `<tr>` 행 시리즈 | 단일 `<div class="result">` |
| 제목 요소 | `<a class="result-link">` | `<h2><a class="result__a">` |
| URL 표시 | `<span class="link-text">` | `<a class="result__url">` |
| 스니펫 | `<td class="result-snippet">` | `<a class="result__snippet">` |
| 파비콘 | 없음 | `<img class="result__icon__img">` |
| CSS 접근 | 최소 CSS | BEM 네이밍 현대식 CSS |
| DOCTYPE | HTML 4.01 Transitional | XHTML 1.0 Transitional |

---

## 8) 스크래퍼용 추출 패턴

### 패턴: 모든 결과 추출

```
1. 모든 <div class="result results_links results_links_deep web-result"> 선택
2. 각 결과 div에 대해:
   a. 제목: <h2 class="result__title"> > <a> > 텍스트
   b. 제목 href: <h2 class="result__title"> > <a> > href (리다이렉트 URL)
   c. 표시 URL: <a class="result__url"> > 텍스트
   d. 파비콘 src: <img class="result__icon__img"> > src
   e. 스니펫: <a class="result__snippet"> > 텍스트/내부 HTML
```

### 패턴: CSS 선택자 방식 (DOM 쿼리)

```javascript
// 모든 결과 가져오기
const results = document.querySelectorAll('div.result.results_links.web-result');

results.forEach((result) => {
  const title = result.querySelector('h2.result__title a.result__a').textContent;
  const href = result.querySelector('h2.result__title a.result__a').href;
  const displayUrl = result.querySelector('a.result__url').textContent;
  const faviconSrc = result.querySelector('img.result__icon__img')?.src;
  const snippet = result.querySelector('a.result__snippet').textContent;
});
```

### 패턴: 리다이렉트에서 실제 URL 추출

```
1. href 가져오기: //duckduckgo.com/l/?uddg=<ENCODED>&rut=<HASH>
2. "uddg" 쿼리 파라미터 값 추출
3. 값을 URL-디코딩
4. 결과: 실제 대상 URL
```

### 패턴: 파비콘 URL 처리

```
원본 형식: //external-content.duckduckgo.com/ip3/<DOMAIN>.ico
예시: //external-content.duckduckgo.com/ip3/www.python.org.ico

결과 URL에서 도메인을 추출해 파비콘 URL 구성:
- URL이 https://example.com/path인 경우
- 파비콘: //external-content.duckduckgo.com/ip3/example.com.ico
```

---

## 9) HTML 엔티티 및 인코딩

### 결과에 포함된 일반 HTML 엔티티
- `&#x27;` → `'` (아포스트로피)
- `&amp;` → `&` (앰퍼샌드)
- `&lt;` → `<` (미만)
- `&gt;` → `>` (초과)
- `&quot;` → `"` (따옴표)

### 쿼리 키워드 강조
- 볼드 태그: `<b>python</b>` 일치하는 키워드 강조
- 스니펫 텍스트에 나타날 수 있음
- 처리 전 HTML 엔티티 디코딩

### 링크의 URL 인코딩
- `uddg` 파라미터는 URL 인코딩됨
- 읽을 수 있는 URL을 얻으려면 URL-디코딩 필요
- 실제 대상 URL은 `uddg` URL-디코딩으로 얻음

---

## 10) 메타 정보

### 페이지 제목
- **선택자**: `<title>`
- **형식**: `<query> at DuckDuckGo`
- **예시**: `python at DuckDuckGo`

### 메타 태그
- **Robots**: `<meta name="robots" content="noindex, nofollow" />`
- **Referrer**: `<meta name="referrer" content="origin" />`
- **Viewport**: `<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=3.0, user-scalable=1" />`
- **Charset**: UTF-8

### OpenSearch 발견
- **선택자**: `<link rel="search" ... href="//duckduckgo.com/opensearch_html_v2.xml">`
- **타입**: `application/opensearchdescription+xml`
- **제목**: `DuckDuckGo (HTML)`

### 스타일시트
- **주 CSS**: 해시 파일명, 예: `//duckduckgo.com/dist/h.238c80a7d9b754cfcdd5.css`
- **미디어**: `handheld, all`

---

## 11) 빈 상태 및 엣지 경우

### 결과 없음
- 결과 컨테이너는 비어있음 (자식 결과 `<div>` 없음)
- 검색 폼 및 헤더는 여전히 표시됨

### 단일 결과
- 여러 결과와 동일한 구조
- 단일 결과 `<div>`만 포함

### 페이지네이션
- `#links.results` 끝부분의 `<div class="nav-link">` 내부 POST 폼으로 제공됨
- 관찰된 hidden 필드: `q`, `s`, `nextParams`, `v`, `o`, `dc`, `api`, `vqd`, 선택적 `kl`

### 봇 챌린지
- 봇 탐지 트리거되면 다른 페이지 구조
- 챌린지 모달 또는 다른 응답 구조 확인

---

## 12) 구조 비교: Lite vs HTML

| 항목 | Lite (`lite.duckduckgo.com`) | HTML (`html.duckduckgo.com`) |
|---|---|---|
| 결과 컨테이너 | 테이블 행(`<tr>`) | 결과 블록(`<div class="result...">`) |
| 제목 선택자 | `a.result-link` | `h2.result__title > a.result__a` |
| 스니펫 선택자 | `td.result-snippet` | `a.result__snippet` |
| 표시 URL 선택자 | `span.link-text` | `a.result__url` |
| 파비콘 노드 | 없음 | `img.result__icon__img` |
| 페이지네이션 컨테이너 | `form.next_form` | `div.nav-link > form` |
