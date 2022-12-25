package lsp

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// Document represents a document in a workspace.
type Document struct {
	Path    string
	Content string
	doc     *types.Document
}
