# Install Mongo

follow this official [documents](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/)

```shell
brew tap mongodb/brew
```
```shell
brew install mongodb-community
```
start mongoDB service
```shell
brew services start mongodb-community
```
show status of running service
```shell
brew services list
```
output:
```shell
Name              Status  User Plist
mongodb-community started tony /Users/tony/Library/LaunchAgents/homebrew.mxcl.mongodb-community.plist
```
login to mongodb server, by default:
- allowed only localhost
- default port = 27017
- no username/password required
```shell
mongo
```
output:
```shell
MongoDB shell version v4.4.5
connecting to: mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("ef036a17-8247-4476-9f2f-a3319808e46c") }
MongoDB server version: 4.4.5
---
The server generated these startup warnings when booting:
        2021-05-15T17:42:04.753+07:00: Access control is not enabled for the database. Read and write access to data and configuration is unrestricted
        2021-05-15T17:42:04.753+07:00: Soft rlimits too low
        2021-05-15T17:42:04.753+07:00:         currentValue: 256
        2021-05-15T17:42:04.753+07:00:         recommendedMinimum: 64000
---
---
        Enable MongoDB's free cloud-based monitoring service, which will then receive and display
        metrics about your deployment (disk utilization, CPU, operation statistics, etc).

        The monitoring data will be available on a MongoDB website with a unique URL accessible to you
        and anyone you share the URL with. MongoDB may use this information to make product
        improvements and to suggest MongoDB products and deployment options to you.

        To enable free monitoring, run the following command: db.enableFreeMonitoring()
        To permanently disable this reminder, run the following command: db.disableFreeMonitoring()
---
>
```

stop mongoDB service
```shell
brew services stop mongodb-community
```

# Go get driver for mongo
- I will use this driver for learning purpose (to follow along)
```
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
```

- for production, we should use official driver [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)

# In this step:

Don't run this code

Just making updates - a several step process.

We will need a mongo session to use in the CRUD methods.

We need our UserController to have access to a mongo session.

Let's add this to controllers/user.go

```
UserController struct {  
    session *mgo.Session
}
```

And now add this to controllers/user.go

```
func NewUserController(s *mgo.Session) *UserController {  
    return &UserController{s}
}
```

And now add this to main.go

```
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
```

and this

```
uc := controllers.NewUserController(getSession())  
```

1. Enter this at the terminal

```
curl http://localhost:8080/user/1
```