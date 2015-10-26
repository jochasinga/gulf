package views

import (
    "io/ioutil"
)


type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() error {
    filename := p.Title + ".html"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{title, body}, nil
}