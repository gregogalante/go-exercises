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

type helloWorldRequest struct {
	Name string `json:"name"`
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
	// read request json
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// create an interface used for the Marshal function
	response := helloWorldResponse{Message: "Hello " + request.Name}
	// generate encoder cor the current response writer
	encoder := json.NewEncoder(w) // NB: Usage of encoder is 50% faster than Marshal
	// encode response
	encoder.Encode(response)
}