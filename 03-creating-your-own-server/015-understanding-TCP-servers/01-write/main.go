package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	// defer lis.Close() // this will warning of none handle err
	// so we make use of anonymous func to handle err
	defer func() {
		err = lis.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		conn, err := lis.Accept()
		if err != nil {
			// log.Fatalln(err) // Fatal will terminate the program
			// so we use Println with continue instead
			log.Println(err)
			continue
		}

		// now we try to write a string into connection
		// in this case, we just ignore the return (int, err)
		_, _ = io.WriteString(conn, "\nHello from TCP server\n")
		// now we read from conn
		log.Println(conn)
		_, _ = fmt.Fprintln(conn, "How is your day?")
		_, _ = fmt.Fprintf(conn, "%v", "Well, I hope!\n")

		// always close the connection after all
		err = conn.Close()
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
