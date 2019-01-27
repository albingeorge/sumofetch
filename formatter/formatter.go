package formatter

import (
	"github.com/albingeorge/sumofetch/globals"
	"github.com/albingeorge/sumofetch/sumo"
	"strings"
	"time"
)

type FormattedContent struct {
	Header   string
	DateTime time.Time
	Content  string
}

func Format(format []sumo.ResponseFormat) []FormattedContent {
	result := []FormattedContent{}

	callbackCount := 0

	for _, response := range format {
		if response.Code == globals.GATEWAY_SOAP_REQUEST {
			result = append(result, getSoapRequestFormattedContent(response))
			result = append(result, getSoapResponseFormattedContent(response))
		}

		if response.Code == globals.GATEWAY_PAYMENT_REQUEST {
			result = append(result, getRedirectRequestFormattedContent(response))
		}

		if response.Code == globals.PAYMENT_CALLBACK_REQUEST {
			callbackCount = callbackCount + 1
			result = append(result, getCallbackRequestFormattedContent(response))

			// If this is the second callback, decline the transaction
			if callbackCount > 1 {
				r := FormattedContent{
					Header:   "Payment declined due to duplicate callback",
					DateTime: response.DateTime,
				}

				result = append(result, r)
			}
		}

		if response.Code == globals.GATEWAY_REQUEST_TIMEOUT {
			result = append(result, getGatewayTimeoutRequestFormattedContent(response))

			// After adding the log for Soap Request, we should also add that the payment
			// timed out at the current log's time.
			// This need not be there for authorize command, since if authorize times out
			// we can transaction status to verify the payment
			if response.Command != "authorize" {
				r := FormattedContent{
					Header:   "Payment failed because request timed out",
					DateTime: response.DateTime,
					Content:  "",
				}

				result = append(result, r)
			} else {
				r := FormattedContent{
					Header:   "Authorize request timed out",
					DateTime: response.DateTime,
				}

				result = append(result, r)
			}

		}

		if response.Code == globals.GATEWAY_CHECKSUM_VERIFY_FAILED {
			r := FormattedContent{
				Header:   "Payment declined due to invalid checksum",
				DateTime: response.DateTime,
			}

			result = append(result, r)
		}

		if response.Code == globals.GATEWAY_ERROR_DATA_MISMATCH {
			r := FormattedContent{
				Header:   "Payment declined due to data mismatch",
				DateTime: response.DateTime,
			}

			result = append(result, r)
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

func getGatewayTimeoutRequestFormattedContent(response sumo.ResponseFormat) FormattedContent {
	result := FormattedContent{}

	//fmt.Printf("%+v\n", response)

	command := response.Command
	result.Header = strings.Title(command) + " request"
	result.Content = response.SoapRequest

	// In log, we don't show when the request was sent. The log for GATEWAY_REQUEST_TIMEOUT is made
	// when the request actually times out - meaning, we need to set the time back that many seconds
	// based on the action to show when the request was sent

	timeoutSeconds := globals.TimeoutCommandMaps[command]
	timeoutSeconds = -1 * timeoutSeconds
	result.DateTime = response.DateTime.Add(time.Duration(timeoutSeconds) * time.Second)

	return result
}

func formatKeyValuePairs(input map[string]string) string {

	res := ""
	for key, val := range input {
		res = res + "\"" + key + "\": \"" + val + "\"\n"
	}

	return res
}
