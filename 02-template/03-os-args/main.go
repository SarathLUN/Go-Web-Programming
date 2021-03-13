package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
		name +
		`</h1>
		</body>
		</html>
`)
	nf, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
