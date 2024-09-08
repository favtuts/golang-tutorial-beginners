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

# Mocking in Go

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

# Using external data in tests

In Go, you should place external data for tests in a directory called `testdata`. When you build binaries for your programs, the `testdata` directory is ignored, so you can use this approach to store inputs that you want to test your program against.

For example, let’s write a function that generates the `base64` encoding from a binary file:
```go
func getBase64Encoding(b []byte) string {
    return base64.StdEncoding.EncodeToString(b)
}
```

To test that this function produces the correct output, let’s place some sample files and their corresponding `base64` encoding in a `testdata` directory at the root of our project:
```bash
$ ls testdata
img1.jpg img1_base64.txt img2.jpg img2_base64.txt img3.jpg img3_base64.txt
```

To test our `getBase64Encoding()` function, run the code below:
```go
func TestGetBase64Encoding(t *testing.T) {
    cases := []string{"img1", "img2", "img3"}

    for _, v := range cases {
        t.Run(v, func(t *testing.T) {
            b, err := os.ReadFile(filepath.Join("testdata", v+".jpg"))
            if err != nil {
                t.Fatal(err)
            }

            expected, err := os.ReadFile(filepath.Join("testdata", v+"_base64.txt"))
            if err != nil {
                t.Fatal(err)
            }

            got := getBase64Encoding(b)

            if string(expected) != got {
                t.Fatalf("Expected output to be: '%s', but got: '%s'", string(expected), got)
            }
        })
    }
}
```

The bytes for each sample file are read from the file system and then fed into the `getBase64Encoding()` function. The output is subsequently compared to the expected output, which is also retrieved from the `testdata` directory.

Let’s make the test easier to maintain by creating a subdirectory inside of `testdata`. Inside of our subdirectory, we’ll add all of the input files, allowing us to simply iterate over each binary file and compare the actual to the expected output.

Now, we can add more test cases without touching the source code:
```bash
$ go test -v -run=TestGetBase64Encoding
=== RUN   TestGetBase64Encoding
=== RUN   TestGetBase64Encoding/iphone
=== RUN   TestGetBase64Encoding/android
--- PASS: TestGetBase64Encoding (0.00s)
    --- PASS: TestGetBase64Encoding/iphone (0.00s)
    --- PASS: TestGetBase64Encoding/android (0.00s)
PASS
ok      github.com/favtuts/unit-testing 0.004s
```

# Using golden files

If you’re using a Go template, it’s a good idea to test the generated output against the expected output to confirm that the template is working as intended. Go templates are usually large, so it’s not recommended to hard code the expected output in the source code as we’ve done so far in this tutorial.

Let’s explore an alternative approach to Go templates that simplifies writing and maintaining a test throughout a project’s lifecycle.

A golden file is a special type of file that contains the expected output of a test. The test function reads from the golden file, comparing its contents against a test’s expected output.

In the following example, we’ll use an `html/template` to generate an HTML table that contains a row for each book in an inventory:
```go
type Book struct {
    Name          string
    Author        string
    Publisher     string
    Pages         int
    PublishedYear int
    Price         int
}

var tmpl = `<table class="table">
  <thead>
    <tr>
      <th>Name</th>
      <th>Author</th>
      <th>Publisher</th>
      <th>Pages</th>
      <th>Year</th>
      <th>Price</th>
    </tr>
  </thead>
  <tbody>
    {{ range . }}<tr>
      <td>{{ .Name }}</td>
      <td>{{ .Author }}</td>
      <td>{{ .Publisher }}</td>
      <td>{{ .Pages }}</td>
      <td>{{ .PublishedYear }}</td>
      <td>${{ .Price }}</td>
    </tr>{{ end }}
  </tbody>
</table>
`

var tpl = template.Must(template.New("table").Parse(tmpl))

func generateTable(books []Book, w io.Writer) error {
    return tpl.Execute(w, books)
}

func main() {
    books := []Book{
        {
            Name:          "The Odessa File",
            Author:        "Frederick Forsyth",
            Pages:         334,
            PublishedYear: 1979,
            Publisher:     "Bantam",
            Price:         15,
        },
    }

    err := generateTable(books, os.Stdout)
    if err != nil {
        log.Fatal(err)
    }
}
```

The `generateTable()` function above creates the HTML table from a slice of `Book` objects. The code above will produce the following output:

```bash
$ go run main.go
<table class="table">
  <thead>
    <tr>
      <th>Name</th>
      <th>Author</th>
      <th>Publisher</th>
      <th>Pages</th>
      <th>Year</th>
      <th>Price</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>The Odessa File</td>
      <td>Frederick Forsyth</td>
      <td>Bantam</td>
      <td>334</td>
      <td>1979</td>
      <td>$15</td>
    </tr>
  </tbody>
</table>
```
To test the function above, we’ll capture the actual outcome and compare it to the expected outcome. We’ll store the expected result in the `testdata` directory as we did in the previous section, however, we’ll have to make a few changes.

