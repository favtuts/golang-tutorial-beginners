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

# Defer, panic, and recover

Although Go doesn’t have exceptions, it has a similar type of mechanism known as [defer, panic, and recover](https://blog.golang.org/defer-panic-and-recover). Go’s ideology is that adding exceptions like the `try/catch/finally` statement in JavaScript would result in complex code and encourage programmers to label too many basic errors, like failing to open a file, as exceptional.

`Defer` is a language mechanism that puts your function call into a stack. Each deferred function is executed in reverse order when the host function finishes, regardless of whether a `panic` is called or not. The `defer` mechanism is very useful for cleaning up resources:
```go
package main

import (
        "fmt"
)

func A() {
        defer fmt.Println("Keep calm!")
        B()
}
func B() {
        defer fmt.Println("Else...")
        C()
}
func C() {
        defer fmt.Println("Turn on the air conditioner...")
        D()
}
func D() {
        defer fmt.Println("If it's more than 30 degrees...")
}
func main() {
        A()
}
```

The code above would compile as follows:
```bash
If it's more than 30 degrees...
Turn on the air conditioner...
Else...
Keep calm!
```

`panic` is a built-in function that stops the normal execution flow. When you call `panic` in your code, it means you’ve decided that your caller can’t solve the problem. Therefore, you should use `panic` only in rare cases where it’s not safe for your code or anyone integrating your code to continue at that point.

The code sample below demonstrates how `panic` works:
```go
package main

import (
        "errors"
        "fmt"
)

func A() {
        defer fmt.Println("Then we can't save the earth!")
        B()
}
func B() {
        defer fmt.Println("And if it keeps getting hotter...")
        C()
}
func C() {
        defer fmt.Println("Turn on the air conditioner...")
        Break()
}
func Break() {
        defer fmt.Println("If it's more than 30 degrees...")
        panic(errors.New("Global Warming!!!"))

}
func main() {
        A()
}
```

The sample above would compile as follows:
```bash
If it's more than 30 degrees...
Turn on the air conditioner...
And if it keeps getting hotter...
Then we can't save the earth!
panic: Global Warming!!!

goroutine 1 [running]:
main.Break()
        /home/tvt/go-projects/golang-tutorial-beginners/error-handling/panic_example.go:22 +0x90
main.C2()
        /home/tvt/go-projects/golang-tutorial-beginners/error-handling/panic_example.go:18 +0x65
main.B2()
        /home/tvt/go-projects/golang-tutorial-beginners/error-handling/panic_example.go:14 +0x65
main.A2()
        /home/tvt/go-projects/golang-tutorial-beginners/error-handling/panic_example.go:10 +0x65
main.main()
        /home/tvt/go-projects/golang-tutorial-beginners/error-handling/main.go:16 +0x145
exit status 2
```

As shown above, when `panic` is used and not handled, the execution flow stops, all deferred functions are executed in reverse order, and stack traces are printed.

You can use the built-in `recover` function to handle `panic` and return the values passed from a panic call. `recover` must always be called in a `defer` function, otherwise, it will return `nil`:
```go
package main

import (
        "errors"
        "fmt"
)

func A() {
        defer fmt.Println("Then we can't save the earth!")
        defer func() {
                if x := recover(); x != nil {
                        fmt.Printf("Panic: %+v\n", x)
                }
        }()
        B()
}
func B() {
        defer fmt.Println("And if it keeps getting hotter...")
        C()
}
func C() {
        defer fmt.Println("Turn on the air conditioner...")
        Break()
}
func Break() {
        defer fmt.Println("If it's more than 30 degrees...")
        panic(errors.New("Global Warming!!!"))

}
func main() {
        A()
}
```

As you can see in the code sample above, recover prevents the entire execution flow from coming to a halt. We added in a panic function, so the compiler would return the following:
```bash
If it's more than 30 degrees...
Turn on the air conditioner...
And if it keeps getting hotter...
Panic: Global Warming!!!
Then we can't save the earth!
```

To report an error as a return value, you have to call the `recover` function in the same goroutine as where the `panic` function is called, retrieve an error struct from the `recover` function, and pass it to a variable:
```go
package main

import (
        "errors"
        "fmt"
)

func saveEarth() (err error) {

        defer func() {
                if r := recover(); r != nil {
                        err = r.(error)
                }
        }()
        TooLate()
        return
}
func TooLate() {
        A()
        panic(errors.New("Then there's nothing we can do"))
}

func A() {
        defer fmt.Println("If it's more than 100 degrees...")
}
func main() {
        err := saveEarth()
        fmt.Println(err)
}
```

# Error wrapping

Previously, error wrapping in Go was only accessible via packages like `pkg/errors`. However, [Go v1.13 introduced support for error wrapping](https://go.dev/doc/go1.13).

To create wrapped errors, `fmt.Errorf` has a `%w` verb, and for inspecting and unwrapping errors, a couple of functions have been added to the `error` package.

`errors.Unwrap` basically inspects and exposes the underlying errors in a program. It returns the result of calling the `Unwrap` method on `Err` if `Err`’s type contains an `Unwrap` method returning an error. Otherwise, `Unwrap` returns nil:
```go
package errors

type Wrapper interface{
    Unwrap() error
}
```

Below is an example implementation of the `Unwrap` method:
```go
func(e*PathError) Unwrap() error{
    return e.Err
}
```

With the `errors.Is` function, you can compare an error value against the sentinel value. Instead of comparing the sentinel value to one error, this function compares it to every error in the error chain. It also implements an `Is` method on an error so that an error can post itself as a sentinel even though it’s not a sentinel value:
```go
func Is(err, target error) bool
```

In the basic implementation above, `Is` checks and reports if `err` or any of the `errors` in its chain are equal to the target, the sentinel value.

The `errors.As` function provides a way to cast to a specific error type. It looks for the first error in the error chain that matches the sentinel value, and if found, it sets the sentinel value to that error value, returning `true`:
```go
package main

import (
        "errors"
        "fmt"
        "os"
)

func main() {
        if _, err := os.Open("non-existing"); err != nil {
                var pathError *os.PathError
                if errors.As(err, &pathError) {
                        fmt.Println("Failed at path:", pathError.Path)
                } else {
                        fmt.Println(err)
                }
        }

}
```