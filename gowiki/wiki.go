package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define Page struct
type Page struct {
	Title string
	Body []byte
}

// Define function save for the Page struct
func (p *Page) save() error {
	// define file name
	filename := p.Title + ".txt"
	// return the write of the file with read-write permissions for the current
	// user (0600)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Define a function to load a page from the title
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Define a view handler used to show an article
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "There was an error :(")
	} else {
		fmt.Println("Request for", title)
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	}
}

// Define the edit handler used to edit an article
// TODO

// Define main function
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}