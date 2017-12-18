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
	// get data and error from the Marshal function
	data, err := json.Marshal(response)
	// check if an error is present
	if err != nil {
		panic("Ooops")
	}
	// return the response as a json string
	fmt.Fprint(w, string(data))
}