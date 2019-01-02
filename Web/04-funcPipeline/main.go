package main

import (
	"log"
	"os"
	"time"
	"text/template"
)


var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))
}

var fm = template.FuncMap{
	"mdy" : monthDayYear,
}

func monthDayYear(t time.Time) string{
	//predefined layout for format 01 is month 02 is day and 06 is year
	return t.Format("01-02-2006")
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}