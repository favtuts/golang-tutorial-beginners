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

You can also opt into package list mode by using `go test .`, which caches successful test results and avoids unnecessary reruns.
```bash
$ go test .
ok      github.com/favtuts/unit-testing 0.002s
$ go test .
ok      github.com/favtuts/unit-testing (cached)
$ go test .
ok      github.com/favtuts/unit-testing (cached)
```

You can run tests in a specific package by passing the relative path to the package, for example, `go test ./package-name`. Additionally, you can use `go test ./...` to run the tests for all the packages in the codebase.

If you append the `-v` flag to `go test`, the test will print out the names of all the executed test functions and the time spent for their execution. Additionally, the test displays the output of printing to the error log, for example, when you use `t.Log()` or `t.Logf()`:
```bash
$ go test -v
=== RUN   TestMultiply
--- PASS: TestMultiply (0.00s)
PASS
ok      github.com/favtuts/unit-testing 0.002s
```

Let’s cause our test to fail by changing `want` to `7`. Run `go test` once again, and inspect its output:
```bash
$ go test -v
=== RUN   TestMultiply
    integers_test.go:11: Expected '7', but got '6'
--- FAIL: TestMultiply (0.00s)
FAIL
exit status 1
FAIL    github.com/favtuts/unit-testing 0.001s
```

# Table driven tests in Go

We use table driven tests, which allow us to define all our tests cases in a slice, iterate over them, and perform comparisons to determine if the test case succeeded or failed:
```go
type testCase struct {
    arg1 int
    arg2 int
    want int
}

func TestMultiplyTableDriven(t *testing.T) {
    cases := []testCase{
        {2, 3, 6},
        {10, 5, 50},
        {-8, -3, 24},
        {0, 9, 0},
        {-7, 6, -42},
    }

    for _, tc := range cases {
        got := Multiply(tc.arg1, tc.arg2)
        if tc.want != got {
            t.Errorf("Expected '%d', but got '%d'", tc.want, got)
        }
    }
}
```

If you run the test again, it will pass successfully:
```bash
$ go test -v
=== RUN   TestMultiply
--- PASS: TestMultiply (0.00s)
=== RUN   TestMultiplyTableDriven
--- PASS: TestMultiplyTableDriven (0.00s)
PASS
ok      github.com/favtuts/unit-testing 0.002s
```