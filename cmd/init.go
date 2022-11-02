package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
)

type initCommand struct {
	*cobra.Command
}

var initCmd = &initCommand{
	Command: &cobra.Command{
		Use:  "init",
		RunE: initCmdRunE,
	},
}

func init() {
	rootCmd.AddCommand(initCmd.Command)
}

func initCmdRunE(cmd *cobra.Command, args []string) error {

	// git init if not exists
	repo, err := git.PlainInit(".", false)
	if err != nil && err != git.ErrRepositoryAlreadyExists {
		return err
	}

	// get git config scoped globally
	config, err := repo.ConfigScoped(config.GlobalScope)
	if err != nil {
		return err
	}

	// check if config has user.name and user.email
	if config.User.Name == "" || config.User.Email == "" {
		return fmt.Errorf("user.name and user.email must be set")
	}

	// check if .jot directory exists, if not, create it
	if _, err := os.Stat(".jot"); os.IsNotExist(err) {
		err = os.Mkdir(".jot", 0755)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	// check if .gitignore exists, if not, create it and add .jot/build to it
	if _, err := os.Stat(".gitignore"); os.IsNotExist(err) {
		f, err := os.Create(".gitignore")
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(".jot/build")
		if err != nil {
			return err
		}
	} else {
		return err
	}

	// get the worktree
	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	// add .gitignore to the index
	_, err = w.Add(".gitignore")
	if err != nil {
		return err
	}

	// commit the changes
	_, err = w.Commit("add .gitignore", &git.CommitOptions{
		Author: &object.Signature{
			Name:  config.User.Name,
			Email: config.User.Email,
			When:  time.Now(),
		},
	})

	return nil
}
