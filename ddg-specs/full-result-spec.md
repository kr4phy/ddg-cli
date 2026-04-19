# DuckDuckGo Full SERP (`duckduckgo.com`) Technical Specification

Reference Date: 2026-04-19  
Target: `https://duckduckgo.com/?q=...` (full JavaScript SERP)

## 1) Rendering Model

- The full SERP is JavaScript-driven.
- Initial HTML contains shell containers and bootstrap scripts.
- Web results are not embedded as static `<div class="result ...">` nodes in the initial response.
- Runtime script initializes deep fetching via `DDG.deep.initialize('/d.js?...', true)`.

---

## 2) Initial HTML Structure (Verified)

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

### Core Containers

| Selector | Role |
|---|---|
| `#react-search-form` | Search form mount point |
| `#react-duckbar` | Navigation/duckbar mount point |
| `#react-root-zci` | Instant-answer mount point |
| `#vertical_wrapper` | Vertical tabs mount point |
| `#react-layout` | Main web results mount point |

---

## 3) Bootstrap Variables in Initial HTML (Verified)

Inline script exposes search/session state:

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

Also present:
- `backendExperimentAssignments` object
- `_bootstrapBackendData`
- `window.__preloadData__`
- `window.__initialSearchFormData__`

---

## 4) Runtime Data Fetch Bootstrap (Verified in Initial HTML)

Inline `nrji()` script triggers:

```javascript
nrj('/t.js?...');
DDG.deep.initialize('/d.js?...&wrap=1&...', true);
DDG.ready(nrji, 1);
```

### Observed `/d.js` query fields

| Parameter | Meaning (observed context) |
|---|---|
| `q` | Query text |
| `t` | Source/type marker |
| `l` | Locale code |
| `s` | Start offset |
| `ct` | Country |
| `bing_market` | Backend market |
| `dp` | Internal token |
| `wpa`, `wpl` | Internal context |
| `perf_id`, `parent_perf_id`, `perf_sampled` | Performance/session tracking |
| `host_region` | Host region |
| `sp`, `dfrsp`, `wrap`, `aps` | Internal response/format flags |

Detailed endpoint probes are documented separately:
- `djs-api-spec.md`
- `tjs-api-spec.md`

---

## 5) Noscript Fallback (Verified)

Inside `data-testid="mainline"`:

```html
<noscript>
  <meta http-equiv="refresh" content="0;URL=/html?q=python">
  <div class="msg msg--noscript">...</div>
</noscript>
```

If JavaScript is unavailable, flow redirects to `/html`.

---

## 6) Anti-bot Hook (Verified)

Tail script includes:

```javascript
DDG.deep.anomalyDetectionBlock({...});
```

This is part of runtime anomaly/bot handling on full SERP.

---

## 7) What Is and Is Not in Initial HTML

### Present
- Query/session/config variables
- React mount containers
- Bootstrap scripts and runtime endpoints (`/t.js`, `/d.js`)
- Noscript redirect path

### Not Present
- Fully rendered web result cards as static HTML under `#react-layout`
- Static title/snippet/url nodes for each web result in initial response

---

## 8) Extraction Boundaries

### Extractable from initial response only
- Query/state metadata (`rq`, `vqd`, locale, perf ids)
- Runtime endpoint templates and parameters
- Container IDs/selectors and page shell structure

### Requires JavaScript runtime DOM
- Final rendered result card nodes in `#react-layout`
- Runtime-only attributes emitted by the client renderer
- Post-render ordering and pagination UI state
