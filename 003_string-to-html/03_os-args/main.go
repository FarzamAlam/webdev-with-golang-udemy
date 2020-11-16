package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	tpl := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>Hello world!</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error while creating file : ", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
