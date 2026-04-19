# DuckDuckGo Lite Result Page Structure Analysis

Reference Date: 2026-04-19  
Target: `https://lite.duckduckgo.com/lite/` search results

## 1) Overall Page Structure

### Basic Layout

```
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" ...>
<html>
  <head>
    <!-- Metadata and stylesheets -->
  </head>
  <body>
    <p class='extra'>&nbsp;</p>
    <div class="header">DuckDuckGo</div>
    <p class='extra'>&nbsp;</p>
    
    <!-- Search Form -->
    <form action="/lite/" method="post">
      <input type="text" name="q" />
      <input type="submit" value="Search" />
      <div class="filters">
        <select name="kl"><!-- Region --></select>
        <select name="df"><!-- Time filter --></select>
      </div>
    </form>
    
    <!-- Navigation Section -->
    <p class="extra">&nbsp;</p>
    <table border="0">
      <tr>
        <td><!-- Pagination info (if any) --></td>
        <td><!-- Next/Previous Page Forms --></td>
      </tr>
    </table>
    
    <!-- Empty Results Table -->
    <p class='extra'>&nbsp;</p>
    <table border="0"><!-- Instant Answers or Special Results (if any) --></table>
    
    <!-- Web Results Table -->
    <table border="0">
      <!-- Result entries -->
    </table>
  </body>
</html>
```

---

## 2) Search Form Structure

### Container
- **Selector**: `form[action="/lite/"][method="post"]`
- **Purpose**: Main search form to submit new queries

### Elements

| Element | Selector | Type | Attributes | Notes |
|---------|----------|------|------------|-------|
| Query input | `input[name="q"][type="text"]` | text | `class="query"`, `size="40"` | Contains current search query; value persists from previous search |
| Submit button | `input[type="submit"]` | submit | `class="submit"`, `value="Search"` | Submits form with POST |
| Region select | `select[name="kl"]` | select | `class="submit"` | Dropdown with region codes (e.g., `kr-kr`, `us-en`, `wt-wt` for no region) |
| Time filter select | `select[name="df"]` | select | `class="submit"` | Options: `""` (Any Time), `d` (Past Day), `w` (Past Week), `m` (Past Month), `y` (Past Year) |

---

## 3) Navigation/Pagination Section

### Container
- **Selector**: `table border="0"` (first occurrence after search form)
- **Structure**: Single `<tr>` with two `<td>` cells

### Pagination Form (Next Page)
- **Selector**: `form.next_form[action="/lite/"][method="post"]`
- **Placement**: Inside second `<td>` of navigation table

### Form Fields (All `<input type="hidden">`)

| Field | Parameter | Type | Purpose | Example Value |
|-------|-----------|------|---------|----------------|
| Next button | (submit button) | submit | Triggers pagination | `value="Next Page &gt;"` |
| Search query | `q` | string | Maintains search query | `"python"` |
| Start offset | `s` | integer | Results offset (10, 20, 30, ...) | `10` |
| Cursor param | `nextParams` | string | Additional pagination data | `""` (usually empty) |
| Version flag | `v` | string | Internal lite version | `"l"` |
| Response mode | `o` | string | Output format | `"json"` |
| Counter | `dc` | integer | Internal cursor value | `"11"` |
| API identifier | `api` | string | Internal API mode | `"d.js"` |
| Query token | `vqd` | string | Server-issued validation token | Hash string |
| Region (optional) | `kl` | string | Region code if set | `"wt-wt"`, `"kr-kr"`, etc. |

---

## 4) Web Results Container

### Container
- **Selector**: `table border="0"` (last occurrence, contains all results)
- **Structure**: No thead/tbody, direct `<tr>` and `<td>` children

### Single Result Block Structure

Each result is rendered as a series of `<tr>` rows grouped together:

```html
<!-- Result row 1: Position + Title Link -->
<tr>
  <td valign="top">
    1.&nbsp;  <!-- Position number -->
  </td>
  <td>
    <a rel="nofollow" href="//duckduckgo.com/l/?uddg=...&rut=..." class='result-link'>
      Welcome to Python.org
    </a>
  </td>
</tr>

<!-- Result row 2: Snippet/Description (if present) -->
<tr>
  <td>&nbsp;&nbsp;&nbsp;</td>
  <td class='result-snippet'>
    Python is a versatile and easy-to-learn language...
  </td>
</tr>

<!-- Result row 3: Domain/Display URL -->
<tr>
  <td>&nbsp;&nbsp;&nbsp;</td>
  <td>
    <span class='link-text'>www.python.org</span>
  </td>
</tr>

<!-- Separator row -->
<tr>
  <td>&nbsp;</td>
  <td>&nbsp;</td>
</tr>
```

---

## 5) Result Element Details

### Result Position
- **Selector**: First `<td>` of result block > text content
- **Format**: Integer followed by `.&nbsp;` (e.g., `1.&nbsp;`)
- **Type**: text node
- **Extraction**: Parse number before period

