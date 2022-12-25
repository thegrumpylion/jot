package adoc

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func ParseFile(fn string) (*types.Document, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &configuration.Configuration{}

	stat, err := os.Stat(fn)
	if err != nil {
		return nil, err
	}
	config.LastUpdated = stat.ModTime()

	return Parse(file, config)
}

func Parse(r io.Reader, config *configuration.Configuration) (*types.Document, error) {
	p, err := parser.Preprocess(r, config)
	if err != nil {
		return nil, err
	}

	return parser.ParseDocument(strings.NewReader(p), config)
}

func Fragments(r io.Reader, config *configuration.Configuration) (interface{}, error) {

	done := make(chan interface{})
	defer close(done)

	p, err := parser.Preprocess(r, config)
	if err != nil {
		return nil, err
	}

	frags := parser.ParseDocumentFragments(parser.NewParseContext(config), strings.NewReader(p), done)

	for frag := range frags {
		fmt.Println(frag)
		if frag.Error != nil {
			return nil, frag.Error
		}
	}
	return nil, nil
}
