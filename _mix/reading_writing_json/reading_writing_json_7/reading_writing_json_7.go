package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

var name string

// Structs:
// //////////////////////////////////////////////////////////////////////

type validationHandler struct {
	next http.Handler
}

type helloWorldHandler struct{}

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

// Helpers:
// //////////////////////////////////////////////////////////////////////

// Function used to generate a validation handler.
func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

// Function used to generate a hello world handler.
func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

// This function overrides the http ServeHTTP function to decode the request body.
func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	name = request.Name
	h.next.ServeHTTP(rw, r)
}

// This function overrides the http ServeHTTP function to code the response.
func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

// Main:
// //////////////////////////////////////////////////////////////////////

func main() {
	port := 8080
	log.Printf("Server starting on port %v\n", port)

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}