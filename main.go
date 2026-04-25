package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		limit         int
		minimalOutput bool
		region        string
		safeSearch    int
		jsonOutput    bool
	)

	flag.IntVar(&limit, "limit", 10, "Limit the number of results")
	flag.IntVar(&limit, "l", 10, "Alias for --limit")
	flag.BoolVar(&minimalOutput, "minimal-output", false, "Show only title and URL (omit description)")
	flag.BoolVar(&minimalOutput, "m", false, "Alias for --minimal-output")
	flag.StringVar(&region, "region", "wt-wt", "Set search region (for example: wt-wt, us-en, kr-kr)")
	flag.StringVar(&region, "kl", "wt-wt", "Alias for --region")
	flag.IntVar(&safeSearch, "safe-search", -1, "Set safe search: 1=on, -1=moderate, -2=off")
	flag.IntVar(&safeSearch, "kp", -1, "Alias for --safe-search")
	flag.BoolVar(&jsonOutput, "json", false, "Output results as JSON")
	flag.BoolVar(&jsonOutput, "j", false, "Alias for --json")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ddg-cli: search DuckDuckGo from your terminal\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n  %s [options] <query>\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintf(os.Stderr, "  %s github\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -limit 5 golang cli\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -m -region us-en -safe-search 1 github actions\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -json open source licenses\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Notes:")
		fmt.Fprintln(os.Stderr, "  - Put the query after options.")
		fmt.Fprintln(os.Stderr, "  - All remaining arguments are joined as a single query string.")
	}

	flag.Parse()

	if limit < 1 {
		log.Fatal("limit must be a positive integer")
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	query := strings.TrimSpace(strings.Join(args, " "))

	results, err := ScrapeDuckDuckGo(query, 1, limit, region, safeSearch)
	if err != nil {
		log.Fatal(err)
	}

	if len(results) == 0 {
		fmt.Println("No results found.")
		return
	}

	if jsonOutput {
		jsonData, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			log.Fatal("Error encoding results to JSON:", err)
		}
		fmt.Println(string(jsonData))
		return
	} else {
		for _, result := range results {
			fmt.Printf("%d.\t%s\n", result.Index, result.Title)
			fmt.Printf("\tURL: %s\n", result.URL)
			if !minimalOutput && result.Description != "" {
				fmt.Printf("\tDescription: %s\n", result.Description)
			}
			fmt.Println()
		}
	}
}
