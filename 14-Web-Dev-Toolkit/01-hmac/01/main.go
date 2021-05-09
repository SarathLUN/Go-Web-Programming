package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
)

func main() {
	c := getHash("test@test.com")
	fmt.Println(c)
	c = getHash("test@example.com")
	fmt.Println(c)
}

func getHash(s string) string {
	h := hmac.New(sha256.New, []byte("our-private-key"))
	_, err := io.WriteString(h, s)
	if err != nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
