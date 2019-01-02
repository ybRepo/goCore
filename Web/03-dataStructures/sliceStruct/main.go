package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type student struct {
		Fname string //Fname required first letter to be capitalized so that it can be exported/accessible by the html page
		Lname string
		Program string
	}

func init(){
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main () {

	student1 := student{
		Fname: "Yashar",
		Lname: "B",
		Program: "Web Dev",
	}

	student2 := student{
		Fname: "John",
		Lname: "McArthur",
		Program: "Science",
	}

	student3 := student{
		Fname: "Candice",
		Lname: "Wong",
		Program: "Math",
	}

	students := []student{student1, student2, student3}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", students)
	if err != nil {
		log.Fatal(err)
	}
}

