package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("templates/*gohtml")
	must(err)

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
