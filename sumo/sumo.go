package sumo

import (
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

// Sumocreds - Credentials structure
type Sumocreds struct {
	APISession    string
	SumoServiceID string
}

func (sumo Sumocreds) setHeaders(req *http.Request) {
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

// Search - Creates a search qyery, and fetches it's results
func (sumo Sumocreds) Search(query string) []map[string]string {

	queryID := sumo.createSearchQueryID(query)

	// exportID := sumo.createExportID(queryID)

	fmt.Println("queryID", queryID)

	m := map[string]string{"a": "b"}
	ret := []map[string]string{m}
	return ret
}

// func (sumo Sumocreds) sendAPIRequtest(url string) string {

// 	return ""
// }

// New -  Initiates a new Sumocreds object
func New(conf config.Config) Sumocreds {
	creds := Sumocreds{APISession: conf.APISession, SumoServiceID: conf.SumoServiceID}
	return creds
}
