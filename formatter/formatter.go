package formatter

import (
	"bitbucket.org/albingeorgee/sumofetch/sumo"
	"fmt"
)

func Format(format []sumo.ResponseFormat) {
	for _, response := range format {
		formatResponse(response)
	}
}

func formatResponse(response sumo.ResponseFormat) {
	fmt.Println(response)
}
