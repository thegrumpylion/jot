package main

import (
	"context"
	"fmt"

	"github.com/blugelabs/bluge"
	"github.com/spf13/cobra"
)

type searchCommand struct {
	*cobra.Command
}

var searchCmd = &searchCommand{
	Command: &cobra.Command{
		Use:  "search",
		RunE: searchCmdRunE,
		Args: cobra.ExactArgs(1),
	},
}

func init() {
	rootCmd.AddCommand(searchCmd.Command)
}

func searchCmdRunE(cmd *cobra.Command, args []string) error {

	// bluge default config
	config := bluge.DefaultConfig(".jot/index")

	// set default search field to "content"
	config.DefaultSearchField = "content"

	// new match query on args[0]
	query := bluge.NewMatchQuery(args[0])

	// open index reader
	reader, err := bluge.OpenReader(config)
	if err != nil {
		return err
	}

	// close reader
	defer reader.Close()

	// search request all matches
	request := bluge.NewAllMatches(query)

	// search
	results, err := reader.Search(context.Background(), request)
	if err != nil {
		return err
	}

	// print results
	next, err := results.Next()
	for err == nil && next != nil {
		next.VisitStoredFields(func(field string, value []byte) bool {
			if field == "_id" {
				fmt.Println(string(value))
			}
			return true
		})
		if err != nil {
			return err
		}
		next, err = results.Next()
	}

	return nil
}
