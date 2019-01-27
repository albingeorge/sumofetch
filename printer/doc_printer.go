package printer

import (
	"baliance.com/gooxml/document"
	"fmt"
	"os"
)
import "github.com/albingeorge/sumofetch/formatter"

func PrintDoc(input []formatter.FormattedContent) {

	testName := ""
	if len(os.Args) > 2 {
		testName = os.Args[2]
	} else {
		fmt.Println("Test name not provided. Exiting without printing to document")
		return
	}

	doc := document.New()

	addText(doc, testName, "Title", false)

	addText(doc, "Screenshots", "Heading1", false)

	addText(doc, "Logs", "Heading1", false)

	for _, content := range input {
		addText(doc, "\n"+content.Header, "Heading2", true)
		addText(doc, "\nDate Time: "+content.DateTime.Format("01/02/2006 15:04:05"), "", false)
		addText(doc, "\n"+content.Content, "", false)

		addText(doc, "", "", false)
	}

	doc.SaveToFile(testName + ".docx")
}

func addText(doc *document.Document, content string, style string, bold bool) {

	para := doc.AddParagraph()

	if style != "" {
		para.SetStyle(style)
	}

	run := para.AddRun()
	run.Properties().SetBold(bold)
	run.AddText(content)
}
