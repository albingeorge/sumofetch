package printer

import "baliance.com/gooxml/document"
import "github.com/albingeorge/sumofetch/formatter"

func PrintDoc(input []formatter.FormattedContent) {
	doc := document.New()

	doc.AddParagraph()
}
