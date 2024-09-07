# Advanced unit testing patterns in Go
* https://tuts.heomi.net/advanced-unit-testing-patterns-in-go/

# Init a Go project

Locate the project directory
```bash
$ mkdir unit-testing-advanced
$ cd unit-testing-advanced
```

Then create the module
```bash
$ go mod init github.com/favtuts/unit-testing
```


# Testing HTTP handlers

## Method 1: Checking status code

Let’s consider a basic test that checks the status code of the following HTTP handler:
```go
func index(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}
```

The `index()` handler above should return a 200 OK response for every request. Let’s verify the handler’s response with the following test:
```go
func TestIndexHandler(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodGet, "/", nil)

    index(w, r)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status: %d, but got: %d", http.StatusOK, w.Code)
    }
}
```

Run the test:
```bash
$ go test -v
=== RUN   TestIndexHandler
--- PASS: TestIndexHandler (0.00s)
PASS
ok      github.com/favtuts/unit-testing 0.002s
```