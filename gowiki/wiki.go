package main

import (
	"fmt"
	"io/ioutil"
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

// Define main function
func main() {
	title := "How to be Gregorio"
	// create a page and save it
	p1 := &Page{Title: title, Body: []byte("It's not possible")}
	p1.save()
	// load the page
	p2, err := loadPage(title)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p2.Body)
	}
}