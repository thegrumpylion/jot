package adoc

import (
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
