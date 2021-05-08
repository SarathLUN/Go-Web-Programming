package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)

	c := getCookie(w, req)
	c = appendValue(w, c)
	xs := strings.Split(c.Value, "|")

	err := tpl.ExecuteTemplate(w, "index.gohtml", xs)
	if err != nil {
		log.Fatalln(err)
	}
}

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// values
	p1 := "disneyland.jpg"
	p2 := "atbeach.jpg"
	p3 := "hollywood.jpg"
	// append
	s := c.Value

	// if s not yet contain p1, add p1 into
	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	// if s not yet contain p2, add p2 into
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	// if s not yet contain p3, add p3 into
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}
	// set back to cookie
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
