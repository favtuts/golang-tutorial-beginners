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

## Method 3: Mocking in Go

A fairly simple pattern for mocking an HTTP client in Go is to create a custom interface. Our interface will define the methods used in a function and pass different implementations depending on where the function is called from.

The custom interface for our HTTP client above should look like the following code block:
```go
type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}
```

Our signature for `getJoke()` will look like the code block below:
```go
func getJokeHandlerFunc(client HTTPClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
```

The original body of the `getJoke()` handler is moved inside of the return value. The `client` variable declaration is removed from the body in favor of the `HTTPClient` interface.

The `HTTPClient` interface wraps a `Do()` method, which accepts an HTTP request and returns an HTTP response and an error.

We need to provide a concrete implementation of `HTTPClient` when we call `getJokeHandlerFunc()` in the `main()` function:
```go
func main() {
    mux := http.NewServeMux()

    client := http.Client{
        Timeout: 10 * time.Second,
    }

    mux.HandleFunc("/joke", getJokeHandlerFunc(&client))

    http.ListenAndServe(":1212", mux)
}
```

The `http.Client` type implements the `HTTPClient` interface, so the program continues to call the Random dad joke API. We need to update the tests with a different `HTTPClient` implementation that does not make HTTP requests over the network.

First, we’ll create a mock implementation of the `HTTPClient` interface:
```go
type MockClient struct {
    DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
    return m.DoFunc(req)
}
```

In the code block above, the `MockClient` struct implements the `HTTPClient` interface through its provision of the `Do` method, which calls a `DoFunc` property. Now, we need to implement the `DoFunc` function when we create an instance of `MockClient` in the test:
```go
func TestWrapperGetJokeHandler(t *testing.T) {
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

            c := &MockClient{}

            c.DoFunc = func(req *http.Request) (*http.Response, error) {
                return &http.Response{
                    Body:       io.NopCloser(strings.NewReader(v.body)),
                    StatusCode: v.statusCode,
                }, nil
            }

            getJoke(c)(w, r)

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

In the code snippet above, `DoFunc` is adjusted for each test case, so it returns a custom response. Now, we’ve avoided all of the network calls, so the test will pass at a much faster rate:
```bash
$ go test -v -run=TestWrapperGetJokeHandler
=== RUN   TestWrapperGetJokeHandler
=== RUN   TestWrapperGetJokeHandler/R7UfaahVfFd
=== RUN   TestWrapperGetJokeHandler/173782
=== RUN   TestWrapperGetJokeHandler/#00
--- PASS: TestWrapperGetJokeHandler (0.00s)
    --- PASS: TestWrapperGetJokeHandler/R7UfaahVfFd (0.00s)
    --- PASS: TestWrapperGetJokeHandler/173782 (0.00s)
    --- PASS: TestWrapperGetJokeHandler/#00 (0.00s)
PASS
ok      github.com/favtuts/unit-testing 0.002s
```

You can use this same principle when your handler depends on another external system, like a database. Decoupling the handler from any specific implementation allows you to easily mock the dependency in the test while retaining the real implementation in your application’s code.