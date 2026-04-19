# DuckDuckGo 파라미터 정리 (GET/POST)

기준 시각: 2026-04-19  
참고 문서(GET): `https://duckduckgo.com/duckduckgo-help-pages/settings/params`  
관찰 대상(POST): `https://lite.duckduckgo.com/lite/` 및 응답 내 폼(`next_form`, `anomaly.js`)

## 1) GET 파라미터 (공식 settings/params 기준)

기본 검색어 파라미터:

- `q` (string): 검색어

### Result Settings

| 파라미터 | 타입 | 의미 | 값 |
|---|---|---|---|
| `kl` | string | Region | `xa-ar`, `xa-en`, `ar-es`, `au-en`, `at-de`, `be-fr`, `be-nl`, `br-pt`, `bg-bg`, `ca-en`, `ca-fr`, `ct-ca`, `cl-es`, `cn-zh`, `co-es`, `hr-hr`, `cz-cs`, `dk-da`, `ee-et`, `fi-fi`, `fr-fr`, `de-de`, `gr-el`, `hk-tzh`, `hu-hu`, `in-en`, `id-id`, `id-en`, `ie-en`, `il-he`, `it-it`, `jp-jp`, `kr-kr`, `lv-lv`, `lt-lt`, `xl-es`, `my-ms`, `my-en`, `mx-es`, `nl-nl`, `nz-en`, `no-no`, `pe-es`, `ph-en`, `ph-tl`, `pl-pl`, `pt-pt`, `ro-ro`, `ru-ru`, `sg-en`, `sk-sk`, `sl-sl`, `za-en`, `es-es`, `se-sv`, `ch-de`, `ch-fr`, `ch-it`, `tw-tzh`, `th-th`, `tr-tr`, `ua-uk`, `uk-en`, `us-en`, `ue-es`, `ve-es`, `vn-vi`, `wt-wt` |
| `kp` | integer | Safe Search | `1`(On), `-1`(Moderate), `-2`(Off) |
| `kz` | integer | Open Instant Answers | `1`, `-1` |
| `kc` | integer | Auto-load Images | `1`, `-1` |
| `kav` | integer | Auto-load Results | `1`, `-1` |
| `kn` | integer | New Window | `1`, `-1` |
| `kf` | string | Favicons/WOT | `1`, `w`, `fw`, `-1` |
| `kaf` | integer | Full URLs | `1`, `-1` |
| `kac` | integer | Auto-suggest | `1`, `-1` |

### Privacy Settings

| 파라미터 | 타입 | 의미 | 값 |
|---|---|---|---|
| `kd` | integer | Redirect | `1`, `-1` |
| `kh` | integer | HTTPS | `1`, `-1` |
| `kg` | string | Address bar method | `g`(GET), `p`(POST) |
| `k5` | integer | Video Playback | `1`, `2`, `-1` |

### Colour Settings

| 파라미터 | 타입 | 의미 | 값 |
|---|---|---|---|
| `kj` | string | Header color | `r3`,`d`,`g`,`g2`,`b`,`b2`,`r`,`r2`,`p`,`o`,`w` 또는 색 코드 |
| `kx` | string | URL color | `r`,`g`,`l`,`b`,`p`,`o`,`e` 또는 색 코드 |
| `k7` | string | Background color | `w`,`d`,`g`,`g2`,`b`,`b2`,`r`,`r2`,`p`,`o` 또는 색 코드 |
| `k8` | string | Text color | `g` 또는 색 코드 |
| `k9` | string | Links color | `g`,`b` 또는 색 코드 |
| `kaa` | string | Visited links color | `p`(문서상 기본/보라), 또는 색 코드 |

### Look & Feel Settings

| 파라미터 | 타입 | 의미 | 값 |
|---|---|---|---|
| `kae` | string | Theme | `-1`,`c`,`r`,`d`,`t` 또는 색 코드 |
| `ks` | string | Size | `n`,`l`,`t`,`m`,`s` |
| `kw` | string | Width | `n`,`w`,`s` |
| `km` | string | Placement | `m`,`l` |
| `ka` | string | Link font | `a`,`c`,`g`,`h`,`p`,`n`,`e`,`s`,`o`,`t`,`b`,`v` 또는 폰트명 |
| `ku` | integer | Underline | `1`,`-1` |
| `kt` | string | Text font | `a`,`c`,`g`,`h`,`p`,`n`,`e`,`s`,`o`,`t`,`b`,`v` 또는 폰트명 |

### Interface Settings

| 파라미터 | 타입 | 의미 | 값 |
|---|---|---|---|
| `ko` | string | Header 표시 | `1`,`s`,`-1`,`-2` |
| `k1` | integer | Advertisements | `1`,`-1` |
| `kv` | string | Page #s | `1`,`n`,`-1` |
| `kaj` | string | Units of Measure | `1`,`n`,`-1` |
| `t` | string | Source 식별자 | 문자열 |

---

## 2) POST 파라미터 (`lite.duckduckgo.com/lite` 관찰 기준)

### A. 기본 검색 폼 (`POST /lite/`)

| 파라미터 | 타입 | 의미 |
|---|---|---|
| `q` | string | 검색어 |
| `kl` | string | 지역/언어(선택) |
| `df` | string | 기간 필터(선택: `d`,`w`,`m`,`y`, 빈값) |

### B. 결과 페이지네이션 폼 (`POST /lite/`, `class="next_form"`)

| 파라미터 | 타입 | 의미 |
|---|---|---|
| `q` | string | 검색어 |
| `s` | integer | 시작 오프셋(다음 페이지) |
| `nextParams` | string | 추가 커서 파라미터(대개 빈 문자열) |
| `v` | string | 내부 버전 플래그(관찰값: `l`) |
| `o` | string | 내부 응답 모드(관찰값: `json`) |
| `dc` | integer | 내부 커서/카운터 값 |
| `api` | string | 내부 API 식별(관찰값: `d.js`) |
| `vqd` | string | 서버 발급 쿼리 토큰 |
| `kl` | string | 지역/언어 유지(선택) |

### B-2. 결과 페이지네이션 폼 (`POST /html/`, `div.nav-link > form`)

관찰된 hidden 필드는 lite 페이지네이션과 동일:
- `q`, `s`, `nextParams`, `v`, `o`, `dc`, `api`, `vqd`, 선택적 `kl`

HTML 엔드포인트는 제출 버튼 값이 `Next`, lite는 `Next Page >`로 관찰됨.

### C. 봇 챌린지 폼 (`POST //duckduckgo.com/anomaly.js`)

| 파라미터 | 타입 | 의미 |
|---|---|---|
| `image-check_<id>` | boolean | CAPTCHA 타일 체크박스 (동적 이름) |
| `challenge-submit` | string | 챌린지 제출 값 |

## 3) 참고/주의

- GET 파라미터 목록은 공식 Help 문서의 현재 공개 항목을 그대로 반영했다.
- POST 파라미터는 lite 페이지 HTML 폼을 실제 관찰한 값으로, 봇 탐지/실험 설정에 따라 일부 내부 필드는 바뀔 수 있다.
