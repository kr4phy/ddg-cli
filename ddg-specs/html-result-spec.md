# DuckDuckGo HTML Result Page Structure Analysis

Reference Date: 2026-04-19  
Target: `https://html.duckduckgo.com/html/` search results

## 1) Overall Page Structure

### Basic Layout

```
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" ...>
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <!-- Metadata and stylesheets -->
  </head>
  <body class="body--html">
    <a name="top" id="top"></a>
    
    <!-- Hidden state form -->
    <form action="/html/" method="post">
      <input type="text" name="state_hidden" id="state_hidden" />
    </form>

    <div>
      <div class="site-wrapper-border"></div>

      <!-- Header with search form -->
      <div id="header" class="header cw header--html">
        <!-- Logo and Search Form -->
        <form name="x" class="header__form" action="/html/" method="post">
          <div class="search search--header">
            <input name="q" type="text" value="..." />
            <input name="b" type="submit" />
          </div>
          <div class="frm__select">
            <select name="kl"><!-- Region --></select>
          </div>
          <div class="frm__select frm__select--last">
            <select name="df"><!-- Time filter --></select>
          </div>
        </form>
      </div>

      <!-- Results Container -->
      <div>
        <div class="serp__results">
          <div id="links" class="results">
            <!-- Individual results -->
          </div>
        </div>
      </div>
    </div>
  </body>
</html>
```

---

## 2) Search Form Structure

### Header Form Container
- **Selector**: `form.header__form[action="/html/"][method="post"]`
- **Purpose**: Main search form with query, region, and time filters

### Form Elements

| Element | Selector | Type | Attributes | Notes |
|---------|----------|------|------------|-------|
| Query input | `input[name="q"][type="text"]` | text | `class="search__input"`, `id="search_form_input_homepage"`, `autocomplete="off"` | Current search query persists in value |
| Submit button | `input[name="b"][type="submit"]` | submit | `class="search__button search__button--html"`, `id="search_button_homepage"` | May have empty value attribute |
| Region select | `select[name="kl"]` | select | `class="frm__select"` (parent) | Dropdown with region codes |
| Time filter select | `select[name="df"]` | select | `class="frm__select frm__select--last"` (parent) | Options: "" (Any), d, w, m, y |

### Form Wrapper Classes
- `.search` - Search input wrapper
- `.search--header` - Header-specific search
- `.frm__select` - Filter select container (region)
- `.frm__select--last` - Last filter select container (time)

---

## 3) Results Container Structure

### Main Results Wrapper
- **Selector**: `div.serp__results > div#links.results`
- **Structure**: Container div with class `results` and id `links`
- **Content**: Sequence of individual result `<div>` elements

---

## 4) Individual Result Block Structure

### Container
- **Selector**: `div.result.results_links.results_links_deep.web-result`
- **Classes**:
  - `result` - Base result styling
  - `results_links` - Link-based result
  - `results_links_deep` - Deep linking result
  - `web-result` - Web search result

### Result Body
- **Selector**: `div.links_main.links_deep.result__body`
- **Purpose**: Container for all result content
- **Note**: Contains comment "This is the visible part"

### Single Result Block HTML Pattern

```html
<div class="result results_links results_links_deep web-result">
  <div class="links_main links_deep result__body">
    
    <!-- Title/Link -->
    <h2 class="result__title">
      <a rel="nofollow" class="result__a" href="//duckduckgo.com/l/?uddg=...&rut=...">
        Welcome to Python.org
      </a>
    </h2>

    <!-- URL and Favicon Section -->
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

    <!-- Snippet/Description -->
    <a class="result__snippet" href="//duckduckgo.com/l/?uddg=...&rut=...">
      Python knows the usual control flow statements...
    </a>

    <div class="clear"></div>
  </div>
</div>
```

---

## 5) Result Element Details

### Result Title
- **Selector**: `h2.result__title > a.result__a`
- **Attributes**:
  - `rel="nofollow"`: Always present
  - `href`: DDG redirect URL (format: `//duckduckgo.com/l/?uddg=<encoded>&rut=<hash>`)
  - `class="result__a"`: Title link styling
