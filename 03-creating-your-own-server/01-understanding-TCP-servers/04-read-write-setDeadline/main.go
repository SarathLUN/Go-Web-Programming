package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		log.Println("Connection activated and will be expired in next 10 second.")
		fmt.Fprintln(conn, "Harry up! your connection will be expired in next 10 second.")
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard you say: %s\n", line)
	}
	defer conn.Close()

	// now we get here
	// the connection will time out
	// that breaks us out of the scanner loop
	log.Println("Connection Timeout, YOUR CODE GOT HERE!")

}
