package adoc

import (
	"fmt"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func PrintElements(elements []interface{}, indent string) {
	for _, e := range elements {
		switch v := e.(type) {
		case *types.DocumentHeader:
			fmt.Printf("%sDocumentHeader\n", indent)
			PrintElements(v.Title, indent+"  ")
			PrintElements(v.Elements, indent+"  ")
		case *types.Paragraph:
			fmt.Printf("%sParagraph\n", indent)
			PrintElements(v.Elements, indent+"  ")
		case *types.ListElement:
			fmt.Printf("%sListElement\n", indent)
		case *types.StringElement:
			fmt.Printf("%sStringElement: %s\n", indent, v.Content)
		case *types.List:
			fmt.Printf("%sList\n", indent)
			indent += "  "
			for _, e := range v.Elements {
				switch v := e.(type) {
				case *types.UnorderedListElement:
					fmt.Printf("%sUnorderedListElement %s\n", indent, v.BulletStyle)
					PrintElements(v.Elements, indent+"  ")
				case *types.OrderedListElement:
					fmt.Printf("%sOrderedListElement %s\n", indent, v.Style)
					PrintElements(v.Elements, indent+"  ")
				}
			}
		case *types.DelimitedBlock:
			fmt.Printf("%sDelimitedBlock %s\n", indent, v.Kind)
			PrintElements(v.Elements, indent+"  ")
		default:
			fmt.Printf("%sDEFAULT %T\n", indent, v)
		}
	}
}
