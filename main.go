package main

import (
	"fmt"
	"os"

	"bitbucket.org/albingeorgee/sumofetch/config"
	"bitbucket.org/albingeorgee/sumofetch/sumo"
)

func main() {
	conf, err := config.GetConfigs()

	if err != nil {
		fmt.Println(err)
	}

	sumocred := sumo.New(conf)

	paymentID := os.Args[1]

	result := sumocred.Search(paymentID + " (gateway_soap_request or GATEWAY_REQUEST_TIMEOUT or PAYMENT_CALLBACK_REQUEST)  | json auto nodrop")

	fmt.Println(result)
}
