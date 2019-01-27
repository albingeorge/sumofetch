package formatter

import (
	"bitbucket.org/albingeorgee/sumofetch/globals"
	"bitbucket.org/albingeorgee/sumofetch/sumo"
	"strings"
	"time"
)

type FormattedContent struct {
	Header   string
	DateTime time.Time
	Content  string
	Footer   string
}

func Format(format []sumo.ResponseFormat) []FormattedContent {
	result := []FormattedContent{}
	for _, response := range format {
		if response.Code == globals.GATEWAY_SOAP_REQUEST {
			result = append(result, getSoapRequestFormattedContent(response))
			result = append(result, getSoapResponseFormattedContent(response))
		}

		if response.Code == globals.GATEWAY_PAYMENT_REQUEST {
			result = append(result, getRedirectRequestFormattedContent(response))
		}

		if response.Code == globals.PAYMENT_CALLBACK_REQUEST {
			result = append(result, getCallbackRequestFormattedContent(response))
		}
	}

	return result
}

func getSoapRequestFormattedContent(response sumo.ResponseFormat) FormattedContent {
	result := FormattedContent{}
	command := response.Command
	result.Header = strings.Title(command) + " request"

	result.DateTime = response.DateTime

	result.Content = response.SoapRequest

	return result
}

func getSoapResponseFormattedContent(response sumo.ResponseFormat) FormattedContent {
	result := FormattedContent{}
	command := response.Command
	result.Header = strings.Title(command) + " response"

	result.DateTime = response.DateTime

	result.Content = response.SoapResponse

	return result
}

func getRedirectRequestFormattedContent(response sumo.ResponseFormat) FormattedContent {
	result := FormattedContent{}

	result.Header = "Redirect request"
	result.DateTime = response.DateTime

	result.Content = `URL: https://cert.swasrec.npci.org.in/redirect_ias/home/IssuerReg
Method: POST

Parameters:
`
	result.Content = result.Content + formatKeyValuePairs(response.RedirectRequest)

	return result
}

func getCallbackRequestFormattedContent(response sumo.ResponseFormat) FormattedContent {
	result := FormattedContent{}

	result.Header = "Callback request"
	result.DateTime = response.DateTime

	result.Content = `Parameters:
`
	result.Content = result.Content + formatKeyValuePairs(response.CallbackRequest)

	return result
}

func formatKeyValuePairs(input map[string]string) string {

	res := ""
	for key, val := range input {
		res = res + "\"" + key + "\": \"" + val + "\"\n"
	}

	return res
}
