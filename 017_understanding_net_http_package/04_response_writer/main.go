package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Farzam", "Alam")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")

}
