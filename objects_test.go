package jot

import "testing"

func TestFileDoc(t *testing.T) {
	_, err := NewFileDocument("README.adoc", &NullStateMachine{})
	if err != nil {
		t.Fatal(err)
	}
}
