package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Define the API endpoint
	url := "https://api.weather.gov/points/39.7456,-97.0892"

	// Send the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making the request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading the response body: %v", err)
	}

	// Print the response
	fmt.Println("Response:", string(body))
}
