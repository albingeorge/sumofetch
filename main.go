package main

import (
	"fmt"
	"github.com/albingeorge/sumofetch/config"
	"github.com/albingeorge/sumofetch/formatter"
	"github.com/albingeorge/sumofetch/printer"
	"github.com/albingeorge/sumofetch/sumo"
	"os"
)

func main() {
	conf, err := config.GetConfigs()

	if err != nil {
		fmt.Println(err)
	}

	sumocred := sumo.New(conf)

	paymentID := os.Args[1]

	results := sumocred.Search(paymentID + " (GATEWAY_SOAP_REQUEST or GATEWAY_REQUEST_TIMEOUT or PAYMENT_CALLBACK_REQUEST or GATEWAY_PAYMENT_REQUEST or GATEWAY_CHECKSUM_VERIFY_FAILED or GATEWAY_ERROR_DATA_MISMATCH)  | json auto nodrop")
	//formatter.Format(results)
	formattedResult := formatter.Format(results)

	printer.PrintStdout(formattedResult)
}
