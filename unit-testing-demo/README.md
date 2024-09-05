# A deep dive into unit testing in Go
* https://tuts.heomi.net/a-deep-dive-into-unit-testing-in-go/

# Init a Go project

Locate the project directory
```bash
$ mkdir unit-testing-demo
$ cd unit-testing-demo
```

Then create the module
```bash
$ go mod init github.com/favtuts/unit-testing
```

# Writing your first test in Go

First, create a file called `integers.go` and add the following code:
```go
// integers.go
package main

import (
    "fmt"
)

// Multiply returns the product of two integers
func Multiply(a, b int) int {
    return a * b
}
```

Let’s write a test to verify that the `Multiply()` function works correctly. In the current directory, create a file called `integers_test.go` and add the following code to it:
```go
// integers_test.go
package main

import "testing"

func TestMultiply(t *testing.T) {
    got := Multiply(2, 3)
    want := 6

    if want != got {
        t.Errorf("Expected '%d', but got '%d'", want, got)
    }
}
```

# Running Go tests

Now, let’s use the `go test` command to run our test in the terminal. The `go test` command compiles the sources, files, and tests found in the current directory, then runs the resulting test binary. When testing is done, a summary of the test, either `PASS` or `FAIL`, will be printed to the console, as seen in the code block below:

```bash
$ go test
PASS
ok      github.com/favtuts/unit-testing 0.001s
```

When you use `go test` as above, caching is disabled, so the tests are executed every time.