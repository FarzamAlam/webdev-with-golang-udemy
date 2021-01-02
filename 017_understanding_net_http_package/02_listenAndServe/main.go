package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Any code you want in this func.")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
