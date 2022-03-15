package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	p := &Page{"ssss", []byte("sknaklsndalmd")}
	p.Save()
	k := new(Page)
	k.Load(p.title)
	fmt.Println(string(k.body))
}

type Page struct {
	title string
	body  []byte
}

func (p *Page) Save() error {
	return ioutil.WriteFile(p.title, p.body, 0666)
}

func (p *Page) Load(title string) (err error) {
	p.title = title
	p.body, err = ioutil.ReadFile(title)
	return
}
