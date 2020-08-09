// parse http request data including path[]

//

// serve handler response to url path

// listen to port

package main

import (
	"io/ioutil"
)

// define content hol;der struct

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)

	return &Page{Title: title, Body: body}
}

//he save method returns an error value because that is
// the return type of WriteFile (a standard library function that writes a byte slice to a file).
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}


func main() {

}
