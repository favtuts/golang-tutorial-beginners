# Golang Interfaces Explained
* https://tuts.heomi.net/golang-interfaces-explained/


# Setup project module
```bash
$ go mod init golanginterfaces
```

# What is an interface in Go?

An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.

One example of an interface type from the standard library is the [fmt.Stringer](https://pkg.go.dev/fmt/#Stringer) interface, which looks like this:
```go
type Stringer interface {
    String() string
}
```

The following `Book` type satisfies the interface because it has a `String() string` method:
```go
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}
```

The following `Count` type also satisfies the `fmt.Stringer` interface
```go
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}
```

Wherever you see declaration in Go (such as a variable, function parameter or struct field) which has an interface type, you can use an object of any type so long as it satisfies the interface.
```go
func WriteLog(s fmt.Stringer) {
    log.Print(s.String())
}
```

Because this `WriteLog()` function uses the `fmt.Stringer` interface type in its parameter declaration, we can pass in any object that satisfies the `fmt.Stringer` interface. For example, we could pass either of the `Book` and `Count` types that we made earlier to the `WriteLog()` method, and the code would work OK.

The full example code:
```go
package main

import (
    "fmt"
    "strconv"
    "log"
)

// Declare a Book type which satisfies the fmt.Stringer interface.
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
    log.Print(s.String())
}

func main() {
    // Initialize a Count object and pass it to WriteLog().
    book := Book{"Alice in Wonderland", "Lewis Carrol"}
    WriteLog(book)

    // Initialize a Count object and pass it to WriteLog().
    count := Count(3)
    WriteLog(count)
}
```

Test the `fmt.Stringer` interface demo:
```bash
$ go run .
2024/08/29 15:09:54 Book: Alice in Wonderland - Lewis Carrol
2024/08/29 15:09:54 3
```