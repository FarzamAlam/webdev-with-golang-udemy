package main

import (
	"io"
	"net/http"
)

type hotdog int

type hotcat int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Barking...")
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Meau...")
}

func main() {
	var d hotdog
	var c hotcat
	http.Handle("/dog", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil)
}
