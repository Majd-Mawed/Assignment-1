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

	name := ""

	parts := strings.Split(r.URL.Path, "/")
	if parts[4] == "" {
		name = "all"
	} else {
		name = "name/" + parts[4]
	}

	url := "https://restcountries.com/v3.1/" + name

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
	nabuni := []NABUNIBORDERS{}

	// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
	err = decoder.Decode(&nabuni)
	if err != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	naburl := "https://restcountries.com/v3.1/alpha?codes="

	for i := 0; i < len(nabuni[0].Borders); i++ {
		naburl += nabuni[0].Borders[i] + ","
	}

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
	nabuni1 := []NABUNI{}

	// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
	err = nabdecoder.Decode(&nabuni1)
	if naberr != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+naberr.Error(), http.StatusBadRequest)
		return
	}

	if len(parts) > 5 {
		for a := 0; a < len(nabuni1); a++ {
			uniname := parts[5]
			url2 := "http://universities.hipolabs.com/search?name=" + uniname + "&country=" + nabuni1[a].Name.Name

			lastNewRequest, lasterr := http.NewRequest(http.MethodGet, url2, nil)
			if lasterr != nil {
				fmt.Errorf("Error in creating request:", lasterr.Error())
			}

			lastNewRequest.Header.Add("content-type", "application/json")

			lastclient := &http.Client{}
			defer lastclient.CloseIdleConnections()

			lastres, lasterr := client.Do(lastNewRequest)
			//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
			if lasterr != nil {
				fmt.Errorf("Error in response:", lasterr.Error())
			}

			// Instantiate decoder
			lastdecoder := json.NewDecoder(lastres.Body)
			// Ensure parser fails on unknown fields (baseline way of detecting different structs than expected ones)
			// Note: This does not lead to a check whether an actually provided field is empty!

			// Prepare empty struct to populate
			lastuni := []UNI{}

			// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
			err = lastdecoder.Decode(&lastuni)
			if lasterr != nil {
				// Note: more often than not is this error due to client-side input, rather than server-side issues
				http.Error(w, "Error during decoding: "+lasterr.Error(), http.StatusBadRequest)
				return
			}
			//fmt.Fprintf(w, "%v", lastuni)

			if len(lastuni) > 1 {
				for i := 0; i < len(lastuni)/2; i++ {
					fmt.Fprintf(w, "Name: ")
					fmt.Fprintf(w, lastuni[i].Name)
					fmt.Fprintf(w, "\n")
					fmt.Fprintf(w, "Country: ")
					fmt.Fprintf(w, lastuni[i].Country)
					fmt.Fprintf(w, "\n")
					for y := 0; y < len(lastuni[i].Webpages); y++ {
						fmt.Fprintf(w, "Webpages: ")
						fmt.Fprintf(w, lastuni[i].Webpages[y])
						fmt.Fprintf(w, "\n")
					}
					fmt.Fprintf(w, "Isocode: ")
					fmt.Fprintf(w, lastuni[i].Isocode)
					fmt.Fprintf(w, "\n")

					newurl := "https://restcountries.com/v3.1/name/" + lastuni[i].Country
					newNewRequest, newerr := http.NewRequest(http.MethodGet, newurl, nil)
					if newerr != nil {
						fmt.Errorf("Error in creating request:", newerr.Error())
					}
					newNewRequest.Header.Add("content-type", "application/json")
					newclient := &http.Client{}
					defer newclient.CloseIdleConnections()
					newres, newerr := client.Do(nabNewRequest)
					//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
					if newerr != nil {
						fmt.Errorf("Error in response:", newerr.Error())
					}
					newdecoder := json.NewDecoder(newres.Body)
					newuni := []NABUNIINFO{}
					newerr = newdecoder.Decode(&newuni)
					if newerr != nil {
						// Note: more often than not is this error due to client-side input, rather than server-side issues
						http.Error(w, "Error during decoding: "+newerr.Error(), http.StatusBadRequest)
						return
					}
					fmt.Fprintf(w, "Languages: ")
					fmt.Fprintf(w, "%v", newuni[a].Languages)
					fmt.Fprintf(w, "\n")
					fmt.Fprintf(w, "Maps: ")
					fmt.Fprintf(w, "%v", newuni[a].Map)
					fmt.Fprintf(w, "\n")
					fmt.Fprintf(w, "\n")

				}
			} else {
				for i := 0; i < len(lastuni); i++ {
					fmt.Fprintf(w, "Name: ")
					fmt.Fprintf(w, lastuni[i].Name)
					fmt.Fprintf(w, "\n")
					fmt.Fprintf(w, "Country: ")
					fmt.Fprintf(w, lastuni[i].Country)
					fmt.Fprintf(w, "\n")
					for y := 0; y < len(lastuni[i].Webpages); y++ {
						fmt.Fprintf(w, "Webpages: ")
						fmt.Fprintf(w, lastuni[i].Webpages[y])
						fmt.Fprintf(w, "\n")
					}
					fmt.Fprintf(w, "Isocode: ")
					fmt.Fprintf(w, lastuni[i].Isocode)
					fmt.Fprintf(w, "\n")

					newurl := "https://restcountries.com/v3.1/name/" + lastuni[i].Country
					newNewRequest, newerr := http.NewRequest(http.MethodGet, newurl, nil)
					if newerr != nil {
						fmt.Errorf("Error in creating request:", newerr.Error())
					}
					newNewRequest.Header.Add("content-type", "application/json")
					newclient := &http.Client{}
					defer newclient.CloseIdleConnections()
					newres, newerr := client.Do(nabNewRequest)
					//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
					if newerr != nil {
						fmt.Errorf("Error in response:", newerr.Error())
					}
					newdecoder := json.NewDecoder(newres.Body)
					newuni := []NABUNIINFO{}
					newerr = newdecoder.Decode(&newuni)
					if newerr != nil {
						// Note: more often than not is this error due to client-side input, rather than server-side issues
						http.Error(w, "Error during decoding: "+newerr.Error(), http.StatusBadRequest)
						return
					}
					fmt.Fprintf(w, "Languages: ")
					fmt.Fprintf(w, "%v", newuni[a].Languages)
					fmt.Fprintf(w, "\n")
					fmt.Fprintf(w, "Maps: ")
					fmt.Fprintf(w, "%v", newuni[a].Map)
					fmt.Fprintf(w, "\n")
					fmt.Fprintf(w, "\n")
				}
			}
		}

	} else {
		fmt.Fprintf(w, "%v", nabuni)
	}

	// Return status code (good practice)
	http.Error(w, "OK", http.StatusOK)
}
