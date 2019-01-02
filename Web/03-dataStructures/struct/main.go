package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
		Fname string //Fname required first letter to be capitalized so that it can be exported/accessible by the html page
		Lname string
	}

func init(){
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main () {
	student := person{
		Fname: "Yashar",
		Lname: "Boosharya",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", student)
	if err != nil {
		log.Fatal(err)
	}
}

