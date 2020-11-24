package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}
