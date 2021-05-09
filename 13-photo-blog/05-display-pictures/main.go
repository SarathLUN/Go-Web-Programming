package main

import (
	"crypto/sha1"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	// todo: updated: add route to serve pictures
	http.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir("./uploads"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	c := getCookie(w, req)

	// handle form submit
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		log.Println(mf)
		if err != nil {
			log.Println(err)
		}
		defer func(mf multipart.File) {
			err := mf.Close()
			if err != nil {
				log.Println(err)
			}
		}(mf)

		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		log.Println(h)
		_, err = io.Copy(h, mf)
		if err != nil {
			log.Println(err)
		}
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// create new file
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		p := filepath.Join(wd, "uploads", "pics", fname)
		nf, err := os.Create(p)
		log.Println(nf)
		if err != nil {
			log.Println(err)
		}
		defer func(nf *os.File) {
			err := nf.Close()
			if err != nil {
				log.Println(err)
			}
		}(nf)
		// copy
		_, err = mf.Seek(0, 0)
		if err != nil {
			log.Println(err)
		}
		wt, err := io.Copy(nf, mf)
		log.Println(wt)
		if err != nil {
			log.Println(err)
		}
		// add filename to this user's cookie
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	// todo: updated: sliced cookie values to only send over images
	err := tpl.ExecuteTemplate(w, "index.gohtml", xs[1:])
	if err != nil {
		log.Fatalln(err)
	}
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {

	// append
	s := c.Value

	// if s not yet contain fname, add fname into
	if !strings.Contains(s, fname) {
		s += "|" + fname
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
