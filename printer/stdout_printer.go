package printer

import (
	"fmt"
	"github.com/albingeorge/sumofetch/formatter"
)

func PrintStdout(input []formatter.FormattedContent) {
	for _, content := range input {
		fmt.Println("\033[1m\033[4m" + content.Header + "\033[00m" + "\n")

		fmt.Println("DateTime: " + content.DateTime.Format("01/02/2006 15:04:05") + "\n")

		fmt.Println(content.Content + "\n")
	}
}
