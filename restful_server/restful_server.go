package restful_server

import (
	"RESTclient/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
Handler for articles
*/
func ArticleApiRouter() func(http.ResponseWriter, *http.Request) {

	return func(rw http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			fmt.Println("Router received " + r.Method + " method.")
			utils.HandleArticleGet(rw, r)
		case http.MethodPost:
			fmt.Println("Router received " + r.Method + " Method.")
			utils.HandleArticlePost(rw, r)
		default: // If none of the above conditions, return this by default.
			fmt.Println("Method not defined (Method: " + r.Method + ")")
			http.Error(rw, "Method "+r.Method+" not supported.", http.StatusNotImplemented)
		}
	}
}

/*
Handler for default homepage, listing service features.
*/
func RestFulHomePage(rw http.ResponseWriter, r *http.Request) {
	// If the method is Get, it will return the content, otherwise, will rise an error
	if r.Method == http.MethodGet {
		// An example Article to show as response in home page call
		article := utils.Article{Title: "Article title", Desc: "Article description", Content: "Article content"}
		output, err := json.Marshal(article) // Converting struct to Json
		if err != nil {
			http.Error(rw, "Error during JSON marshaling.", http.StatusInternalServerError)
		}
		// To have a pretty output for existence endpoints, we defined a struct for them.
		endpoints := utils.Endpoints{
			utils.Endpoint{Url: "/", Description: "GET: Welcome to the PROG2005 Cloud Technologies home page!"},
			utils.Endpoint{Url: "/article?id=SomeThing", Description: "GET: returns an article by id (where {id} is the article ID of concern)"},
			utils.Endpoint{Url: "/article", Description: "GET: returns all articles"},
			utils.Endpoint{Url: "/article", Description: "POST: Takes information for a new article and adds it to the memory. Article information should be provided using the schema below (but without the ID field, since that is generated and returned by the service upon adding)." +
				";  Schema: " + string(output) + ". If no input structure is provided, a random new article is added to memory."}}

		rw.Header().Set("Content-Type", "application/json") // Writes the response content type
		err = json.NewEncoder(rw).Encode(endpoints)
		if err != nil {
			http.Error(rw, "Error during JSON encoding (Error: "+err.Error()+")", http.StatusInternalServerError)
		}
	} else {
		// If the method is not Get.
		http.Error(rw, "Method "+r.Method+" not supported.", http.StatusNotImplemented)
	}
}

func RestFulApi() {
	// Home page which lists available endpoints.
	http.HandleFunc(utils.HomeEndPoint, RestFulHomePage)
	// An endpoint in which returns
	//								- Specific article by id (Method: Get; parameter: is),
	//								- All articles (Method: Get; parameter: no parameter),
	//								- Adds an article (Method: Post, body: Article Struct),
	//								- Adds a random article (Method: Post, body: empty body)
	http.HandleFunc(utils.ArticleEndPoint, ArticleApiRouter())
	// In Cloud environments, specially (PaaS), ports are managed by the system. To get the allocated port we
	// implemented a function which gets the port from system environment.
	Port := utils.GetPort(utils.MainPort)
	fmt.Println("RestFul API: Launched under: http://localhost" + Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
