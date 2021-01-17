package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error while listening :%v\n", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error while accepting :%v\n", err)
			continue
		}
		io.WriteString(conn, "I know you are connected.")
		conn.Close()

	}
}
