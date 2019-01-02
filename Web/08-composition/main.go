package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age int
}

var tpl *template.Template

func (p person) CurrentAge() int {
	return p.Age
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesAgr(x int) int {
	return x * 5
}

func init(){
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main () {
	p := person {
		Name: "James Smith",
		Age: 42,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal(err)
	}
}

