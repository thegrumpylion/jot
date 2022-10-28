package jot

import "testing"

func TestFileDoc(t *testing.T) {
	_, err := NewFileDocument("README.adoc")
	if err != nil {
		t.Fatal(err)
	}
}
