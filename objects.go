package jot

import (
	"fmt"
	"os"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

type Document interface {
	Document() types.Document
	Plugins() map[string]Plugin
}

type Plugin interface {
}

type FileDocument struct {
	path string
	doc  *types.Document
}

func NewFileDocument(path string, sm StateMachine) (*FileDocument, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &configuration.Configuration{}
	config.Attributes.Set("kostis", "whatever")

	// p, err := parser.Preprocess(file, config)
	// if err != nil {
	// 	return nil, err
	// }

	fmt.Println("Attributes")
	fmt.Println(config.Attributes)

	doc, err := parser.ParseDocument(file, config)
	if err != nil {
		return nil, err
	}

	fmt.Println("Attributes again")
	fmt.Println(config.Attributes, doc.Header().Attributes)

	for _, v := range doc.Elements {
		switch e := v.(type) {
		case *types.DocumentHeader:
			fmt.Println("DocumentHeader")
			fmt.Println(e.Title...)
			fmt.Println(e.Attributes, len(e.Attributes))
		case *types.Paragraph:
			fmt.Println("Paragraph")
			fmt.Println(e.Elements...)
		default:
			fmt.Printf("def %T\n", e)
		}
	}

	return &FileDocument{
		path: path,
		doc:  doc,
	}, nil

}
