package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func ParseFile(filename string) []byte {
	file, e := os.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	return file
}

/*
Responds with fixed JSON output sourced from provided file.
*/
func StubHandlerCountries(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		log.Println("Received " + r.Method + " request on Countries stub handler. Returning mocked information.")
		w.Header().Add("content-type", "application/json")
		output := ParseFile("./res/countries.json")
		fmt.Fprint(w, string(output))
		break
	default:
		http.Error(w, "Method not supported", http.StatusNotImplemented)
	}
}

/*
Responds with fixed JSON output sourced from provided file.
*/
func StubHandlerOccurrences(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		log.Println("Received " + r.Method + " request on Occurrences stub handler. Returning mocked information.")
		w.Header().Add("content-type", "application/json")
		output := ParseFile("./res/occurrences.json")
		fmt.Fprint(w, string(output))
		break
	default:
		http.Error(w, "Method not supported", http.StatusNotImplemented)
	}
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
	http.HandleFunc("/countries/no", StubHandlerCountries)
	// Naturally, you can introduce multiple handlers to emulate different data sources
	http.HandleFunc("/species/occurrences", StubHandlerOccurrences)

	log.Println("Running on port", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}
