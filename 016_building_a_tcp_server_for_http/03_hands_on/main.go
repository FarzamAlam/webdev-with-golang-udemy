package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error while listening : %v\n", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Printf("Error while accepting conn : %v\n", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	url := strings.Fields(ln)[1]
	msg := ""
	if m == "GET" && url == "/" {
		msg = "Home"
	} else if m == "GET" && url == "/about" {
		msg = "About"
	} else if m == "GET" && url == "/contact" {
		msg = "Contact"
	} else {
		msg = "Not Found"
	}

	respond(conn, msg)
}

func respond(conn net.Conn, msg string) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>` + msg + `</strong></body></html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
