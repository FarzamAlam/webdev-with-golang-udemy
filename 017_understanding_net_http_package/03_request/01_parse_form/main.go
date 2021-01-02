package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error while serving : %v\n", err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
