package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/thegrumpylion/jot/internal/lsp"
)

type lsCommand struct {
	*cobra.Command
	address string
}

var lsCmd = &lsCommand{
	Command: &cobra.Command{
		Use:   "ls",
		Short: "Language Server",
		RunE:  lsCmdRunE,
	},
}

func init() {
	lsCmd.Flags().StringVarP(&lsCmd.address, "address", "a", "0.0.0.0:56857", "Address to serve the site on")
	rootCmd.AddCommand(lsCmd.Command)
}

func lsCmdRunE(cmd *cobra.Command, args []string) error {
	server := lsp.NewServer("0.0.0")
	fmt.Println("lsp serving")
	return server.Serve(context.Background(), time.Minute*3600)
}
