# DuckDuckGo `/t.js` Endpoint Technical Specification

Reference Date: 2026-04-19  
Scope: Probe-based specification for `https://duckduckgo.com/t.js`

## 1) Endpoint Identity

- Host: `duckduckgo.com`
- Path: `/t.js`
- Method: `GET`
- Bootstrap call site in full SERP:
  - `nrj('/t.js?...')`

Observed query keys from bootstrap URL:

`q, t, l, s, ct, bing_market, p_ent, ex, dp, wpa, wpl, perf_id, parent_perf_id, perf_sampled, host_region, dfrsp, aps`

---

## 2) `/t.js` ↔ `/d.js` Sequencing and Role Split

In the same bootstrap function (`nrji()`), call order is:

1. `nrj('/t.js?...')`
2. `DDG.deep.initialize('/d.js?...', true)`

Observed index order in parsed function body:
- `t.js` call index: `0`
- `d.js` call index: `416`
- `t.js` before `d.js`: `true`

Parameter-set relationship:

- Keys in both endpoints:
  - `q, t, l, s, ct, bing_market, p_ent, ex, dp, wpa, wpl, perf_id, parent_perf_id, perf_sampled, host_region, dfrsp, aps`
- Keys only observed on `/d.js`:
  - `vqd, sp, wrap`
- Keys only observed on `/t.js`:
  - _(none in sampled bootstrap URL)_

Interpretation from this probe:
- `/t.js` is the earlier companion call.
- `/d.js` carries additional result/deep-render control keys (`vqd`, `sp`, `wrap`).

---

## 3) Response Envelope (Observed)

### 3.1 Context variants

| Case | Status | Content-Type | Body length |
|---|---:|---|---:|
| cookie + referer + UA | 200 | `application/x-javascript` | 0 |
| cookie + UA | 200 | `application/x-javascript` | 0 |
| UA only | 200 | `application/x-javascript` | 0 |

### 3.2 Parameter sensitivity probes

| Variant | Status | Content-Type | Body length |
|---|---:|---|---:|
| baseline | 200 | `application/x-javascript` | 0 |
| remove `q` | 200 | `application/x-javascript` | 0 |
| remove `dp` | 200 | `application/x-javascript` | 0 |
| remove `perf_id` | 200 | `application/x-javascript` | 0 |
| `host_region=use` | 200 | `application/x-javascript` | 0 |

Observed behavior in this probe window:
- `/t.js` responded with empty body across tested contexts/variants.
- No challenge-script payload family was observed from `/t.js` itself.

---

## 4) Probe-Scope Conclusion

- `/t.js` appears as an earlier bootstrap-side companion request in full SERP initialization.
- `/d.js` is the endpoint where deep/result challenge script behavior was observed.
- In this probe, `/t.js` was metadata/telemetry-like in behavior (stable `200`, empty body).

For deep endpoint behavior and challenge transitions, see:
- `djs-api-spec.md`
