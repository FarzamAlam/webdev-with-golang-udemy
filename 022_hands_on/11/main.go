package main

import (
	"bufio"
	"fmt"
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
