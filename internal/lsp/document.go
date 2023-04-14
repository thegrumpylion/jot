package lsp

import (
	"bufio"
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
	lines   []string
	offsets []uint32
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
	d.calculateOffsets()

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
	d.calculateOffsets()

	return nil
}

// GetElement returns the element for the given position.
func (d *Document) GetElement(line, charachter uint32) interface{} {
	inf, _ := adoc.Fragments(strings.NewReader(d.Content), configuration.NewConfiguration())
	for _, v := range inf {
		fmt.Println("DocumentFragment", v.Position)
		adoc.PrintElements(v.Elements, "")
	}
	fmt.Println("offset", d.getOffset(line, charachter))
	return inf
}

// calculateOffset calculates the offset for the given line and charachter.
func (d *Document) calculateOffsets() {

	fmt.Println("calculateOffsets")

	// create the scanner
	scanner := bufio.NewScanner(strings.NewReader(d.Content))

	// initialize the lines and offsets
	d.lines = []string{}
	d.offsets = []uint32{0}

	// read string line by line
	for scanner.Scan() {
		line := scanner.Text()
		d.lines = append(d.lines, line)

		lastOffset := d.offsets[len(d.offsets)-1]
		d.offsets = append(d.offsets, lastOffset+uint32(len(line))+1)
	}
}

// getOffset returns the offset for the given line and charachter.
func (d *Document) getOffset(line, charachter uint32) uint32 {
	return d.offsets[line] + charachter
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
		case *types.List:

		default:
			fmt.Printf("DEFAULT %T\n", v)
		}
	}

	return nil, false
}
