package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main () {
	name := []string{"Yahsar", "Frank", "James", "Heather"}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", name)
	if err != nil {
		log.Fatal(err)
	}
}

