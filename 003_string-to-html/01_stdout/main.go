package main

import "fmt"

func main() {
	name := "Farzam Alam"

	tpl := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Hello World!</title>
		</head>
		<body>
			<h1> ` + name + `</h1>
		</body>
	</html>
	`
	fmt.Println(tpl)
}
