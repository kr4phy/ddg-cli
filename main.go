package main

import (
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
	)

	flag.IntVar(&limit, "limit", 10, "Limit the number of results")
	flag.BoolVar(&minimalOutput, "minimal-output", false, "Only display the title and URL of each result, omitting descriptions")
	flag.BoolVar(&minimalOutput, "m", false, "Alias for --minimal-output")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <query>\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
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

	results, err := ScrapeDuckDuckGo(query, 1, limit)
	if err != nil {
		log.Fatal(err)
	}

	if len(results) == 0 {
		fmt.Println("No results found.")
		return
	}

	for _, result := range results {
		fmt.Printf("%d.\t%s\n", result.Index, result.Title)
		fmt.Printf("\tURL: %s\n", result.URL)
		if !minimalOutput && result.Description != "" {
			fmt.Printf("\tDescription: %s\n", result.Description)
		}
		fmt.Println()
	}
}
