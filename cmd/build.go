package main

import "github.com/spf13/cobra"

type buildCommand struct {
	*cobra.Command
}

var buildCmd = &buildCommand{
	Command: &cobra.Command{
		Use:  "build",
		RunE: buildCmdRunE,
	},
}

func init() {
	rootCmd.AddCommand(buildCmd.Command)
}

func buildCmdRunE(cmd *cobra.Command, args []string) error {
	return nil
}
