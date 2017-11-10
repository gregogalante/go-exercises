package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	log.Printf("Server starting on port %v\n", port)

	// call the Handle method to create an handler for a filesystem server
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	// start the HTTP server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}