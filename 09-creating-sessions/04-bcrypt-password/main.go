package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

// handle templates
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// handle struct and data
type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
}

var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// process form submit
	if r.Method == http.MethodPost {
		// get form value
		un := r.FormValue("userName")
		p := r.FormValue("password")
		f := r.FormValue("firstName")
		l := r.FormValue("lastName")
		// username already token?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken!", http.StatusForbidden)
			return
		}
		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		// store user in dbUser
		xP, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Encryption error.", http.StatusInternalServerError)
			return
		}
		u := user{un, xP, f, l}
		dbUsers[un] = u
		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	// check if already logged in
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// user already logged in, so load bar
	u := getUser(r)
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}
