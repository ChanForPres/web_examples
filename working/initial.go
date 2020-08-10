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

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	//catch error
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

//he save method returns an error value because that is
// the return type of WriteFile (a standard library function that writes a byte slice to a file).
// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return ioutil.WriteFile(filename, p.Body, 0600)
// }

func (p *Page) save() error {
	filename := p.Title + ".txt"
	//	res,.. error := ioutil.WriteFile()

	return ioutil.WriteFile(filename, p.Body, 0600)
}

//
//func main() {
//	p1 := &Page{Title: "TestPage", Body: []byte("INSIDE BODY")} //create pagee
//	// save page
//	p1.save()
//	// load page
//	loadPage(p1.Title)
//	// print page body
//	fmt.Println(string(p1.Body))
//}
//
