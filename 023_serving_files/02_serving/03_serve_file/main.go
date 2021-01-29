package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", trianga)
	http.HandleFunc("/trianga.jpg", triangaPic)
	http.ListenAndServe(":8080", nil)
}

func trianga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="trianga.jpg">`)
}

func triangaPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "trianga.jpg")
}
