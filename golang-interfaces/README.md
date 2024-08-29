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

# Reducing boilerplate code

We have a Customer struct containing some data about a customer. In one part of our codebase we want to write the customer information to a [bytes.Buffer](https://pkg.go.dev/bytes/#Buffer), and in another part of our codebase we want to write the customer information to an [os.File](https://pkg.go.dev/os/#File) on disk. But in both cases, we want to serialize the customer struct to JSON first.

The first thing you need to know is that Go has an [io.Writer](https://pkg.go.dev/io/#Writer) interface type which looks like this:
```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

And we can leverage the fact that both `bytes.Buffer` and the `os.File` type satisfy this interface, due to them having the `bytes.Buffer.Write()` and `os.File.Write()` methods respectively.

Let’s take a look at a simple implementation:
```go
package main

import (
    "bytes"
    "encoding/json"
    "io"
    "log"
    "os"
)

// Create a Customer type
type Customer struct {
    Name string
    Age  int
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the customer struct to JSON, and if the marshal worked
// successfully, then calls the relevant io.Writer's Write() method.
func (c *Customer) WriteJSON(w io.Writer) error {
    js, err := json.Marshal(c)
    if err != nil {
        return err
    }

    _, err = w.Write(js)
    return err
}

func main() {
    // Initialize a customer struct.
    c := &Customer{Name: "Alice", Age: 21}

    // We can then call the WriteJSON method using a buffer...
    var buf bytes.Buffer
    err := c.WriteJSON(&buf)
    if err != nil {
        log.Fatal(err)
    }

    // Or using a file.
    f, err := os.Create("/tmp/customer")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()


    err = c.WriteJSON(f)
    if err != nil {
        log.Fatal(err)
    }
}
```

we can create the `Customer.WriteJSON()` method once, and we can call that method any time that we want to write to something that satisfies the `io.Writer` interface.

We can test the code:
```bash
$ go run .
$ cat /tmp/customer
{"Name":"Alice","Age":21}
```