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
func HandleUniRequest(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	name := parts[4]

	url := "http://universities.hipolabs.com/search?name=" + name

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
	uni := []UNI{}

	// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
	err = decoder.Decode(&uni)
	if err != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	naburl := "https://restcountries.com/v3.1/name/" + uni[0].Country

	nabNewRequest, naberr := http.NewRequest(http.MethodGet, naburl, nil)
	if naberr != nil {
		fmt.Errorf("Error in creating request:", naberr.Error())
	}

	// Setting content type -> effect depends on the service provider
	nabNewRequest.Header.Add("content-type", "application/json")

	nabclient := &http.Client{}
	defer nabclient.CloseIdleConnections()

	nabres, naberr := client.Do(nabNewRequest)
	//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
	if naberr != nil {
		fmt.Errorf("Error in response:", naberr.Error())
	}

	// Instantiate decoder
	nabdecoder := json.NewDecoder(nabres.Body)
	// Ensure parser fails on unknown fields (baseline way of detecting different structs than expected ones)
	// Note: This does not lead to a check whether an actually provided field is empty!

	// Prepare empty struct to populate
	nabuni := []NABUNI{}

	// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
	err = nabdecoder.Decode(&nabuni)
	if naberr != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+naberr.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v", uni)
	fmt.Fprintf(w, "%v", nabuni)

	// Return status code (good practice)
	http.Error(w, "\nOK", http.StatusOK)
}
