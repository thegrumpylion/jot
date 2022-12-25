package lsp

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/thegrumpylion/jot/internal/adoc"
)

// Document represents a document in a workspace.
type Document struct {
	Path    string
	Content string
	doc     *types.Document
}

// NewDocument creates a new document.
func NewDocument(path, content string) (*Document, error) {
	d := &Document{
		Path:    path,
		Content: content,
	}
	doc, err := adoc.Parse(strings.NewReader(content), configuration.NewConfiguration())
	if err != nil {
		return nil, err
	}

	d.doc = doc

	return d, nil
}

// Update updates the document content.
func (d *Document) Update(content string) error {
	d.Content = content
	doc, err := adoc.Parse(strings.NewReader(content), configuration.NewConfiguration())
	if err != nil {
		return err
	}
	d.doc = doc
	return nil
}

// GetElement returns the element for the given position.
func (d *Document) GetElement(line, charachter uint32) interface{} {
	fmt.Println("GET ELEMENT", line, charachter)
	for k, v := range d.doc.ElementReferences {
		fmt.Println(k, v)
	}
	el, _ := getElement(d.doc.Elements, line, charachter)
	return el
}

func getElement(elements []interface{}, line, charachter uint32) (interface{}, bool) {
	for _, e := range elements {
		switch v := e.(type) {
		case *types.DocumentHeader:
			el, found := getElement(v.Title, line, charachter)
			if found {
				return el, found
			}
			el, found = getElement(v.Elements, line, charachter)
			if found {
				return el, found
			}
		case *types.Paragraph:
			fmt.Println("PARAGRAPH", v)
		case *types.ListElement:
			fmt.Println("UL", v)
		case *types.StringElement:
		default:
			fmt.Printf("DEFAULT %T\n", v)
		}
	}

	return nil, false
}
