# Error handling in Go: Best practices 
* https://tuts.heomi.net/error-handling-in-go-best-practices/

# Init a Go project

Locate the project directory
```bash
$ mkdir error-handling
$ cd error-handling
```

Then create the module
```bash
$ go mod init github.com/favtuts/error-handling
```

# Handling errors through multiple return values

There is a `error` interface type declared in [Go's built-in package](https://pkg.go.dev/builtin@go1.19.2), and its zero value is `nil`.
```go
type error interface {
   Error() string
}
```
Usually, returning an `error` means that there is a problem, and returning `nil` means there were no errors:
```go
func iterate(x, y int) (int, error) {

}
result, err := iterate(x, y)
if err != nil {
    // handle the error appropriately
} else {
    // you're good to go
}
```
Whenever the `iterate` function is called and `err` is not equal to `nil`, the `error` returned should be handled appropriately.

The only drawback with handling errors this way is that there’s no enforcement from Go’s compiler. You have to decide how the function you created returns the error.

You could define an `error` struct and place it in the position of the returned values. One way to do this is by using the built-in `errorString` struct. You can also find the code below in [Go’s source code](https://go.dev/src/errors/errors.go):
```go
package errors

func New(text string) error {
    return &errorString {
        text
    }
}

type errorString struct {
    s string
}

func(e * errorString) Error() string {
    return e.s
}
```

In the code sample above, `errorString` embeds a `string`, which is returned by the Error method. To create a custom error, you’ll have to define your `error` struct and use [method sets](https://golang.org/ref/spec#Method_sets) to associate a function to your struct:
```go
package main

// Define an error struct
type CustomError struct {
    msg string
}
// Create a function Error() string and associate it to the struct.
func(error * CustomError) Error() string {
    return error.msg
}
// Then create an error object using MyError struct.
func CustomErrorInstance() error {
    return &CustomError {
        "File type not supported"
    }
}
```

In the `main.go` you can test the CustomError object
```go
package main

import "fmt"

func main() {
	err := CustomErrorInstance()
	fmt.Printf("Error: %+v", err)
}
```

If you run the main.go
```bash
$ go run main.go
```
You may see the error: [/main.go:6:9: undefined: CustomErrorInstance](https://stackoverflow.com/questions/28153203/undefined-function-declared-in-another-file)

To fix it, let's run:
```bash
$ go run .
```

This `go run .` will run all of your files. The entry point is the function `main()` which has to be unique to the `main` package.


The newly created custom error can then be restructured to use the built-in `error` struct:
```go
import "errors"
func CustomeErrorInstance() error {
    return errors.New("File type not supported")
}
```
One limitation of the built-in `error` struct is that it does not come with stack traces, making it very difficult to locate where an error occurred. The error could pass through a number of functions before being printed out.

To handle this, you could install the [pkg/errors](https://github.com/pkg/errors/blob/master/errors.go) package, which provides basic error handling primitives like stack trace recording, error wrapping, unwrapping, and formatting. To install this package, run the command below in your terminal:
```bash
$ go get github.com/pkg/errors
```

When you need to add stack traces or any other information to make debugging your errors easier, use the `New` or `Errorf` functions to provide errors that record your stack trace. `Errorf` implements the `fmt.Formatter` interface, which lets you format your errors using the `fmt` package runes, `%s`, `%v`, and `%+v`:
```go
import(
    "github.com/pkg/errors"
    "fmt"
)
func X() error {
    return errors.Errorf("Could not write to file")
}

func CustomErrorPkgErrors() {
    return X()
}

func main() {
    fmt.Printf("Error: %+v", CustomErrorPkgErrors())
}
```

To print stack traces instead of a plain error message, you have to use `%+v` instead of `%v` in the format pattern. The stack traces will look similar to the code sample below:
```bash
Error: Could not write to file (pkg.errors)
main.X
	/home/tvt/go-projects/golang-tutorial-beginners/error-handling/pkg_errors.go:8
main.CustomErrorPkgErrors
	/home/tvt/go-projects/golang-tutorial-beginners/error-handling/pkg_errors.go:12
main.main
	/home/tvt/go-projects/golang-tutorial-beginners/error-handling/main.go:8
runtime.main
	/home/tvt/.goenv/versions/1.22.4/src/runtime/proc.go:271
runtime.goexit
	/home/tvt/.goenv/versions/1.22.4/src/runtime/asm_amd64.s:1695
```