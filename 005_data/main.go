package main

import (
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	tpl.Execute(os.Stdout, "Mr. Alam")
}
