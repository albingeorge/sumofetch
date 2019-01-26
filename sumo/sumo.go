package sumo

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/albingeorgee/sumofetch/config"
)

type sumo interface {
	createSearchQueryID(string) string

	Search(string) map[string]string
}

// Sumocreds - Credentials structure
type Sumocreds struct {
	BaseUrl   string
	AccessID  string
	AccessKey string
}

func (sumo Sumocreds) setHeaders(req *http.Request) {
	req.SetBasicAuth(sumo.AccessID, sumo.AccessKey)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
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
func (sumo Sumocreds) Search(query string) []ResponseFormat {

	queryID := sumo.createSearchQueryID(query)

	// exportID := sumo.createExportID(queryID)

	fmt.Println("queryID", queryID)

	ret := sumo.fetchQueryResults(queryID)

	return ret
}

// New -  Initiates a new Sumocreds object
func New(conf config.Config) Sumocreds {
	baseUrl := "https://api.eu.sumologic.com/api/v1/search/jobs"

	creds := Sumocreds{BaseUrl: baseUrl, AccessID: conf.AccessID, AccessKey: conf.AccessKey}
	return creds
}
