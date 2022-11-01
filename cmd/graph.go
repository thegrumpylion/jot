package main

import "github.com/spf13/cobra"

type graphCommand struct {
	*cobra.Command
}

var graphCmd = &graphCommand{
	Command: &cobra.Command{
		Use:  "graph",
		RunE: graphCmdRunE,
	},
}

func init() {
	rootCmd.AddCommand(graphCmd.Command)
}

func graphCmdRunE(cmd *cobra.Command, args []string) error {
	return nil
}
