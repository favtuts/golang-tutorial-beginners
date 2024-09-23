# Using a PostgreSQL Database in Go (in Production) 
* https://tuts.heomi.net/using-a-postgresql-database-in-go-in-production/

# Init a Go project

Locate the project directory
```bash
$ mkdir structs-interfaces
$ cd structs-interfaces
```

Then create the module
```bash
$ go mod init github.com/favtuts/postgresql-database
```

Create main.go with the content:
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, Go")
}
```

Run the main function
```bash
$ go run .
Hello, Go
```