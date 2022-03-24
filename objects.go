package jot

import "github.com/bytesparadise/libasciidoc/pkg/types"

type StateMachine interface {
	State() string
	Transition(state string, data interface{}) error
}

type Document interface {
	Document() types.Document
	StateMachine() StateMachine
}
