package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", trianga)
	http.HandleFunc("/trianga.jpg", triangaPic)
	http.ListenAndServe(":8080", nil)
}

func trianga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src= "trianga.jpg">`)
}

func triangaPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("trianga.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
