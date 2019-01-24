package sumo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/albingeorgee/sumofetch/config"
)

type sumo interface {
	createSearchQueryID(string) string

	sendAPIRequtest(string) string
	Search(string) map[string]string
}

type sumocreds struct {
	APISession    string
	SumoServiceID string
}

func (sumo sumocreds) createSearchQueryID(query string) string {
	url := "https://service.eu.sumologic.com/json/v2/searchquery/create"

	searchInputs := generateSearchQueryInputs("abc")

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(searchInputs))

	sumo.setHeaders(req)
	response := sendRequest(req)

	id := fetchQueryIDFromResponse(response)

	return id
}

func (sumo sumocreds) setHeaders(req *http.Request) {
	req.Header.Set("ApiSession", sumo.APISession)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "SUMOSERVICEID="+sumo.SumoServiceID)
}

func sendRequest(req *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func (sumo sumocreds) Search(query string) []map[string]string {

	queryID := sumo.createSearchQueryID(query)

	fmt.Println("queryID", queryID)

	m := map[string]string{"a": "b"}
	ret := []map[string]string{m}
	return ret
}

// func (sumo sumocreds) sendAPIRequtest(url string) string {

// 	return ""
// }

func New(conf config.Config) sumocreds {
	creds := sumocreds{APISession: conf.APISession, SumoServiceID: conf.SumoServiceID}
	return creds
}
