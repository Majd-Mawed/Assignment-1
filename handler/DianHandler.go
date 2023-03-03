package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func HandleDiagRequest(w http.ResponseWriter, r *http.Request) {
	uniresp, unierr := http.Get("http://universities.hipolabs.com/")
	if unierr != nil {
		fmt.Println("Error:", unierr)
		return
	}
	defer uniresp.Body.Close()
	fmt.Fprintf(w, "universitiesapi: ")
	fmt.Fprintf(w, uniresp.Status)
	fmt.Fprintf(w, "\n")

	countryresp, countryerr := http.Get("https://restcountries.com/")
	if countryerr != nil {
		fmt.Println("Error:", countryerr)
		return
	}
	defer countryresp.Body.Close()
	fmt.Fprintf(w, "countriesapi: ")
	fmt.Fprintf(w, countryresp.Status)
	fmt.Fprintf(w, "\n")

	parts := strings.Split(r.URL.Path, "/")
	version := parts[2]
	fmt.Fprintf(w, "version: ")
	fmt.Fprintf(w, version)
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "Uptime: ")
	fmt.Fprintf(w, "%s", uptime())

}

// https://stackoverflow.com/questions/37992660/golang-retrieve-application-uptime
