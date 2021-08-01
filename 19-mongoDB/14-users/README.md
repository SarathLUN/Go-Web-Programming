# locking down your database

## create admin super user

```shell
use admin
```

```shell
db.createUser( \
  { \
    user: "root", \
    pwd: "root", \
    roles: [ { role: "root", db: "admin" } ] \
  } \
)
```

[built in user roles](https://docs.mongodb.com/manual/reference/built-in-roles/)

#### exit mongo & then start again
```shell
mongod --auth --port 27017 --dbpath /usr/local/var/mongodb
```
- I have also configured alias command-line: mongodAuth for quick start mongod service
```shell
mongodAuth
```
- keep mongod auth running and start another terminal for login
```shell
mongosh -u "root" -p "root" --authenticationDatabase "admin"
```
- I have configured alias command-line: mongoRoot for quick login
```shell
mongoRoot
```

#### see current user

```shell
db.runCommand({connectionStatus : 1})
```

## create regular user
Give this user readwrite permissions on the ```store``` db.

```shell
use store \
db.createUser( \
  { \
    user: "bond", \
    pwd: "moneypenny007", \
    roles: [ { role: "readWrite", db: "store" } ] \
  } \
)
```

#### login as normal user
- start another terminal

```shell
mongosh -u "bond" -p "moneypenny007" --authenticationDatabase "store"
```

#### see current user

```shell
db.runCommand({connectionStatus : 1})
```

#### lock down the database

[enable auth](https://docs.mongodb.com/master/tutorial/enable-authentication/)

[getting auth running on mongo](https://docs.mongodb.com/manual/tutorial/enable-authentication/)

#### test authorized user

```shell
mongo -u "bond" -p "moneypenny007" --authenticationDatabase "store"
```

```shell
use store
```

```shell
show collections
```

```shell
db.customers.find()
```

```shell
db.customers.insert({"role" : "double-zero", "name" : "Elon Musk", "age" : 47 })
```

#### test un-authorize user

- launch a new terminal window

```shell
mongosh
```

- should be unauthorized:
```shell
show collections
```

# drop user
- execute this by a privilege user
```shell
db.dropUser("<user name>")
```
- example:
```shell
db.dropUser("bond")
```