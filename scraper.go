package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	Index       int
	Title       string
	URL         string
	Description string
}

func extractResultURL(href string) string {
	u, err := url.Parse("https:" + href)
	if err != nil {
		return href
	}

	uddg := u.Query().Get("uddg")
	if uddg == "" {
		return href
	}

	decodedURL, err := url.QueryUnescape(uddg)
	if err != nil {
		return uddg
	}
	
	return decodedURL
}

func ScrapeDuckDuckGo(query string, page int, limit int) ([]SearchResult, error) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/147.0.0.0 Safari/537.36"

	// Generate the search URL
	escapedQuery := url.QueryEscape(query)
	requestURL := fmt.Sprintf("https://lite.duckduckgo.com/lite/?q=%s", escapedQuery)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	index := 1

	doc.Find("a.result-link").Each(func(i int, s *goquery.Selection) {
		if len(results) >= limit {
			return
		}

		title := strings.TrimSpace(s.Text())
		href, _ := s.Attr("href")
		actualURL := extractResultURL(href)
		
		description := ""
		parent := s.Parent().Parent()
		parent.NextFiltered("tr").Find("td.result-snippet").Each(func(i int, s *goquery.Selection) {
			description = strings.TrimSpace(s.Text())
		})

		result := SearchResult{
			Index: index,
			Title: title,
			URL: actualURL,
			Description: description,
		}
		results = append(results, result)
		index++
	})

	return results, nil
}
