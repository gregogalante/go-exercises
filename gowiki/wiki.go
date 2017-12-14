package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
)

// Define Page struct
type Page struct {
	Title string
	Body []byte
}

// Define function save for the Page struct
func (p *Page) save() error {
	// define file name
	filename := "pages/" + p.Title + ".txt"
	// return the write of the file with read-write permissions for the current
	// user (0600)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Define a function to load a page from the title
func loadPage(title string) (*Page, error) {
	filename := "pages/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Define a function to manage handler general errors
func handlerErr(w http.ResponseWriter) {
	fmt.Fprintf(w, "There was an error :(")
}

// Define a function to manage handler page templates
func handlerTemplate(w http.ResponseWriter, tmp string, p *Page) {
	t, err := template.ParseFiles("templates/" + tmp + ".html")
	if err != nil {
		handlerErr(w)
		return
	}

	t.Execute(w, p)
}

// Define a view handler used to show an article
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		handlerErr(w)
		return
	}

	handlerTemplate(w, "view", p)
}

// Define the edit handler used to edit an article
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	handlerTemplate(w, "edit", p)
}

// Define main function
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}