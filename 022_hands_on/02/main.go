package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Barking...")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalf("Error while parsing template : %v\n", err)
	}
	err = tpl.ExecuteTemplate(w, "template.gohtml", "Farzi")
	if err != nil {
		log.Fatalf("Error while executing the template: %v\n", err)
	}
}