- **Content**: Link text (result title)

### Result Link (Redirect URL)
- **Format**: `//duckduckgo.com/l/?uddg=<URL_ENCODED>&rut=<TRACKING_HASH>`
- **uddg parameter**: Contains encoded destination URL
- **Extraction**: URL-decode the `uddg` parameter value

### Favicon/Result Icon
- **Selector**: `span.result__icon > img.result__icon__img`
- **Attributes**:
  - `src`: `//external-content.duckduckgo.com/ip3/<domain>.ico`
  - `width="16"`, `height="16"`
  - `alt=""`: Empty alt attribute
- **Purpose**: Site favicon/icon
- **Format**: Extracted from domain (e.g., `www.python.org` → `.../ip3/www.python.org.ico`)

### Display URL
- **Selector**: `a.result__url`
- **Attributes**:
  - `href`: Same redirect URL as title
- **Content**: Display URL (e.g., `www.python.org` or `www.python.org/path`)
- **Purpose**: Shows domain/path for user reference

### Result Snippet/Description
- **Selector**: `a.result__snippet`
- **Attributes**:
  - `href`: Same redirect URL as title/domain
- **Content**: Text with HTML `<b>` tags around query keywords
- **Format**: Description text, HTML entity encoded special characters
- **Note**: Clickable link element (not just text span)

---

## 6) CSS Classes Reference

| Class | Element | Purpose |
|-------|---------|---------|
| `body--html` | `<body>` | HTML version body styling |
| `header` | `<div>` | Header container |
| `header--html` | `<div>` | HTML-specific header styling |
| `cw` | `<div>` | Container width/layout class |
| `header__logo-wrap` | `<a>` | Logo link wrapper |
| `header__form` | `<form>` | Header search form |
| `search` | `<div>` | Search input wrapper |
| `search--header` | `<div>` | Header search styling |
| `search__input` | `<input>` | Query input field |
| `search__button` | `<input>` | Submit button |
| `search__button--html` | `<input>` | HTML version submit button |
| `frm__select` | `<div>` | Select container (region filter) |
| `frm__select--last` | `<div>` | Last select container (time filter) |
| `serp__results` | `<div>` | Results page container |
| `result` | `<div>` | Single result wrapper |
| `results_links` | `<div>` | Link-based result |
| `results_links_deep` | `<div>` | Deep linking result |
| `web-result` | `<div>` | Web search result type |
| `links_main` | `<div>` | Main result content |
| `links_deep` | `<div>` | Deep linking content |
| `result__body` | `<div>` | Result content body |
| `result__title` | `<h2>` | Result title heading |
| `result__a` | `<a>` | Title link styling |
| `result__extras` | `<div>` | Extra information container |
| `result__extras__url` | `<div>` | URL extras section |
| `result__icon` | `<span>` | Favicon container |
| `result__icon__img` | `<img>` | Favicon image |
| `result__url` | `<a>` | Display URL link |
| `result__snippet` | `<a>` | Description/snippet text link |
| `clear` | `<div>` | CSS clear float |

---

## 7) Key Differences from Lite Version

| Feature | Lite (`lite.duckduckgo.com`) | HTML (`html.duckduckgo.com`) |
|---------|---|---|
| Page structure | Tables for layout | Div-based with semantic classes |
| Result container | `<table border="0">` | `<div class="serp__results">` |
| Result row | Series of `<tr>` rows | Single `<div class="result">` |
| Title element | `<a class="result-link">` | `<h2><a class="result__a">` |
| URL display | `<span class="link-text">` | `<a class="result__url">` |
| Snippet | `<td class="result-snippet">` | `<a class="result__snippet">` |
| Favicon | No favicon | `<img class="result__icon__img">` |
| CSS approach | Minimal CSS | Modern CSS classes with BEM naming |
| DOCTYPE | HTML 4.01 Transitional | XHTML 1.0 Transitional |

---

## 8) Extraction Patterns for Scrapers

### Pattern: Extract All Results

