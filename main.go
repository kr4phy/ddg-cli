package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "ddg-cli",
		Usage:     "Search DuckDuckGo from the command line",
		UsageText: "ddg-cli [options] <query>",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "limit",
				Value: 10,
				Usage: "Limit the number of results",
			},
			&cli.BoolFlag{
				Name:    "minimal-output",
				Aliases: []string{"m"},
				Value:   false,
				Usage:   "Only display the title and URL of each result, omitting descriptions",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.NArg() == 0 {
				return cli.ShowRootCommandHelp(c)
			}

			query := strings.TrimSpace(c.Args().First())
			if query == "" {
				return cli.ShowRootCommandHelp(c)
			}

			limit := c.Int("limit")
			if limit < 1 {
				return fmt.Errorf("limit must be a positive integer")
			}
			results, err := ScrapeDuckDuckGo(query, 1, limit)
			if err != nil {
				return err
			}

			if len(results) == 0 {
				fmt.Println("No results found.")
				return nil
			}

			for _, result := range results {
				fmt.Printf("%d. %s\n", result.Index, result.Title)
				fmt.Printf("   URL: %s\n", result.URL)
				if !c.Bool("minimal-output") && result.Description != "" {
					fmt.Printf("   Description: %s\n", result.Description)
				}
				fmt.Println()
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
