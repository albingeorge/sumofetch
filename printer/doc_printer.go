package printer

import (
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	//addText(doc, "Screenshots", "", false)

	addScreenshots(doc)

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

func addScreenshots(doc *document.Document) {

	if len(os.Args) > 3 {

		root := os.Args[3]

		imageReferences := []common.ImageRef{}

		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

			// If it's a directory, skip it
			if info.IsDir() {
				return nil
			}

			img, err := common.ImageFromFile(path)
			if err != nil {
				fmt.Println("Unable to create image from ", path, ". Error: ", err)
				return nil
			}

			iref, err := doc.AddImage(img)

			if err != nil {
				fmt.Println("Error adding image: ", err)
				return nil
			}

			imageReferences = append(imageReferences, iref)

			return nil
		})

		if err != nil {
			return
		}

		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText("Screenshots")
		run.AddBreak()

		for _, iref := range imageReferences {
			//_, err := run.AddDrawingAnchored(iref)
			inl, err := run.AddDrawingInline(iref)

			iref.RelID()

			if err != nil {
				log.Fatalf("unable to add inline image: %s", err)
			}

			inl.SetSize(5.5*measurement.Inch, 3.5*measurement.Inch)

			run.AddBreak()
			run.AddBreak()
		}

	} else {
		fmt.Println("Skipping adding screenshots, since no path provided")
	}
}
