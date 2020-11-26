package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	g1 := struct {
		Score1 int
		Score2 int
	}{
		11,
		9,
	}
	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatal(err)
	}
}
