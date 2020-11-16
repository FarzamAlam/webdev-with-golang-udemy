package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	must(tpl.Execute(os.Stdout, nil))
	must(tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil))
	must(tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil))
	must(tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
