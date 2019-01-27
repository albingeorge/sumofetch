package main

import (
	"bitbucket.org/albingeorgee/sumofetch/config"
	"bitbucket.org/albingeorgee/sumofetch/formatter"
	"bitbucket.org/albingeorgee/sumofetch/printer"
	"bitbucket.org/albingeorgee/sumofetch/sumo"
	"fmt"
	"os"
)

func main() {
	conf, err := config.GetConfigs()

	if err != nil {
		fmt.Println(err)
	}

	sumocred := sumo.New(conf)

	paymentID := os.Args[1]

	results := sumocred.Search(paymentID + " (GATEWAY_SOAP_REQUEST or GATEWAY_REQUEST_TIMEOUT or PAYMENT_CALLBACK_REQUEST or GATEWAY_PAYMENT_REQUEST)  | json auto nodrop")

	formattedResult := formatter.Format(results)

	printer.PrintStdout(formattedResult)
}
