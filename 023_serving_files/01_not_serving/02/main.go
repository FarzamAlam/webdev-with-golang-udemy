package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveImage)
	http.ListenAndServe(":8080", nil)
}
func serveImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<!--image doesn't serve-->
	<img src="trianga.jpg">
	`)
}
