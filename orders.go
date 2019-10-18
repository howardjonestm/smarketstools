package smarketstools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Exchange order
type Order struct {
	ContractID          string `json:"contract_id"`
	Label               string `json:"label"`
	MarketID            string `json:"market_id"`
	MinAcceptedQuantity int    `json:"minimum_accepted_quantity"`
	Price               int    `json:"price"`
	Quantity            int    `json:"quantity"`
	Side                string `json:"side"`
	Type                string `json:"type"`
}

func PlaceOrder(order Order, token string) {
	url := "https://api.smarkets.com/v3/orders/"

	sessionAuthHeader := fmt.Sprintf("Session-Token %s", string(token))

	orderJSON, err := json.Marshal(order)
	checkErr(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(orderJSON))
	checkErr(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", sessionAuthHeader)

	client := &http.Client{}
	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body", string(body))

}
