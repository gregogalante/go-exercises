package main

import (
	"fmt"
	"log"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	port := 8080
	portString := fmt.Sprintf(":%v", port)

	log.Printf("Server starting on port %v\n", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	http.ListenAndServe(portString, mux)
}
