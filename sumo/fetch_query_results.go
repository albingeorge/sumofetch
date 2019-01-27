package sumo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/albingeorge/sumofetch/globals"
	"net/http"
	"strconv"
	"time"
)

func (sumo Sumocreds) fetchQueryResults(queryID string) []ResponseFormat {
	url := sumo.BaseUrl + "/" + queryID + "/messages?offset=0&limit=20"

	req, _ := http.NewRequest("GET", url, nil)
	sumo.setHeaders(req)

	time.Sleep(2 * time.Second)
	response := sendRequest(req)

	res := parseResponse(response)

	return res
}

type Content struct {
	Code                    string `json:"code"`
	Timestamp               string `json:"_messagetime"`
	PaymentId               string `json:"context.payment_id"`
	SoapResponse            string `json:"context.response.callpaysecureresult"`
	SoapRequest             string `json:"context.request"`
	Command                 string `json:"context.command"`
	AccuResponseCode        string `json:"context.gateway_input.accuresponsecode"`
	Session                 string `json:"context.gateway_input.session"`
	AccuGuid                string `json:"context.gateway_input.accuguid"`
	AccuRequestId           string `json:"context.gateway_input.accurequestid"`
	RequestCommand          string `json:"context.request.command"`
	RequestAccuCardholderId string `json:"context.request.content.accucardholderid"`
	RequestAccuGuid         string `json:"context.request.content.accuguid"`
	RequestAccuReturnURL    string `json:"context.request.content.accureturnurl"`
	RequestSession          string `json:"context.request.content.accurequestsession"`
	RequestAccuRequestId    string `json:"context.request.content.accurequestid"`
}

type MessageMaps struct {
	Content Content `json:"map"`
}

type JsonData struct {
	Messages []MessageMaps `json:"messages"`
}

type ResponseFormat struct {
	Code            string
	DateTime        time.Time
	Command         string
	SoapRequest     string
	SoapResponse    string
	CallbackRequest map[string]string
	RedirectRequest map[string]string
}

// Sample format which is getting parsed
//     `{
//			"messages": [
//				{
//					"map": {
//						"code": "/app/app/Gateway/Paysecure/RequestHandlerTrait.php"
//					}
//				}
//			]
//		}`
func parseResponse(response []byte) []ResponseFormat {
	var r JsonData
	err := json.Unmarshal(response, &r)

	if err != nil {
		fmt.Println("Json Unmarshalling error", err)
	}

	res := formatResponseMap(r)

	return res
}

func formatResponseMap(json JsonData) []ResponseFormat {
	res := []ResponseFormat{}

	// Loops in the reverse order
	for i := len(json.Messages); i > 0; i-- {
		message := json.Messages[i-1]

		r, err := parseSingleMessage(message.Content)

		if err == nil {
			res = append(res, r)
		}
	}

	//for _, message := range json.Messages {
	//	r, err := parseSingleMessage(message.Content)
	//
	//	if err == nil {
	//		res = append(res, r)
	//	}
	//}

	return res
}

func parseSingleMessage(content Content) (ResponseFormat, error) {

	date := formatDateTime(content.Timestamp)

	r := ResponseFormat{
		Code:     content.Code,
		DateTime: date,
	}
	// This is soap requests in readable format. Ignore these
	if content.Code == globals.GATEWAY_PAYMENT_REQUEST {
		//fmt.Println("content.RequestCommand", content.RequestCommand)
		if content.RequestCommand != "" {
			return r, errors.New("Ignore this")
		}

		// If redirect Request
		r.RedirectRequest = map[string]string{
			"AccuCardholderId": content.RequestAccuCardholderId,
			"AccuGuid":         content.RequestAccuGuid,
			"AccuReturnURL":    content.RequestAccuReturnURL,
			"session":          content.RequestSession,
			"AccuRequestId":    content.RequestAccuRequestId,
		}

	}

	if content.Code == globals.PAYMENT_CALLBACK_REQUEST {
		r.CallbackRequest = map[string]string{
			"AccuResponseCode": content.AccuResponseCode,
			"session":          content.Session,
			"AccuGuid":         content.AccuGuid,
			"AccuRequestId":    content.AccuRequestId,
		}
	}

	if content.Code == globals.GATEWAY_SOAP_REQUEST {
		r.SoapRequest = content.SoapRequest
		r.SoapResponse = content.SoapResponse
		r.Command = content.Command
	}

	return r, nil
}

func formatDateTime(d string) time.Time {
	intTime, _ := strconv.ParseInt(d, 10, 64)
	date := time.Unix((intTime / 1000), 0)
	istLoc, _ := time.LoadLocation("Asia/Kolkata")
	date = date.In(istLoc)

	return date
}
