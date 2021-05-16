# Use a map

Use the code in the previous folder "07_solution".

There is a map that holds all the user data.

Every time a user is created or deleted, write this map as JSON to a file.

Also, when your program starts, if there is a file with JSON data in it, load that data.

IMPORTANT:
Make sure you update your import statements to import packages from the correct location!

# Solution
1. ``models/user.go`` create 2 functions

```go
func StoreUsers(m map[string]User)  {
	f, err := os.Create("data")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(m)
}

func LoadUsers() map[string]User {
	m := make(map[string]User)
	f, err := os.Open("data")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		log.Println(err)
	}

	return m
}
```

2. ``controllers/user.go`` in function ``CreateUsers``, store map to the file

```go
// store the user
uc.session[u.Id] = u
// store user to file
models.StoreUsers(uc.session)
```

3. ``controllers/user.go`` in function ``DeleteUsers``, store map to the file

```go
// re-write data file
models.StoreUsers(uc.session)
```

4. ``main.go`` in function ``GetSession()`` is now load from data file

```go
return models.LoadUsers()
```

# Now Start the server and Test

1. create user
```shell
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
```
- output:
```shell
{"id":"f3b16a74-a3b3-40ea-ae84-718b090ee488","name":"Miss Moneypenny","gender":"female","age":27}
```
- list file if ``data`` was created
```shell
➜  02 git:(main) ✗ ls -l
total 24
-rw-r--r--  1 tony  staff  2003 May 16 14:59 README.md
drwxr-xr-x  3 tony  staff    96 May 16 14:52 controllers
-rw-r--r--  1 tony  staff   139 May 16 14:55 data
-rw-r--r--  1 tony  staff   605 May 16 14:40 main.go
drwxr-xr-x  3 tony  staff    96 May 16 14:31 models

➜  02 git:(main) ✗ cat data
{"f3b16a74-a3b3-40ea-ae84-718b090ee488":{"id":"f3b16a74-a3b3-40ea-ae84-718b090ee488","name":"Miss Moneypenny","gender":"female","age":27}}
```

2. query the user back
```shell
curl http://localhost:8080/user/f3b16a74-a3b3-40ea-ae84-718b090ee488
```
- output:
```shell
{"id":"f3b16a74-a3b3-40ea-ae84-718b090ee488","name":"Miss Moneypenny","gender":"female","age":27}
```
3. delete user
```shell
curl -X DELETE http://localhost:8080/user/f3b16a74-a3b3-40ea-ae84-718b090ee488
```
- output:
```shell
Deleted user:f3b16a74-a3b3-40ea-ae84-718b090ee488
```
