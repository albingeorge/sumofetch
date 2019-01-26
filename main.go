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

	r := sumocred.Search(paymentID + " (GATEWAY_SOAP_REQUEST or GATEWAY_REQUEST_TIMEOUT or PAYMENT_CALLBACK_REQUEST or GATEWAY_PAYMENT_REQUEST)  | json auto nodrop")

	for _, c := range r {
		fmt.Println(c)
	}
}
