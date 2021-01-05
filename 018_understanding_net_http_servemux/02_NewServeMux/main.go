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
	io.WriteString(w, "Meao...")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
