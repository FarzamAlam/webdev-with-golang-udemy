package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/bar", http.HandlerFunc(bar))
	http.Handle("/me", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bar ran")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalf("Error while parsing : %v\n", err)
	}
	err = tpl.ExecuteTemplate(w, "template.gohtml", "Farzi")
	if err != nil {
		log.Fatalf("Error while executing the template : %v\n", err)
	}
}
