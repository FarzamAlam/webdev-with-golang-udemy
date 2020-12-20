package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	var hotels []hotel

	holidayInn := hotel{
		Name:    "Holiday Inn",
		Address: "1, Baker street",
		City:    "Silcon Valley",
		Zip:     "123",
		Region:  "Southern",
	}
	clarksInn := hotel{
		Name:    "Clarks Inn",
		Address: "2, Baker street",
		City:    "Silcon Valley",
		Zip:     "123",
		Region:  "Central",
	}
	renaissance := hotel{
		Name:    "Renaissance",
		Address: "3, Baker street",
		City:    "Silcon Valley",
		Zip:     "123",
		Region:  "Northern",
	}
	hotels = append(hotels, holidayInn, clarksInn, renaissance)

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}

}