### Result Title/Link
- **Selector**: `a.result-link`
- **Attributes**:
  - `rel="nofollow"`: Always present
  - `href`: Redirect URL with DDG parameters
    - Format: `//duckduckgo.com/l/?uddg=<url_encoded_redirect_url>&rut=<tracking_hash>`
    - `uddg`: URL-encoded destination URL
      - After decoding, this is the canonical extraction source for the target URL
      - `span.link-text` is a display value and can differ (scheme/query/normalization differences)
    - `rut`: Click tracking/attribution hash

### Result URL Extraction
- **Redirect URL structure**: `//duckduckgo.com/l/?uddg=<encoded>&rut=<hash>`
- **Encoded parameter name**: `uddg`
- **Quick path (easy)**: use `span.link-text` as display URL/domain
- **Exact path (strict)**: URL-decode `uddg` from `a.result-link[href]`
- **Note**: `span.link-text` can omit scheme and may differ from full canonical destination URL

### Result Snippet
- **Selector**: `td.result-snippet`
- **Content**: HTML with `<b>` tags around query keywords
- **Format**: Plain text with bold markup
- **Presence**: Optional; may not appear for all results
- **Extraction**: Remove HTML tags or keep them depending on use case

### Display URL / Domain
- **Selector**: `span.link-text`
- **Content**: Domain/path display URL (not full redirect URL)
- **Format**: `www.domain.com` or `www.domain.com/path`
- **Note**: For display purposes; actual destination must be extracted from link href

---

## 6) CSS Classes Reference

| Class | Element | Purpose |
|-------|---------|---------|
| `extra` | `<p>` | Spacing/padding paragraph |
| `header` | `<div>` | Page header with site name |
| `query` | `<input>` | Search query text input |
| `submit` | `<input>`, `<select>` | Submit button or select controls |
| `filters` | `<div>` | Container for region/time filter selects |
| `next_form` | `<form>` | Pagination form for next page |
| `navbutton` | `<input>` | Pagination submit button |
| `result-link` | `<a>` | Result title hyperlink |
| `result-snippet` | `<td>` | Result description text |
| `link-text` | `<span>` | Display URL/domain name |

---

## 7) Instant Answers / Special Results

### Container
- **Selector**: `table border="0"` (second occurrence, before main results table)
- **Content**: May be empty or contain special answer blocks
- **Note**: Not fully detailed in standard searches; appears for specific query types

---

## 8) Extraction Patterns for Scrapers

### Pattern: Extract All Results

```
1. Find all <tr> groups in results table
2. For each group starting with result position:
   a. Position: text of first <td> (parse number)
   b. Title: <a class="result-link"> text
   c. Link href: <a class="result-link"> href attribute
   d. Snippet: <td class="result-snippet"> inner text (optional)
   e. Display URL: <span class="link-text"> text
   f. Separator detected: empty <tr> with only &nbsp;
```

### Pattern: URL Extraction Modes

```
Quick mode:
1. Read <span class="link-text"> text
2. Use as display URL/domain

Exact mode:
1. Read <a class="result-link"> href
2. Extract "uddg" query parameter
3. URL-decode the value
4. Result: actual destination URL
```

### Pattern: Extract by Row Count

```
Typical result block = 3 to 4 rows:
- Row 1: Position + Title (class="result-link")
- Row 2: Snippet (class="result-snippet") [optional]
- Row 3: Display URL (class="link-text")
- Row 4: Separator (empty)

Group rows by position number to identify result boundaries.
```

---

## 9) Meta Information

### Page Title
- **Selector**: `<title>`
- **Format**: `<query> at DuckDuckGo`
- **Example**: `python at DuckDuckGo`

### Meta Tags
- **Robots meta**: `<meta name="robots" content="noindex, nofollow" />`
- **Referrer**: `<meta name="referrer" content="origin" />`
- **Charset**: UTF-8

### Stylesheets
- **Main CSS**: Hashed filename, e.g., `//duckduckgo.com/dist/lr.48ddfe4eadf6a534e93f.css`
- **Note**: Lite version uses minimal CSS for fast loading

---

## 10) Encoding and Special Cases

### HTML Entities
- Query keyword highlights use `<b>` tags
- Special characters encoded as HTML entities (e.g., `&#x27;` for apostrophe)
- Snippet text may contain HTML entities; decode when processing

### URL Encoding
- Result links use `//` protocol-relative URLs
- Query parameters in redirect URL are properly encoded
- `uddg` parameter requires URL-decoding
- The decoded `uddg` value may not text-match `span.link-text` because display URL formatting can omit or normalize parts of the destination URL

### Pagination Token (vqd)
- **Type**: Alphanumeric hash string
- **Purpose**: Validates pagination requests (prevents abuse)
- **Requirement**: Must be included in subsequent pagination requests
- **Format**: Example: `4-316971612543929568047774606561803676280`

---

## 11) Empty States and Edge Cases

### No Results
- Results table is present but empty (no `<tr>` children)
- Search form and navigation still displayed

### Bot Challenge Page
- If bot detection triggered, entire response differs
- Contains `<div class="anomaly-modal">` instead of results
- Requires CAPTCHA completion before continuing

### Single Result Page
- Still contains "Next Page" form even if only partial results
- Form contains all state parameters for consistency
