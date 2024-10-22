package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Data    []TxData `json:"data"`
	Success bool     `json:"success"`
}

type TxData struct {
	Note            string `json:"note"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	TransactionTime int64  `json:"transactionTime"`
}

const (
	apiKey    = ""
	apiSecret = ""
	baseURL   = "https://api.binance.com"
)

func main() {
	endpoint := "/sapi/v1/pay/transactions"
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	query := url.Values{}
	query.Set("timestamp", timestamp)
	signature := createSignature(query.Encode())
	query.Set("signature", signature)
	fullURL := fmt.Sprintf("%s%s?%s", baseURL, endpoint, query.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("X-MBX-APIKEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for _, tx := range response.Data {
		fmt.Printf("Note: %s\n", tx.Note)
		fmt.Printf("Amount: %s\n", tx.Amount)
		fmt.Printf("Currency: %s\n", tx.Currency)
		fmt.Printf("Transaction Time: %d\n\n", tx.TransactionTime)
	}
}

func createSignature(data string) string {
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
