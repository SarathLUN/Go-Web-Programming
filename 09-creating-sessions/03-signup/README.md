# Separate Go Files:
In this example we learn how we can separate of concern on multiple Go file.

- file ```main.go```: is still the entry point of the package main
- file ```session.go```: extended functions.

Now we have multiple Go files in our package, and we can't execute our package with ```go run main.go``` anymore.
To execute this package we need to use command:

```shell
go run .
```
Or
```shell
go run *.go
```

# Step 1

Created ```func getUser``` and put it in a new file, session.go. This refactor allows us to use the same code in index and bar.

```shell
func getUser(w http.ResponseWriter, req *http.Request) user 
```

# Step 2

Created ```func signup``` and removed the signup code from ```func index```. A new field for password was added to the user struct.

```shell
func signup(w http.ResponseWriter, req *http.Request)
```

# Step 3

Created ```func alreadyLoggedIn``` and put it on the session.go page. This refactor allows us to use the same code in bar and signup.

```shell
func alreadyLoggedIn(req *http.Request) bool
```