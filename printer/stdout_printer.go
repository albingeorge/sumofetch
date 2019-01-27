package printer

import (
	"bitbucket.org/albingeorgee/sumofetch/formatter"
	"fmt"
)

func PrintStdout(input []formatter.FormattedContent) {
	for _, content := range input {
		fmt.Println(content.Header + "\n")

		fmt.Println("DateTime: " + content.DateTime.Format("01/02/2006 15:04:05") + "\n")

		fmt.Println(content.Content + "\n")

		fmt.Println(content.Footer + "\n\n")
	}
}
