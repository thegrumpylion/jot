package main

import "github.com/spf13/cobra"

type searchCommand struct {
	*cobra.Command
}

var searchCmd = &searchCommand{
	Command: &cobra.Command{
		Use:  "search",
		RunE: searchCmdRunE,
	},
}

func init() {
	rootCmd.AddCommand(searchCmd.Command)
}

func searchCmdRunE(cmd *cobra.Command, args []string) error {
	return nil
}
