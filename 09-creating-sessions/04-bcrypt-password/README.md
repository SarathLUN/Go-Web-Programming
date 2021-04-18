You should never store a password without encrypting it.

We will use ```bcrypt``` to encrypt our passwords.

# Step 1:
You will need to go get this package and import it:
```shell
go get golang.org/x/crypto/bcrypt
```

```go
import (
    uuid "github.com/satori/go.uuid"
    "golang.org/x/crypto/bcrypt"
    "html/template"
    "net/http"
)
```

# Step 2:
Change data type of field:```password``` in struct:```user``` to be the ```[]byte```

```go
type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
}
```

# Step 3:
Use bcrypt to encrypt your password before storing it.
```go
xP, err := bcrypt.GenerateFromPassword([]byte(p),bcrypt.MinCost)
```

# Step 4: Testing
```shell
go run .
```