Assume we have the following list of books in an inventory:
```go
var inventory = []Book{
    {
        Name:          "The Solitare Mystery",
        Author:        "Jostein Gaarder",
        Publisher:     "Farrar Straus Giroux",
        Pages:         351,
        PublishedYear: 1990,
        Price:         12,
    },
    {
        Name:          "Also Known As",
        Author:        "Robin Benway",
        Publisher:     "Walker Books",
        Pages:         208,
        PublishedYear: 2013,
        Price:         10,
    },
    {
        Name:          "Ego Is the Enemy",
        Author:        "Ryan Holiday",
        Publisher:     "Portfolio",
        Pages:         226,
        PublishedYear: 2016,
        Price:         18,
    },
}
```

The expected output for this list of books will span across many lines, therefore, it is difficult to place it as a string literal inside of the source code:
```html
<table class="table">
  <thead>
    <tr>
      <th>Name</th>
      <th>Author</th>
      <th>Publisher</th>
      <th>Pages</th>
      <th>Year</th>
      <th>Price</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>The Solitaire Mystery</td>
      <td>Jostein Gaarder</td>
      <td>Farrar Straus Giroux</td>
      <td>351</td>
      <td>1990</td>
      <td>$12</td>
    </tr>
    <tr>
      <td>Also Known As</td>
      <td>Robin Benway</td>
      <td>Walker Books</td>
      <td>308</td>
      <td>2013</td>
      <td>$10</td>
    </tr>
    <tr>
      <td>Ego Is The Enemy</td>
      <td>Ryan Holiday</td>
      <td>Portfolio</td>
      <td>226</td>
      <td>2016</td>
      <td>$18</td>
    </tr>
  </tbody>
</table>
```

In addition to being practical for larger outputs, a golden file can be automatically updated and generated.

While it’s possible to write a helper function to create and update golden files, we can take advantage of [goldie](https://github.com/sebdah/goldie), a utility that was created specifically for golden files.

Install the latest version of goldie with the command below:
```bash
$ go get -u github.com/sebdah/goldie/v2
```

Let’s go ahead and use goldie in a test for the `generateTable()` function:
```go
func TestGenerateTable(t *testing.T) {
    var buf bytes.Buffer

    err := generateTable(inventory, &buf)
    if err != nil {
        t.Fatal(err)
    }

    actual := buf.Bytes()

    g := goldie.New(t)
    g.Assert(t, "books", actual)
}
```

The test above captures the output of the `generateTable()` function in a buffer of bytes. Then, it passes the contents of the buffer to the `Assert()` method on the `goldie` instance. The contents on the buffer will be compared to the contents of the `books.golden` file in the `testdata` directory.

Initially, running the test will fail because we have not created the `books.golden` file yet:
```bash
$ go test -v -run=TestGenerateTable
=== RUN   TestGenerateTable
    books_golden_test.go:21: Golden fixture not found. Try running with -update flag.
--- FAIL: TestGenerateTable (0.00s)
FAIL
exit status 1
FAIL    github.com/favtuts/unit-testing 0.003s
```

The error message suggests that we add the `-update `flag, which will create the `books.golden` file with the contents of the buffer:
```bash
$ go test -v -update -run=TestGenerateTable
=== RUN   TestGenerateTable
--- PASS: TestGenerateTable (0.00s)
PASS
ok      github.com/favtuts/unit-testing 0.003s
```

On subsequent runs, we should remove the `-update` flag so that our golden file is not continually updated.

Any changes to the template should cause the test to fail. For example, if you update the price field to Euros (`€`) instead of USD (`$`), you’ll immediately receive an error. These errors occur because the output of the `generateTable()` function no longer matches the contents of the golden file.

Goldie provides diffing capabilities to help you spot the change when these errors occur:
```bash
$ go test -v -run=TestGenerateTable
=== RUN   TestGenerateTable
    books_golden_test.go:21: Result did not match the golden fixture. Diff is below:
        
        --- Expected
        +++ Actual
        @@ -18,3 +18,3 @@
               <td>1990</td>
        -      <td>$12</td>
        +      <td>€12</td>
             </tr><tr>
        @@ -25,3 +25,3 @@
               <td>2013</td>
        -      <td>$10</td>
        +      <td>€10</td>
             </tr><tr>
        @@ -32,3 +32,3 @@
               <td>2016</td>
        -      <td>$18</td>
        +      <td>€18</td>
             </tr>
        
--- FAIL: TestGenerateTable (0.00s)
FAIL
exit status 1
FAIL    github.com/favtuts/unit-testing 0.003s
```

In the output above, the change is clearly highlighted. These changes are deliberate, so we can make our test pass again by updating the golden file using the `-update` flag:
```bash
$ go test -v -update -run=TestGenerateTable
```