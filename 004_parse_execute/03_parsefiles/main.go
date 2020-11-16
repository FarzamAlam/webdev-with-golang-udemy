package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gohtml")
	must(err)

	must(tpl.Execute(os.Stdout, nil))

	tpl, err = tpl.ParseFiles("two.gohtml", "three.gohtml")
	must(err)

	must(tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil))
	must(tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil))
	must(tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
