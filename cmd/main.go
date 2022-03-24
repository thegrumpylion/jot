package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "jot",
	RunE: rootCmdRunE,
}

func rootCmdRunE(cmd *cobra.Command, args []string) error {
	fmt.Println("jot")
	return nil
}

func main() {
	if err := rootCmd.ExecuteContext(context.TODO()); err != nil {
		fmt.Println("error", err)
	}
}
