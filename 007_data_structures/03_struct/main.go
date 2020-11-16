package main

import (
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	tpl.Execute(os.Stdout, buddha)
}
