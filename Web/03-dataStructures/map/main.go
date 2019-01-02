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
	person := map[string]string{
		"fname" : "Yahsar", 
		"lname" : "boosharya", 
		"email" : "adsa@email.com",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", person)
	if err != nil {
		log.Fatal(err)
	}
}

