package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
Dedicated handler for POST requests
*/
func HandleNabUniRequest(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	countryname := parts[4]
	uniname := parts[5]

	url := "https://restcountries.com/search?name=" + countryname + "/" + uniname

	NewRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Setting content type -> effect depends on the service provider
	NewRequest.Header.Add("content-type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()

	res, err := client.Do(NewRequest)
	//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	// Instantiate decoder
	decoder := json.NewDecoder(res.Body)
	// Ensure parser fails on unknown fields (baseline way of detecting different structs than expected ones)
	// Note: This does not lead to a check whether an actually provided field is empty!

	// Prepare empty struct to populate
	nabuni := []NABUNI{}

	// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
	err = decoder.Decode(&nabuni)
	if err != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(nabuni)
	//fmt.Fprintf(w, "%v", nabuni)

	// Return status code (good practice)
	http.Error(w, "OK", http.StatusOK)
}
