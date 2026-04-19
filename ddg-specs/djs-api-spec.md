# DuckDuckGo `/d.js` Endpoint Technical Specification

Reference Date: 2026-04-19  
Scope: Probe-based specification for `https://duckduckgo.com/d.js`

## 1) Endpoint Identity

- Host: `duckduckgo.com`
- Path: `/d.js`
- Method: `GET`
- Bootstrap call site in full SERP:
  - `DDG.deep.initialize('/d.js?...', true)`

Observed bootstrap query keys:

`q, t, l, s, ct, vqd, bing_market, p_ent, ex, dp, wpa, wpl, perf_id, parent_perf_id, perf_sampled, host_region, sp, dfrsp, wrap, aps`

---

## 2) Response Envelope (Observed)

- Status: `202` (probe responses)
- Content-Type: `application/x-javascript; charset=UTF-8`
- Payload type: JavaScript source (not plain JSON)

Common payload families:
1. `window.execDeep ... is506`
2. `window.execDeep ... isJsaChallenge`
3. `window.execDeep ... anomalyDetectionBlock`
4. JSA script fragments that end with `DDG.deep.initialize(...&jsa=..., false)`

---

## 3) State Transition Matrix (Headers/Cookies)

First request used the full bootstrap `/d.js?...` URL.

| Case | First response | First len | Second-stage behavior |
|---|---|---:|---|
| cookie + referer + UA | `isJsaChallenge` | 2125 | computed `jsa` follow-up => `anomalyDetectionBlock` (len 539) |
| cookie + UA | `isJsaChallenge` | 2340 | computed `jsa` follow-up => `anomalyDetectionBlock` (len 537) |
| referer + UA | `isJsaChallenge` | 2535 | computed `jsa` follow-up => `anomalyDetectionBlock` (len 541) |
| UA only | `isJsaChallenge` | 2559 | computed `jsa` follow-up => `anomalyDetectionBlock` (len 543) |
| cookie + referer, no UA | `anomalyDetectionBlock` | 538 | no JSA step observed |

### Observed transition pattern

- Typical observed sequence:
  - `isJsaChallenge` → client computes `jsa` → follow-up `/d.js?...&jsa_hash=...&jsa=...` → `anomalyDetectionBlock`
- UA omission changed first-step class to immediate anomaly block in this probe set.

---

## 4) Parameter Sensitivity Matrix

Baseline context: cookie + referer + UA.

| Variant | Status | Body len | Class | Key observation |
|---|---:|---:|---|---|
| baseline | 202 | 1898 | `isJsaChallenge` | normal challenge entry |
| remove `vqd` | 202 | 2462 | `isJsaChallenge` | still challenge script |
| mutate `vqd` | 202 | 2456 | `isJsaChallenge` | still challenge script |
| remove `dp` | 202 | 1377 | other/script-fragment | generated follow-up contained malformed prefix (`'1&jsa_hash=...'`) |
| mutate `dp` | 202 | 2775 | `isJsaChallenge` | challenge still issued |
| `s=30` | 202 | 2464 | `isJsaChallenge` | offset does not bypass challenge |
| `s=90` | 202 | 2358 | `isJsaChallenge` | offset does not bypass challenge |
| `wrap=0` | 202 | 2368 | other/script-fragment | `window.execDeep` wrapper not present in sampled response |
| remove `wrap` | 202 | 2561 | other/script-fragment | wrapperless JSA fragment observed |
| `host_region=use` | 202 | 2068 | `isJsaChallenge` | challenge maintained |
| remove `host_region` | 202 | 2216 | `isJsaChallenge` | challenge maintained |

---

## 5) Pagination Stability with Same `vqd`

Tested offsets with one `vqd` value:

`s = 0, 10, 20, ... , 100`

Observed result:
- All requests returned status `202`
- All sampled classes were `isJsaChallenge`
- Body length varied per request, class family stayed consistent

`vqd` sample check (three fresh page loads for same query):
- Sampled `vqd` remained identical across the three immediate reloads.

---

## 6) Locale/Region Axis (`l`, `ct`, `kl`, `bing_market`)

### 6.1 Full-page bootstrap mapping (`?q=python&kl=...`)

| `kl` input | `d.js l` | `d.js ct` | `d.js bing_market` | `d.js wpl` | Class |
|---|---|---|---|---|---|
| `us-en` | `us-en` | `KR` | `en-US` | `en` | `isJsaChallenge` |
| `kr-kr` | `kr-kr` | `KR` | `ko-KR` | _(none)_ | `isJsaChallenge` |
| `jp-jp` | `jp-jp` | `KR` | `ja-JP` | _(none)_ | `isJsaChallenge` |
| `fr-fr` | `fr-fr` | `KR` | `fr-FR` | _(none)_ | `isJsaChallenge` |
| `wt-wt` | `us-en` | `KR` | `en-US` | `en` | `isJsaChallenge` |

Observed behavior:
- `kl` strongly maps to `l` and `bing_market`.
- `ct` remained `KR` in this environment.
- `wpl` appeared for English mapping and was absent in sampled non-English mappings.

### 6.2 HTML SERP structure check across `kl` (sampled pages)

Across sampled `kl` pages (`us-en`, `kr-kr`, `jp-jp`, `fr-fr`, `wt-wt`), the DOM shape remained the same:
- `div#links.results`
- `div.result.results_links.results_links_deep.web-result`
- `h2.result__title > a.result__a`
- `a.result__url`
- `a.result__snippet`
- `img.result__icon__img`

What changed by locale:
- ranking/content language/snippets (e.g., locale-specific titles/snippets)
- optional metadata text near URL block (observed in some locale results)

---

## 7) Probe-Scope Conclusion

In this probe window, `/d.js` behaved as a challenge-script endpoint, not a direct static JSON payload endpoint:
- script-driven anti-automation branching (`is506`, `isJsaChallenge`, anomaly block)
- behavior influenced by parameter shape and request context
- locale inputs changed query mapping fields but not challenge family

For `/t.js` sequencing and endpoint behavior, see:
- `tjs-api-spec.md`
