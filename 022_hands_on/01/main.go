package main

import (
	"fmt"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Barking....")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Farzam Alam")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index.")
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
