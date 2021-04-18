package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

//handle the templates
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//handle data & struct
type user struct {
	UserName  string
	FirstName string
	LastName  string
}

var dbUsers = map[string]user{}     // userID, user
var dbSession = map[string]string{} // sessionID, userID(FK)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	// get cookie
	cookie, err := r.Cookie("session")
	// cookie not exist, redirect to index
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// cookie exist, get the userName from dbSession
	userName, ok := dbSession[cookie.Value]
	// can't find, redirect to index
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// found the username, then lookup user info
	u := dbUsers[userName]
	// execute the template to display user info
	err = tpl.ExecuteTemplate(w, "bar.gohtml", u)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// get the cookie
	cookie, err := r.Cookie("session")
	// cookie not exist will return error
	if err != nil {
		sID := uuid.NewV4()
		// create new cookie
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		// set new cookie
		http.SetCookie(w, cookie)
	}
	// cookie exist, get the user
	var u user
	if userName, ok := dbSession[cookie.Value]; ok {
		u = dbUsers[userName]
	}

	// process form submission
	if r.Method == http.MethodPost {
		// get value from input
		userName := r.FormValue("userName")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		// create user
		u = user{userName, firstName, lastName}
		// override dbSession via cookie.Value
		dbSession[cookie.Value] = userName
		// override user info in dbUser via userName
		dbUsers[userName] = u
	}

	// execute the template
	err = tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		log.Fatalln(err)
	}
}
