package sumo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sumo Sumocreds) createSearchQueryID(query string) string {
	url := sumo.BaseUrl

	searchInputs := generateSearchQueryInputs(query)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(searchInputs))

	sumo.setHeaders(req)
	response := sendRequest(req)

	id := fetchQueryIDFromResponse(response)

	return id
}

// Generate the body for search query for the past 2 hours
// Will accept the time range if required in the future
func generateSearchQueryInputs(query string) []byte {

	//current := time.Now()
	//nowString := current.Format(time.RFC3339)
	//
	//then := current.Add(time.Duration(-6) * time.Hour)
	//twoHoursBack := then.Format(time.RFC3339)
	//
	//search := map[string]interface{}{
	//	"query":    query,
	//	"from":     twoHoursBack,
	//	"to":       nowString,
	//	"timeZone": "Asia/Kolkata",
	//}

	search := map[string]string{
		"query":    query,
		"from":     "2019-01-04T00:00",
		"to":       "2019-01-04T23:59",
		"timeZone": "Asia/Kolkata",
	}
	jsonString, _ := json.Marshal(search)

	return jsonString
}

func fetchQueryIDFromResponse(res []byte) string {

	type createSearchQueryResult struct {
		ID string
	}

	var search createSearchQueryResult
	err := json.Unmarshal(res, &search)

	if err != nil {
		fmt.Println("Fetching search id from Json response failed!")
	}

	return search.ID
}
