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

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	dz := doubleZero{
		person{
			Name: "Ian Fleming",
			Age:  56,
		},
		true,
	}
	err := tpl.Execute(os.Stdout, dz)
	if err != nil {
		log.Fatal(err)
	}
}
