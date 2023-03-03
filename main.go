package main

import (
	"Assignemnt1/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	http.HandleFunc("/", handler.DefaultHandler)
	http.HandleFunc("/unisearcher/v1/uniinfo/", handler.HandleUniRequest)
	http.HandleFunc("/unisearcher/v1/neighbourunis/", handler.HandleNabUniRequest)
	http.HandleFunc("/unisearcher/v1/Diag/", handler.HandleDiagRequest)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
