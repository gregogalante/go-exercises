package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	"errors"
)

// Define a global template variable to cache templates
var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

// Define a global valid path variable to validate the filename of new pages
// to avoid the possibility to write everywhere on the server
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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

// Define a function to manage handler page templates
func handlerTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Define a function to get the title of the page from the request
func handlerGetTitle(r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)

	if m == nil {
		return "", errors.New("Invalid Page Title")
	}

	return m[2], nil
}

// Define a view handler used to show an article
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}

	handlerTemplate(w, "view", p)
}

// Define the edit handler used to edit an article
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	handlerTemplate(w, "edit", p)
}

// Define the save handler used to save an article
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

// Define a function that work as a middleware for all requests
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		title, err := handlerGetTitle(r)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, title)
	}
}

// Define main function
func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}