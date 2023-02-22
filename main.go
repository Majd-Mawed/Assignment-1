package main

import (
	"Assignment 1/ handler"
	"fmt"
	"log"
	"net/http"
	"os"
)

type uni struct {
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
	http.HandleFunc("/unisearcher/v1/uniinfo/", printhttp)
	http.HandleFunc("/unisearcher/v1/neighbourunis/", printhttp)
	http.HandleFunc("/unisearcher/v1/diag/", printhttp)

	log.Println("Running on port", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func printhttp(w http.ResponseWriter, r *http.Request) {

	urluni := "http://universities.hipolabs.com/"
	urlnab := "https://restcountries.com/"

	// Create new request
	runi, erruni := http.NewRequest(http.MethodGet, urluni, nil)
	if erruni != nil {
		fmt.Errorf("Error in creating request:", erruni.Error())
	}

	rnab, errnab := http.NewRequest(http.MethodGet, urlnab, nil)
	if errnab != nil {
		fmt.Errorf("Error in creating request:", errnab.Error())
	}

	// Setting content type -> effect depends on the service provider
	runi.Header.Add("content-type", "application/json")
	rnab.Header.Add("content-type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	resuni, erruni := client.Do(runi)
	//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
	if erruni != nil {
		fmt.Errorf("Error in response:", erruni.Error())
	}

	resnab, errnab := client.Do(rnab)
	//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
	if errnab != nil {
		fmt.Errorf("Error in response:", errnab.Error())
	}

	fmt.Println("Status:", resuni.Status)
	fmt.Println("Status code:", resuni.StatusCode)

	fmt.Println("Content type:", resuni.Header.Get("content-type"))
	fmt.Println("Protocol:", resuni.Proto)

	fmt.Println("Status:", resnab.Status)
	fmt.Println("Status code:", resnab.StatusCode)

	fmt.Println("Content type:", resnab.Header.Get("content-type"))
	fmt.Println("Protocol:", resnab.Proto)
}
