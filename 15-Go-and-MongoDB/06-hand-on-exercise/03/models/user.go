package models

import "time"

// User capitalize to export from package
type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

// Session capitalize to export from package
type Session struct {
	UserName     string
	LastActivity time.Time
}
