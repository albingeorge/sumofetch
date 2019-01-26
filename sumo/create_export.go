package sumo

import "fmt"

func (sumo Sumocreds) createExportID(queryID string) string {
	url := "https://service.eu.sumologic.com/json/v2/searchquery/" + queryID + "/export"

	createExportInputs := generateCreateExportInputs(queryID)

	fmt.Println(url, createExportInputs)

	return ""
}

func generateCreateExportInputs(queryID string) []byte {
	return []byte{}
}
