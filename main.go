package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Define a structure for the API response
type Quote struct {
	CurrentPrice  float64 `json:"c"`
	HighPrice     float64 `json:"h"`
	LowPrice      float64 `json:"l"`
	OpenPrice     float64 `json:"o"`
	PreviousClose float64 `json:"pc"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: stockcli <symbol>")
		os.Exit(1)
	}

	symbol := os.Args[1]

	apiKey := "d0tl6s9r01qlvahcrf40d0tl6s9r01qlvahcrf4g" // Replace this with your actual key

	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Received status code %d", resp.StatusCode)
	}

	var quote Quote
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	fmt.Printf("Stock: %s\n", symbol)
	fmt.Printf("Current Price: $%.2f\n", quote.CurrentPrice)
	fmt.Printf("High Price: $%.2f\n", quote.HighPrice)
	fmt.Printf("Low Price: $%.2f\n", quote.LowPrice)
	fmt.Printf("Open Price: $%.2f\n", quote.OpenPrice)
	fmt.Printf("Previous Close: $%.2f\n", quote.PreviousClose)

	//Save data to a file
	filename := fmt.Sprintf("%s.txt", symbol)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Stock: %s\n", symbol)
	_, err = fmt.Fprintf(file, "Current Price: $%.2f\n", quote.CurrentPrice)
	_, err = fmt.Fprintf(file, "High Price: $%.2f\n", quote.HighPrice)
	_, err = fmt.Fprintf(file, "Low Price: $%.2f\n", quote.LowPrice)
	_, err = fmt.Fprintf(file, "Open Price: $%.2f\n", quote.OpenPrice)
	_, err = fmt.Fprintf(file, "Previous Close: $%.2f\n", quote.PreviousClose)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	fmt.Printf("Data saved to %s\n", filename)
}
