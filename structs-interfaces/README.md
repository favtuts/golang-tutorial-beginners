# Exploring structs and interfaces in Go
* https://tuts.heomi.net/exploring-structs-and-interfaces-in-go/

# Init a Go project

Locate the project directory
```bash
$ mkdir structs-interfaces
$ cd structs-interfaces
```

Then create the module
```bash
$ go mod init github.com/favtuts/structs-interfaces
```

Create main.go with the content:
```go
package main  // specifies the package name

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