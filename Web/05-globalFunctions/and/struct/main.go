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

type user struct{
	Name string
	Occupation string
	Program string
	Grade int
}

func main() {

	u1 := user{
		Name: "Yashar",
		Occupation: "student",
		Program: "Web Dev",
		Grade: 80,
	}

	u2 := user{
		Name: "Tim",
		Occupation: "student",
		Program: "Web Dev",
		Grade: 90,
	}

	u3 := user{
		Name: "",
		Occupation: "",
		Program: "Science",
		Grade: 50,
	}

	u4 := user{
		Name: "John",
		Occupation: "student",
		Program: "Web Dev",
		Grade: 40,
	}

	data := []user{u1, u2, u3, u4}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	} 
}