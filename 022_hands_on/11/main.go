package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	var i int
	var rMethod, rURI string
	scan := bufio.NewScanner(c)
	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS.")
			break
		}
		i++
	}
	switch {
	case rMethod == "GET" && rURI == "/":
		handleIndex(c)
	case rMethod == "GET" && rURI == "/apply":
		handleApply(c)
	case rMethod == "POST" && rURI == "/apply":
		handleApplyPost(c)
	default:
		handleDefault(c)
	}
}

func handleIndex(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>		
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	io.WriteString(c, body)
}
