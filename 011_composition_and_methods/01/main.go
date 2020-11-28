package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	p1 := person{
		Name: "James Bond",
		Age:  42,
	}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatal(err)
	}
}
