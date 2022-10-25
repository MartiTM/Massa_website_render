package main

import (
	"bytes"
	"net/http"
)

var (
	massaAddress string = "A12VViTPc3ZQu3PGp53oYXDPA6iydaeT69abcFTM2QXo6BFKMsv7"
	nodeAddress  string = "http://localhost:33035/"
)

func main() {
	getMassaWebsite(massaAddress, nodeAddress)
}

func getMassaWebsite(massaAddress string, nodeAddress string) (*http.Response, error){
	// we are looking for the key "massa_web" in decimal
	body := []byte(`{
		"jsonrpc":"2.0",
		"method":"get_datastore_entries",
		"params":[[{
			"address":"` + massaAddress + `",
			"key":[109,97,115,115,97,95,119,101,98]
		}]],
		"id":1}`)

	resp, err := http.Post(nodeAddress, "application/json", bytes.NewBuffer(body))

	return resp, err
}
