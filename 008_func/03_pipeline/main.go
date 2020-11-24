package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

func double(n float64) float64 {
	return n * 2
}

func square(n float64) float64 {
	return n * n
}

var fm = template.FuncMap{
	"dbl":  double,
	"sqr":  square,
	"sqrt": math.Sqrt,
}

func main() {
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3.0)
	if err != nil {
		log.Fatal(err)
	}
}
