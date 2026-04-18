package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gocolly/colly"
)

type SearchResult struct {
	Index       int
	Title       string
	URL         string
	Description string
}

func ScrapeDuckDuckGo(query string, page int, limit int) ([]SearchResult, error) {
	// Setup Colly
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
	c.SetRequestTimeout(10 * time.Second)

	// Generate the search URL
	escapedQuery := url.QueryEscape(query)
	requestURL := fmt.Sprintf("https://lite.duckduckgo.com/lite/?q=%s", escapedQuery)

	_ = page
	_ = limit

	return []SearchResult{}, nil
}
