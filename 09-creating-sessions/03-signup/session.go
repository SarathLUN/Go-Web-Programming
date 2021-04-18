package main

import "net/http"

func getUser(req *http.Request) user {
	var u user
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}
	//if the user exists, get user
	if userName, ok := dbSessions[c.Value]; ok {
		u = dbUsers[userName]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	userName := dbSessions[c.Value]
	_, ok := dbUsers[userName]
	return ok
}
