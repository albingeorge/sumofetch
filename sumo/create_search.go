package sumo

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"
)

type createSearchQueryResult struct {
	QueryID string `xml:"searchQueryId"`
}

// Generate the body for search query for the past 2 hours
// Will accept the timerange if required in the future
func generateSearchQueryInputs(query string) []byte {

	current := time.Now()
	now := current.Unix()

	nowString := strconv.FormatInt(now, 10) + "000"

	then := current.Add(time.Duration(-2) * time.Hour)
	thenUnix := then.Unix()

	twoHoursBack := strconv.FormatInt(thenUnix, 10) + "000"

	search := map[string]interface{}{
		"queryString":     query,
		"startMillis":     twoHoursBack,
		"endMillis":       nowString,
		"byReceiptTime":   false,
		"timeZone":        "Asia/Kolkata",
		"isAllTime":       false,
		"fromSavedSearch": false,
	}
	jsonString, _ := json.Marshal(search)

	return jsonString
}

func fetchQueryIDFromResponse(res []byte) string {
	var search createSearchQueryResult
	xml.Unmarshal(res, &search)

	return search.QueryID
}
