# DuckDuckGo Parameters Reference (GET/POST)

Reference Date: 2026-04-19  
Documentation Source (GET): `https://duckduckgo.com/duckduckgo-help-pages/settings/params`  
Observation Target (POST): `https://lite.duckduckgo.com/lite/` and forms in response (`next_form`, `anomaly.js`)

## 1) GET Parameters (Official settings/params documentation)

Base search parameter:

- `q` (string): Search query

### Result Settings

| Parameter | Type | Description | Values |
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

| Parameter | Type | Description | Values |
|---|---|---|---|
| `kd` | integer | Redirect | `1`, `-1` |
| `kh` | integer | HTTPS | `1`, `-1` |
| `kg` | string | Address bar method | `g`(GET), `p`(POST) |
| `k5` | integer | Video Playback | `1`, `2`, `-1` |

### Colour Settings

| Parameter | Type | Description | Values |
|---|---|---|---|
| `kj` | string | Header color | `r3`,`d`,`g`,`g2`,`b`,`b2`,`r`,`r2`,`p`,`o`,`w` or color code |
| `kx` | string | URL color | `r`,`g`,`l`,`b`,`p`,`o`,`e` or color code |
| `k7` | string | Background color | `w`,`d`,`g`,`g2`,`b`,`b2`,`r`,`r2`,`p`,`o` or color code |
| `k8` | string | Text color | `g` or color code |
| `k9` | string | Links color | `g`,`b` or color code |
| `kaa` | string | Visited links color | `p`(default/purple per docs) or color code |

### Look & Feel Settings

| Parameter | Type | Description | Values |
|---|---|---|---|
| `kae` | string | Theme | `-1`,`c`,`r`,`d`,`t` or color code |
| `ks` | string | Size | `n`,`l`,`t`,`m`,`s` |
| `kw` | string | Width | `n`,`w`,`s` |
| `km` | string | Placement | `m`,`l` |
| `ka` | string | Link font | `a`,`c`,`g`,`h`,`p`,`n`,`e`,`s`,`o`,`t`,`b`,`v` or font name |
| `ku` | integer | Underline | `1`,`-1` |
| `kt` | string | Text font | `a`,`c`,`g`,`h`,`p`,`n`,`e`,`s`,`o`,`t`,`b`,`v` or font name |

### Interface Settings

| Parameter | Type | Description | Values |
|---|---|---|---|
| `ko` | string | Header display | `1`,`s`,`-1`,`-2` |
| `k1` | integer | Advertisements | `1`,`-1` |
| `kv` | string | Page numbers | `1`,`n`,`-1` |
| `kaj` | string | Units of Measure | `1`,`n`,`-1` |
| `t` | string | Source identifier | string |

---

## 2) POST Parameters (Based on lite.duckduckgo.com/lite observation)

### A. Basic Search Form (`POST /lite/`)

| Parameter | Type | Description |
|---|---|---|
| `q` | string | Search query |
| `kl` | string | Region/language (optional) |
| `df` | string | Time filter (optional: `d`,`w`,`m`,`y`, or empty) |

### B. Results Pagination Form (`POST /lite/`, `class="next_form"`)

| Parameter | Type | Description |
|---|---|---|
| `q` | string | Search query |
| `s` | integer | Start offset (for next page) |
| `nextParams` | string | Additional cursor parameter (typically empty string) |
| `v` | string | Internal version flag (observed: `l`) |
| `o` | string | Internal response mode (observed: `json`) |
| `dc` | integer | Internal cursor/counter value |
| `api` | string | Internal API identifier (observed: `d.js`) |
| `vqd` | string | Server-issued query token |
| `kl` | string | Region/language persistence (optional) |

### B-2. Results Pagination Form (`POST /html/`, `div.nav-link > form`)

Observed hidden fields match the lite pagination payload:
- `q`, `s`, `nextParams`, `v`, `o`, `dc`, `api`, `vqd`, optional `kl`

The HTML endpoint uses a submit button value `Next`, while lite uses `Next Page >`.

### C. Bot Challenge Form (`POST //duckduckgo.com/anomaly.js`)

| Parameter | Type | Description |
|---|---|---|
| `image-check_<id>` | boolean | CAPTCHA tile checkbox (dynamic name) |
| `challenge-submit` | string | Challenge submission value |

## 3) Notes

- GET parameters list reflects current publicly documented items from official Help documentation.
- POST parameters based on actual observation of lite page HTML forms; some internal fields may vary depending on bot detection/experimental settings.
