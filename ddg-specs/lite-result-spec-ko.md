# DuckDuckGo Lite 결과 페이지 구조 분석

기준 시각: 2026-04-19  
대상: `https://lite.duckduckgo.com/lite/` 검색 결과

## 1) 전체 페이지 구조

### 기본 레이아웃

```
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" ...>
<html>
  <head>
    <!-- 메타데이터 및 스타일시트 -->
  </head>
  <body>
    <p class='extra'>&nbsp;</p>
    <div class="header">DuckDuckGo</div>
    <p class='extra'>&nbsp;</p>
    
    <!-- 검색 폼 -->
    <form action="/lite/" method="post">
      <input type="text" name="q" />
      <input type="submit" value="Search" />
      <div class="filters">
        <select name="kl"><!-- 지역 --></select>
        <select name="df"><!-- 시간 필터 --></select>
      </div>
    </form>
    
    <!-- 네비게이션 섹션 -->
    <p class="extra">&nbsp;</p>
    <table border="0">
      <tr>
        <td><!-- 페이지네이션 정보 --></td>
        <td><!-- 다음/이전 페이지 폼 --></td>
      </tr>
    </table>
    
    <!-- 빈 결과 테이블 -->
    <p class='extra'>&nbsp;</p>
    <table border="0"><!-- 즉시 답변 또는 특수 결과 --></table>
    
    <!-- 웹 결과 테이블 -->
    <table border="0">
      <!-- 결과 항목들 -->
    </table>
  </body>
</html>
```

---

## 2) 검색 폼 구조

### 컨테이너
- **선택자**: `form[action="/lite/"][method="post"]`
- **목적**: 새 쿼리 제출용 주 검색 폼

### 요소

| 요소 | 선택자 | 타입 | 속성 | 설명 |
|------|--------|------|------|------|
| 검색어 입력 | `input[name="q"][type="text"]` | text | `class="query"`, `size="40"` | 현재 검색어 포함; 이전 검색에서 값 유지 |
| 제출 버튼 | `input[type="submit"]` | submit | `class="submit"`, `value="Search"` | POST로 폼 제출 |
| 지역 선택 | `select[name="kl"]` | select | `class="submit"` | 지역 코드 드롭다운 (예: `kr-kr`, `us-en`, `wt-wt`) |
| 시간 필터 | `select[name="df"]` | select | `class="submit"` | 옵션: `""` (모든 시간), `d` (지난 1일), `w` (지난 1주), `m` (지난 1개월), `y` (지난 1년) |

---

## 3) 네비게이션/페이지네이션 섹션

### 컨테이너
- **선택자**: `table border="0"` (검색 폼 바로 다음)
- **구조**: 단일 `<tr>` 내 두 개 `<td>` 셀

### 페이지네이션 폼 (다음 페이지)
- **선택자**: `form.next_form[action="/lite/"][method="post"]`
- **위치**: 네비게이션 테이블의 두 번째 `<td>` 내부

### 폼 필드 (모두 `<input type="hidden">`)

| 필드 | 파라미터 | 타입 | 목적 | 예시 값 |
|------|---------|------|------|---------|
| 다음 버튼 | (submit button) | submit | 페이지네이션 트리거 | `value="Next Page &gt;"` |
| 검색어 | `q` | string | 검색어 유지 | `"python"` |
| 시작 오프셋 | `s` | integer | 결과 오프셋 (10, 20, 30, ...) | `10` |
| 커서 파라미터 | `nextParams` | string | 추가 페이지네이션 데이터 | `""` (대개 비어있음) |
| 버전 플래그 | `v` | string | 내부 lite 버전 | `"l"` |
| 응답 모드 | `o` | string | 출력 형식 | `"json"` |
| 카운터 | `dc` | integer | 내부 커서 값 | `"11"` |
| API 식별자 | `api` | string | 내부 API 모드 | `"d.js"` |
| 쿼리 토큰 | `vqd` | string | 서버 발급 검증 토큰 | 해시 문자열 |
| 지역 (선택) | `kl` | string | 지역 코드 (설정 시) | `"wt-wt"`, `"kr-kr"` 등 |

---

## 4) 웹 결과 컨테이너

