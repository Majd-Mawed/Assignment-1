# 02-JSON-demo

Simple module demonstrating the use of JSON structs and collections.

Per default the server will be available under http://localhost:8080, with endpoints
 * http://localhost:8080/location (for location information)
 * http://localhost:8080/collection (for map structures with untyped value)

## Deploying the demo

* Open in IDE of choice and run main.go

## Using the demo

Use an HTTP client of choice (e.g., Browser, Postman) to point requests to the server URL. GET requests will return the intended structure for either endpoint.

Using Postman (or another configurable HTTP client), explore the variation of JSON content by sending POST request. Explore header information, as well as error handling (e.g., incomplete JSON content).

* Example POST dynamically structured content for collection endpoint:
    * `{"first":"firstValue","second":2,"third":3}`
    
* Example POST content for location endpoint:
  * `{"name":"Gjøvik","code":"2815","country":"Norway","location":{"lat":60.7847024,"lon":10.6891797}}`

Exercises:
- Explore JSON en-/decoding (e.g., hardening, validation)
- Provide functionality where indicated by TODO labels
- Modify structs to explore effects
  

