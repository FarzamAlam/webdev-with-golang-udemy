package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error while listening :%v\n", err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Printf("Error while accepting : %v\n", err)
			continue
		}
		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("THIS IS THE END OF THE HTTP REQUEST.")
				break
			}
		}
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")
	}

}
