package jot

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetDir() (string, error) {
	// get current path
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// find .jot directory from current path and back until root
	for {
		if _, err := os.Stat(path + "/.jot"); err == nil {
			break
		}
		if path == "/" {
			return "", fmt.Errorf("no .jot directory found in the current path or any parent")
		}
		path = filepath.Dir(path)
	}
	return path, nil
}
