package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type UNI struct {
	name      string `json:"name"`
	country   string `json:"country"`
	isocode   string `json:"isocode"`
	webpages  string `json:"webpages"`
	languages string `json:"languages"`
	Map       string `json:"Map"`
}

func main() {

	// Define port
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Standard http server with reference to stubbed handler
	// If you want to adapt this, ensure to adjust path for compatibility with project
	http.HandleFunc("/unisearcher/v1/uniinfo/", UniFunction)
	//http.HandleFunc("/unisearcher/v1/neighbourunis/", printhttp)
	//http.HandleFunc("/unisearcher/v1/diag/", printhttp)

	log.Println("Running on port", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func UniFunction(w http.ResponseWriter, r *http.Request) {

	url := "http://universities.hipolabs.com/"

	// Create new request
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Setting content type -> effect depends on the service provider
	r.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	// HTTP Header content
	fmt.Println("Status:", res.Status)
	fmt.Println("Status code:", res.StatusCode)

	fmt.Println("Content type:", res.Header.Get("content-type"))
	fmt.Println("Protocol:", res.Proto)

	decoder := json.NewDecoder(r.Body)
	// Ensure parser fails on unknown fields (baseline way of detecting different structs than expected ones)
	// Note: This does not lead to a check whether an actually provided field is empty!
	decoder.DisallowUnknownFields()

	// Prepare empty struct to populate
	uni := UNI{}

	// Decode location instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&location)"

	// Validation of input (Golang does not do that itself :()

	// TODO: Write convenience function for validation

	if uni.name == "" {
		http.Error(w, "Invalid input: Field 'Name' is empty.", http.StatusBadRequest)
		return
	}

	if uni.country == "" {
		http.Error(w, "Invalid input: Field 'Postcode' not found.", http.StatusBadRequest)
		return
	}

	// Flat printing
	fmt.Println("Received following request:")
	fmt.Println(uni)

	// Pretty printing
	output, err := json.MarshalIndent(uni, "", "  ")
	if err != nil {
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}

	fmt.Println("Pretty printing:")
	fmt.Println(string(output))

	// TODO: Handle content (e.g., writing to DB, process, etc.)

	// Return status code (good practice)
	http.Error(w, "OK", http.StatusOK)

}
