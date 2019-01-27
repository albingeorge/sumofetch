# sumofetch
Fetch and format Sumologic logs for Certification

### Instructions

1. Fetch the SumoLogic AccessID and Access Key
2. Move `config/config.json.sample` to `config/config.json` and replace the access ID and key
3. Execute the below to fetch the payment's relevant logs

  `go run main.go <payment_id>`

### TODO
- Print output to a word document instead of to stdout
- Handle callback request timeout case
- Handle search job status, i.e poll the search status API till it's ready to be queried. As of now, we're simply waiting for 2 seconds.

### Improvements
 
 - As of now, the time window in which the search is made is hardcoded to 2 hours back till now. This can be improved by accepting maybe a timerange as a command line argument, just like payment id.


 
