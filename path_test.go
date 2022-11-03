package jot_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestPathRegression(t *testing.T) {
	// get current path
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// find .jot directory from current path and back until root
	for {
		if _, err := os.Stat(path + "/.jot"); err == nil {
			break
		}
		if path == "/" {
			t.Fatal("No .jot directory found")
		}
		path = filepath.Dir(path)
	}
	fmt.Println(path)
}