### 컨테이너
- **선택자**: `table border="0"` (마지막 occurrence, 모든 결과 포함)
- **구조**: thead/tbody 없음, 직접 `<tr>` 및 `<td>` 자식

### 단일 결과 블록 구조

각 결과는 `<tr>` 행 시리즈로 렌더링됨:

```html
<!-- 결과 행 1: 위치 + 제목 링크 -->
<tr>
  <td valign="top">
    1.&nbsp;  <!-- 위치 번호 -->
  </td>
  <td>
    <a rel="nofollow" href="//duckduckgo.com/l/?uddg=...&rut=..." class='result-link'>
      Welcome to Python.org
    </a>
  </td>
</tr>

<!-- 결과 행 2: 스니펫/설명 (있으면) -->
<tr>
  <td>&nbsp;&nbsp;&nbsp;</td>
  <td class='result-snippet'>
    Python is a versatile and easy-to-learn language...
  </td>
</tr>

<!-- 결과 행 3: 도메인/표시 URL -->
<tr>
  <td>&nbsp;&nbsp;&nbsp;</td>
  <td>
    <span class='link-text'>www.python.org</span>
  </td>
</tr>

<!-- 구분자 행 -->
<tr>
  <td>&nbsp;</td>
  <td>&nbsp;</td>
</tr>
```

---

## 5) 결과 요소 상세

### 결과 위치
- **선택자**: 결과 블록의 첫 `<td>` > 텍스트 콘텐츠
- **형식**: 정수 뒤에 `.&nbsp;` (예: `1.&nbsp;`)
- **타입**: 텍스트 노드
- **추출**: 마침표 앞 숫자 파싱

### 결과 제목/링크
- **선택자**: `a.result-link`
- **속성**:
  - `rel="nofollow"`: 항상 존재
  - `href`: DDG 파라미터가 포함된 리다이렉트 URL
    - 형식: `//duckduckgo.com/l/?uddg=<url_encoded_redirect_url>&rut=<tracking_hash>`
    - `uddg`: URL 인코딩된 대상 URL
      - URL-디코딩 후 값이 실제 대상 URL 추출의 기준
      - `span.link-text`는 표시용 값이며 스킴/쿼리/정규화 차이로 다를 수 있음
    - `rut`: 클릭 추적/기여 해시

### 실제 URL 추출
- **리다이렉트 URL 구조**: `//duckduckgo.com/l/?uddg=<encoded>&rut=<hash>`
- **인코딩 파라미터명**: `uddg`
- **간편 경로(쉬움)**: `span.link-text`를 표시 URL/도메인으로 사용
- **정확 경로(엄밀)**: `a.result-link[href]`의 `uddg`를 URL-디코딩
- **주의**: `span.link-text`는 스킴이 없거나 전체 정규 URL과 다를 수 있음

### 결과 스니펫
- **선택자**: `td.result-snippet`
- **콘텐츠**: 쿼리 키워드를 감싼 `<b>` 태그가 포함된 HTML
- **형식**: `<b>` 마크업이 있는 순수 텍스트
- **존재 여부**: 선택사항; 모든 결과에 표시되지 않을 수 있음
- **추출**: 사용 사례에 따라 HTML 태그 제거 또는 유지

### 표시 URL / 도메인
- **선택자**: `span.link-text`
- **콘텐츠**: 도메인/경로 표시 URL (완전 리다이렉트 URL 아님)
- **형식**: `www.domain.com` 또는 `www.domain.com/path`
- **주의**: 표시용; 실제 대상은 링크 href에서 추출해야 함

---

## 6) CSS 클래스 참고

| 클래스 | 요소 | 목적 |
|--------|------|------|
| `extra` | `<p>` | 간격/패딩 |
| `header` | `<div>` | 사이트명 헤더 |
| `query` | `<input>` | 검색어 텍스트 입력 |
| `submit` | `<input>`, `<select>` | 제출 버튼 또는 선택 컨트롤 |
| `filters` | `<div>` | 지역/시간 필터 선택 컨테이너 |
| `next_form` | `<form>` | 다음 페이지 페이지네이션 폼 |
| `navbutton` | `<input>` | 페이지네이션 제출 버튼 |
| `result-link` | `<a>` | 결과 제목 하이퍼링크 |
| `result-snippet` | `<td>` | 결과 설명 텍스트 |
| `link-text` | `<span>` | 표시 URL/도메인명 |

