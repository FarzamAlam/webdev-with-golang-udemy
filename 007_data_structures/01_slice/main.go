package main

import (
	"os"
	"text/template"
)

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	tpl.Execute(os.Stdout, sages)
}
