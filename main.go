package main

import (
	"fmt"

	"bitbucket.org/albingeorgee/sumofetch/config"
	"bitbucket.org/albingeorgee/sumofetch/sumo"
)

func main() {
	conf, err := config.GetConfigs()

	if err != nil {
		fmt.Println(err)
	}

	sumocred := sumo.New(conf)

	result := sumocred.Search("abv")

	fmt.Println(result)
}