---

## 7) 즉시 답변 / 특수 결과

### 컨테이너
- **선택자**: `table border="0"` (두 번째 occurrence, 주 결과 테이블 이전)
- **콘텐츠**: 비어있거나 특수 답변 블록 포함 가능
- **주의**: 표준 검색에서는 완전히 상세하지 않음; 특정 쿼리 유형에 표시됨

---

## 8) 스크래퍼용 추출 패턴

### 패턴: 모든 결과 추출

```
1. 결과 테이블에서 모든 <tr> 그룹 찾기
2. 결과 위치로 시작하는 각 그룹에 대해:
   a. 위치: 첫 <td>의 텍스트 (숫자 파싱)
   b. 제목: <a class="result-link"> 텍스트
   c. 링크 href: <a class="result-link"> href 속성
   d. 스니펫: <td class="result-snippet"> 내부 텍스트 (선택사항)
   e. 표시 URL: <span class="link-text"> 텍스트
   f. 구분자 감지: &nbsp;만 있는 빈 <tr>
```

### 패턴: URL 추출 모드

```
간편 모드:
1. <span class="link-text"> 텍스트 읽기
2. 표시 URL/도메인으로 사용

정확 모드:
1. <a class="result-link"> href 읽기
2. "uddg" 쿼리 파라미터 추출
3. 값 URL-디코딩
4. 결과: 실제 대상 URL
```

### 패턴: 행 개수로 추출

```
일반적인 결과 블록 = 3~4개 행:
- 행 1: 위치 + 제목 (class="result-link")
- 행 2: 스니펫 (class="result-snippet") [선택적]
- 행 3: 표시 URL (class="link-text")
- 행 4: 구분자 (비어있음)

위치 번호로 행을 그룹화해 결과 경계 식별.
```

---

## 9) 메타 정보

### 페이지 제목
- **선택자**: `<title>`
- **형식**: `<query> at DuckDuckGo`
- **예시**: `python at DuckDuckGo`

### 메타 태그
- **Robots 메타**: `<meta name="robots" content="noindex, nofollow" />`
- **Referrer**: `<meta name="referrer" content="origin" />`
- **Charset**: UTF-8

### 스타일시트
- **주 CSS**: 해시 파일명, 예: `//duckduckgo.com/dist/lr.48ddfe4eadf6a534e93f.css`
- **주의**: Lite 버전은 빠른 로딩을 위해 최소 CSS 사용

---

## 10) 인코딩 및 특수 경우

### HTML 엔티티
- 쿼리 키워드 강조는 `<b>` 태그 사용
- 특수 문자는 HTML 엔티티로 인코딩 (예: `&#x27;` for 아포스트로피)
- 스니펫 텍스트는 HTML 엔티티 포함 가능; 처리 시 디코딩

### URL 인코딩
- 결과 링크는 `//` 프로토콜 상대 URL 사용
- 리다이렉트 URL의 쿼리 파라미터는 올바르게 인코딩됨
- `uddg` 파라미터는 URL-디코딩 필요
- 디코딩된 `uddg` 값은 표시용 `span.link-text`와 텍스트가 완전히 일치하지 않을 수 있음(스킴/쿼리/정규화 차이)

### 페이지네이션 토큰 (vqd)
- **타입**: 영숫자 해시 문자열
- **목적**: 페이지네이션 요청 검증 (악용 방지)
- **필요성**: 후속 페이지네이션 요청에 포함되어야 함
- **형식**: 예: `4-316971612543929568047774606561803676280`

---

## 11) 빈 상태 및 엣지 경우

### 결과 없음
- 결과 테이블은 존재하지만 비어있음 (자식 `<tr>` 없음)
- 검색 폼 및 네비게이션은 여전히 표시됨

### 봇 챌린지 페이지
- 봇 탐지 트리거되면 전체 응답이 다름
- 결과 대신 `<div class="anomaly-modal">` 포함
- 계속하기 전에 CAPTCHA 완료 필요

### 단일 결과 페이지
- 부분 결과만 있어도 "다음 페이지" 폼 포함
- 폼은 일관성을 위해 모든 상태 파라미터 포함
