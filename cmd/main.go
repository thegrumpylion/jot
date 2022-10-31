package main

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "jot",
	RunE: rootCmdRunE,
	Args: cobra.ExactArgs(1),
}

func rootCmdRunE(cmd *cobra.Command, args []string) error {

	err := filepath.WalkDir(args[0], func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Type().IsRegular() && filepath.Ext(path) == ".adoc" {
			fmt.Println(path)
		}
		return nil
	})

	return err
}

func main() {
	if err := rootCmd.ExecuteContext(context.TODO()); err != nil {
		fmt.Println("error", err)
	}
}
