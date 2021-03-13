package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Tony Stark"
	str := fmt.Sprintf(`
<!DOCTYPE html>
<html lang='en'>
<head>
<title>Hello World!</title>
</head>
<body>
<h1>` + name + `</h1>
</body>
</html>
`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file:", err)
	}
	defer nf.Close()
	_, err = io.Copy(nf, strings.NewReader(str))
	if err != nil {
		log.Fatalln("error writing the file:", err)
	}
}
