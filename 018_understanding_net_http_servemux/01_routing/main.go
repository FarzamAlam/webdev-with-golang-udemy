package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "Barking...")
	case "/cat":
		io.WriteString(w, "Meow...")
	default:
		io.WriteString(w, "HuHu...")
	}
}
func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
