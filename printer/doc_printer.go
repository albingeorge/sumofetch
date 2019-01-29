package printer

import (
	"baliance.com/gooxml/document"
	"fmt"
	"os"
	"strings"
)
import "github.com/albingeorge/sumofetch/formatter"

func PrintDoc(input []formatter.FormattedContent) {

	testPath, testName := "", ""
	if len(os.Args) > 2 {
		testPath = os.Args[2]
		splittedString := strings.Split(testPath, "/")

		testName = splittedString[len(splittedString)-1]

	} else {
		fmt.Println("Test name not provided. Exiting without printing to document")
		return
	}

	doc := document.New()

	addText(doc, testName, "", true)

	addText(doc, "Screenshots", "", false)

	for _, content := range input {
		addText(doc, content.Header, "", true)
		addText(doc, "Date Time: "+content.DateTime.Format("01/02/2006 15:04:05"), "", false)
		addText(doc, content.Content, "", false)

		addText(doc, "", "", false)
	}

	err := doc.SaveToFile(testPath + ".docx")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error printing to file: " + testPath + ".docx")
	} else {
		fmt.Println("Printed to file: " + testPath + ".docx")
	}

}

func addText(doc *document.Document, content string, style string, bold bool) {

	para := doc.AddParagraph()

	if style != "" {
		para.SetStyle(style)
	}

	run := para.AddRun()
	run.Properties().SetBold(bold)

	str := strings.Split(content, "\n")

	for _, c := range str {
		run.AddText(c)
		run.AddBreak()
	}

}
