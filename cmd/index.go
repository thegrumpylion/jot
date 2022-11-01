package main

import "github.com/spf13/cobra"

type indexCommand struct {
	*cobra.Command
}

var indexCmd = &indexCommand{
	Command: &cobra.Command{
		Use:  "index",
		RunE: indexCmdRunE,
	},
}

func init() {
	rootCmd.AddCommand(indexCmd.Command)
}

func indexCmdRunE(cmd *cobra.Command, args []string) error {
	return nil
}
