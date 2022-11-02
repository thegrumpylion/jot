package main

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/blugelabs/bluge"
	"github.com/spf13/cobra"
)

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

	// bluge default config
	config := bluge.DefaultConfig(".jot/index")

	// open writer
	writer, err := bluge.OpenWriter(config)
	if err != nil {
		return err
	}

	// close writer
	defer writer.Close()

	// walk diractory for .adoc files and index them
	if err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".adoc" {
			// read file
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			// create new doc
			doc := bluge.NewDocument(path)
			// add content
			doc.AddField(bluge.NewTextField("content", string(b)))
			// add to index
			if err := writer.Update(doc.ID(), doc); err != nil {
				return err
			}
		}
		return nil

	}); err != nil {
		return err
	}

	return nil
}
