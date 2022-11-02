package main

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
)

type rootCommand struct {
	*cobra.Command
}

var rootCmd = &rootCommand{
	Command: &cobra.Command{
		Use:  "jot",
		RunE: rootCmdRunE,
	},
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
