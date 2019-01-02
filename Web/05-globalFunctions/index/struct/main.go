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

	// create the new data struct combining the slice and a string and defining its values. 
	//Field Names are capitalized so that they can be exported.
	data := struct{
		Words []string
		Name string
	}{
		xs,
		"Yashar",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	} 
}