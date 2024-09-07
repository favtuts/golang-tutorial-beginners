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
$ go test -v -run=TestIndexHandler
=== RUN   TestIndexHandler
--- PASS: TestIndexHandler (0.00s)
PASS
ok      github.com/favtuts/unit-testing 0.002s
```

# Method 2: External dependencies

Now, let’s consider another common scenario in which our HTTP handler has a dependency on an external service:
```go
func getJoke(w http.ResponseWriter, r *http.Request) {
    u, err := url.Parse(r.URL.String())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    jokeId := u.Query().Get("id")
    if jokeId == "" {
        http.Error(w, "Joke ID cannot be empty", http.StatusBadRequest)
        return
    }

    endpoint := "https://icanhazdadjoke.com/j/" + jokeId

    client := http.Client{
        Timeout: 10 * time.Second,
    }

    req, err := http.NewRequest(http.MethodGet, endpoint, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    req.Header.Set("Accept", "text/plain")

    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    defer resp.Body.Close()

    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if resp.StatusCode != http.StatusOK {
        http.Error(w, string(b), resp.StatusCode)
        return
    }

    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    w.Write(b)
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/joke", getJoke)

    http.ListenAndServe(":1212", mux)
}
```

In the code block above, the `getJoke` handler expects an `id` query parameter, which it uses to fetch a joke from the [Random dad joke API](https://icanhazdadjoke.com/).

Let’s write a test for this handler:
```go
func TestGetJokeHandler(t *testing.T) {
    table := []struct {
        id         string
        statusCode int
        body       string
    }{
        {"R7UfaahVfFd", 200, "My dog used to chase people on a bike a lot. It got so bad I had to take his bike away."},
        {"173782", 404, `Joke with id "173782" not found`},
        {"", 400, "Joke ID cannot be empty"},
    }

    for _, v := range table {
        t.Run(v.id, func(t *testing.T) {
            w := httptest.NewRecorder()
            r := httptest.NewRequest(http.MethodGet, "/joke?id="+v.id, nil)

            getJoke(w, r)

            if w.Code != v.statusCode {
                t.Fatalf("Expected status code: %d, but got: %d", v.statusCode, w.Code)
            }

            body := strings.TrimSpace(w.Body.String())

            if body != v.body {
                t.Fatalf("Expected body to be: '%s', but got: '%s'", v.body, body)
            }
        })
    }
}
```

Run the test:
```bash
$ go test -v -run=TestGetJokeHandler
=== RUN   TestGetJokeHandler
=== RUN   TestGetJokeHandler/R7UfaahVfFd
=== RUN   TestGetJokeHandler/173782
=== RUN   TestGetJokeHandler/#00
--- PASS: TestGetJokeHandler (0.80s)
    --- PASS: TestGetJokeHandler/R7UfaahVfFd (0.53s)
    --- PASS: TestGetJokeHandler/173782 (0.27s)
    --- PASS: TestGetJokeHandler/#00 (0.00s)
PASS
ok      github.com/favtuts/unit-testing 0.802s
```

Note that the test in the code block above makes HTTP requests to the real API. Doing so affects the dependencies of the code being tested, which is bad practice for unit testing code. Instead, we should mock the HTTP client. We have several different methods for mocking in Go.