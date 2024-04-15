package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EmailVerificationResponse struct {
	Data struct {
		Result string `json:"result"`
	} `json:"data"`
}

func main() {
	email := "Ghostman312@gmail.com" 
	apiKey := "a78afa97d76af0e3364a3eb68ed12aae83e247a0"       

	url := fmt.Sprintf("https://api.hunter.io/v2/email-verifier?email=%s&api_key=%s", email, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var verificationResponse EmailVerificationResponse
	err = json.NewDecoder(resp.Body).Decode(&verificationResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if verificationResponse.Data.Result == "undeliverable" {
		fmt.Println("Email address does not exist or is undeliverable")
	} else {
		fmt.Println("Email address is valid")
	}
}
