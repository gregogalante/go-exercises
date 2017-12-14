package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

// Responses:
// //////////////////////////////////////////////////////////////////////

type messageResponse struct {
	Message string `json:"message"`
}

// Requests:
// //////////////////////////////////////////////////////////////////////

type helloUserRequest struct {
	Name string `json:"name"`
}

// Handlers:
// //////////////////////////////////////////////////////////////////////

func helloUserHandler(w http.ResponseWriter, r *http.Request) {
	// read request json
	var request helloUserRequest
	var result bool
	var name string
	if r.Method == "GET" {
		result = true
		name = r.URL.Query().Get("name")
	} else if r.Method == "POST" {
		result = decodeRequestBody(w, r, &request)
		name = request.Name
	}
	if !result { return }
	
	// create an interface used for the Marshal function
	response := messageResponse{Message: "Hello " + name}
	// generate encoder cor the current response writer
	encoder := json.NewEncoder(w)
	// encode response
	encoder.Encode(response)
}

// Helpers:
// //////////////////////////////////////////////////////////////////////

func decodeRequestBody(w http.ResponseWriter, r *http.Request, request interface{}) bool {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return false
	}

	return true
}

// Main:
// //////////////////////////////////////////////////////////////////////

func main() {
	port := 8080
	log.Printf("Server starting on port %v\n", port)

	// call the HandleFunc method which creates a Handler type
	http.HandleFunc("/hellouser", helloUserHandler)
	
	// start the HTTP server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
