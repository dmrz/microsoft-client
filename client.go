package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// APIClient comment
type APIClient struct {
	baseEndpoint string
	version      string
	key          string
	client       *http.Client
}

func (api *APIClient) getURL(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", api.baseEndpoint, api.version, endpoint)
}

func (api *APIClient) sentiment(language string, id int, text string) SentimentResponse {
	requestBody := new(bytes.Buffer)
	requestBodyData := struct {
		Documents []Document
	}{
		[]Document{
			{language, strconv.Itoa(id), text},
		},
	}
	json.NewEncoder(requestBody).Encode(requestBodyData)

	req, err := http.NewRequest("POST", api.getURL("sentiment"), requestBody)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", api.key)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	res, err := api.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	sentimentResponse := SentimentResponse{}
	json.NewDecoder(res.Body).Decode(&sentimentResponse)
	return sentimentResponse

}

// For debugging purpose (just to view raw response body)
func printResponseBody(res *http.Response) {
	defer res.Body.Close()
	responseBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseBody := string(responseBodyBytes)
	fmt.Println(responseBody)
}
