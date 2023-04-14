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

func Fragments(r io.Reader, config *configuration.Configuration) ([]types.DocumentFragment, error) {

	done := make(chan interface{})
	defer close(done)

	frags := parser.ParseDocumentFragments(parser.NewParseContext(config), r, done)
	rfrags := parser.RefineFragments(parser.NewParseContext(config), r, done, frags)
	substs := parser.ApplySubstitutions(parser.NewParseContext(config), done, rfrags)
	footnotes := parser.CollectFootnotes(types.NewFootnotes(), done, substs)
	lists := parser.ArrangeLists(done, footnotes)

	out := []types.DocumentFragment{}

	for frag := range lists {
		if frag.Error != nil {
			return nil, frag.Error
		}
		out = append(out, frag)
	}
	return out, nil
}