```
1. Select all <div class="result results_links results_links_deep web-result">
2. For each result div:
   a. Title: <h2 class="result__title"> > <a> > text
   b. Title href: <h2 class="result__title"> > <a> > href (redirect URL)
   c. Display URL: <a class="result__url"> > text
   d. Favicon src: <img class="result__icon__img"> > src
   e. Snippet: <a class="result__snippet"> > text/inner HTML
```

### Pattern: CSS Selector Approach (DOM Query)

```javascript
// Get all results
const results = document.querySelectorAll('div.result.results_links.web-result');

results.forEach((result) => {
  const title = result.querySelector('h2.result__title a.result__a').textContent;
  const href = result.querySelector('h2.result__title a.result__a').href;
  const displayUrl = result.querySelector('a.result__url').textContent;
  const faviconSrc = result.querySelector('img.result__icon__img')?.src;
  const snippet = result.querySelector('a.result__snippet').textContent;
});
```

### Pattern: Extract Real URL from Redirect

```
1. Get href: //duckduckgo.com/l/?uddg=<ENCODED>&rut=<HASH>
2. Extract "uddg" query parameter value
3. URL-decode the value
4. Result: actual destination URL
```

### Pattern: Handle Favicon URLs

```
Original format: //external-content.duckduckgo.com/ip3/<DOMAIN>.ico
Example: //external-content.duckduckgo.com/ip3/www.python.org.ico

Extract domain from result URL to construct favicon URL:
- If URL is https://example.com/path
- Favicon: //external-content.duckduckgo.com/ip3/example.com.ico
```

---

## 9) HTML Entities and Encoding

### Common HTML Entities in Results
- `&#x27;` → `'` (apostrophe)
- `&amp;` → `&` (ampersand)
- `&lt;` → `<` (less than)
- `&gt;` → `>` (greater than)
- `&quot;` → `"` (quote)

### Query Keyword Highlighting
- Bold tags: `<b>python</b>` highlights matching keywords
- May appear in snippet text
- Decode HTML entities before processing

### URL Encoding in Links
- `uddg` parameter is URL-encoded
- Must be URL-decoded to get readable URL
- Actual destination URL is obtained by URL-decoding `uddg`

---

## 10) Meta Information

### Page Title
- **Selector**: `<title>`
- **Format**: `<query> at DuckDuckGo`
- **Example**: `python at DuckDuckGo`

### Meta Tags
- **Robots**: `<meta name="robots" content="noindex, nofollow" />`
- **Referrer**: `<meta name="referrer" content="origin" />`
- **Viewport**: `<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=3.0, user-scalable=1" />`
- **Charset**: UTF-8

### OpenSearch Discovery
- **Selector**: `<link rel="search" ... href="//duckduckgo.com/opensearch_html_v2.xml">`
- **Type**: `application/opensearchdescription+xml`
- **Title**: `DuckDuckGo (HTML)`

### Stylesheets
- **Main CSS**: Hashed filename, e.g., `//duckduckgo.com/dist/h.238c80a7d9b754cfcdd5.css`
- **Media**: `handheld, all`

---

## 11) Empty States and Edge Cases

### No Results
- Results container is empty (no result `<div>` children)
- Search form and header still displayed

### Single Result
- Same structure as multiple results
- Just contains one result `<div>`

### Pagination
- Present as a POST form in `<div class="nav-link">` near the end of `#links.results`
- Observed hidden fields: `q`, `s`, `nextParams`, `v`, `o`, `dc`, `api`, `vqd`, optional `kl`

### Bot Challenge
- Different page structure if bot detection triggered
- Look for presence of challenge modal or different response structure

---

## 12) Structural Comparison: Lite vs HTML

| Feature | Lite (`lite.duckduckgo.com`) | HTML (`html.duckduckgo.com`) |
|---|---|---|
| Result container | Table rows (`<tr>`) | Result blocks (`<div class="result...">`) |
| Title selector | `a.result-link` | `h2.result__title > a.result__a` |
| Snippet selector | `td.result-snippet` | `a.result__snippet` |
| Display URL selector | `span.link-text` | `a.result__url` |
| Favicon node | Not present | `img.result__icon__img` |
| Pagination container | `form.next_form` | `div.nav-link > form` |
