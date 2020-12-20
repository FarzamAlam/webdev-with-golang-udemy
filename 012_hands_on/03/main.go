package main

import (
	"log"
	"os"
	"text/template"
)

type breakfast struct {
	Tea       string
	Coffee    string
	Sandwitch string
}

type lunch struct {
	Thali        string
	PudiSabzi    string
	CholeBhature string
}

type dinner struct {
	Thali       string
	RajmaChawal string
	Biryani     string
}

type menu struct {
	Breakfast breakfast
	Lunch     lunch
	Dinner    dinner
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	menu := menu{
		Breakfast: breakfast{
			Tea:       "10",
			Coffee:    "20",
			Sandwitch: "30",
		},
		Lunch: lunch{
			Thali:        "70",
			PudiSabzi:    "30",
			CholeBhature: "30",
		},
		Dinner: dinner{
			Thali:       "100",
			RajmaChawal: "50",
			Biryani:     "80",
		},
	}
	err := tpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatalln(err)
	}
}
