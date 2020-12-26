package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error while listening %v\n", err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Printf("Error while accepting connection : %v\n", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "We have recieved : %v\n", ln)
	}
	defer conn.Close()
}
