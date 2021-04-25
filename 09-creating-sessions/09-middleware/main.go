package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
	Role      string
}

type session struct {
	username     string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", authorized(logout))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func authorized(h http.HandlerFunc) http.HandlerFunc {
	log.Println("running func authorized()")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !alreadyLoggedIn(w, r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return // don't call original handler
		}
		h.ServeHTTP(w, r)
	})
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func cleanSessions() {
	log.Println("BEFORE CLEAN")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	log.Println("AFTER CLEAN")
	showSessions()
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	// process form submit
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	//process form submit
	if req.Method == http.MethodPost {
		// get form value
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username already taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken!", http.StatusForbidden)
			return
		}
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Encrypt password error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l, r}
		dbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar!", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.username]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSessions() // for demonstration purpose
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func showSessions() {
	log.Println("******")
	for k, v := range dbSessions {
		log.Println(k, v.username)
	}
	log.Println("======")
}

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	// if the user already existed, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.username]
	}
	return u
}
