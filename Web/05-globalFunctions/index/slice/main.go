package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}


func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five",}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	} 
}