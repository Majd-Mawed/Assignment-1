package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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
	fmt.Println("Working")
}
