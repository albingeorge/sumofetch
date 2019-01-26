package sumo

import (
	"fmt"
	"net/http"
)

func (sumo Sumocreds) fetchQueryResults(queryID string)[]map[string]string {
	url := "https://service.eu.sumologic.com/json/v2/searchquery/"+queryID+"/messages/raw?offset=0"

	req, _ := http.NewRequest("GET", url, nil)
	sumo.setHeaders(req)

	//time.Sleep(2 * time.Second)
	response := sendRequest(req)

	fmt.Println("fetchQueryResults response")
	fmt.Println(string(response))

	m := map[string]string{"a": "b"}
	ret := []map[string]string{m}
	return ret
}
