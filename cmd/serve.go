package main

import "github.com/spf13/cobra"

type serveCommand struct {
	*cobra.Command
	address string
}

var serveCmd = &serveCommand{
	Command: &cobra.Command{
		Use:   "serve",
		Short: "Serve the site",
		RunE:  serveCmdRunE,
	},
}

func init() {
	serveCmd.Flags().StringVarP(&serveCmd.address, "address", "a", ":8080", "Address to serve the site on")
	rootCmd.AddCommand(serveCmd.Command)
}

func serveCmdRunE(cmd *cobra.Command, args []string) error {
	return nil
}
