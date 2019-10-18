package smarketstools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type LastExecutedPrices struct {
	LastEx map[string][]LastExecutedPrice `json:"last_executed_prices"`
}

type LastExecutedPrice struct {
	ContractID        string `json:"contract_id"`
	LastExecutedPrice string `json:"last_executed_price"`
	Timestamp         string `json:"timestamp"`
}

type Quote struct {
	Price    int `json:"price"`
	Quantity int `json:"quantity"`
}

type Quotes struct {
	ContractID string
	Bids       []Quote `json:"bids"`
	Offers     []Quote `json:"offers"`
}

func (l LastExecutedPrices) GetID(id string) []LastExecutedPrice {
	return l.LastEx[id]
}

//Given a marketID, retrieves a list of quotes for each contract
func GetQuoteCollection(marketID string, client Client) ([]Quotes, []string) {

	var result map[string]interface{}

	var contracts []string

	url := fmt.Sprintf("https://api.smarkets.com/v3/markets/%s/quotes/", marketID)

	GetJSON(url, &result, client)
	var quoteCollection []Quotes

	for contractID, _ := range result {
		b, _ := json.Marshal(result[contractID])
		quotes := new(Quotes)
		json.Unmarshal(b, &quotes)
		quotes.ContractID = contractID
		quoteCollection = append(quoteCollection, *quotes)
		contracts = append(contracts, contractID)
	}
	return quoteCollection, contracts
}

func GetLastExecutedPrices(marketID string, client Client) *LastExecutedPrices {
	response := new(LastExecutedPrices)
	apiPath := fmt.Sprintf("https://api.smarkets.com/v3/markets/%s/last_executed_prices/", marketID)

	GetJSON(apiPath, response, client)

	return response
}

func GetJSON(url string, target interface{}, client Client) error {
	myClient := &http.Client{Timeout: 10 * time.Second}

	request, err := http.NewRequest("GET", url, nil)
	checkErr(err)

	tokenHeader := fmt.Sprintf("Session-Token %s", client.ApiToken)
	request.Header.Set("Authorization", tokenHeader)

	resp, err := myClient.Do(request)
	checkErr(err)

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	return json.Unmarshal(body, &target)
}
