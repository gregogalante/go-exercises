package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080
	log.Printf("Server starting on port %v\n", port)

	// call the HandleFunc method which creates a Handler type
	http.HandleFunc("/helloworld", helloWorldHandler)
	
	// start the HTTP server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// create an interface used for the Marshal function
	response := helloWorldResponse{Message: "Hello World"}
	// generate encoder cor the current response writer
	encoder := json.NewEncoder(w) // NB: Usage of encoder is 50% faster than Marshal
	// encode response
	encoder.Encode(response)
}