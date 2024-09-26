#  Making REST API Requests in Golang using the HTTP Client 
* https://tuts.heomi.net/making-rest-api-requests-in-golang-using-the-http-client/

# Init a Go project

Locate the project directory
```bash
$ mkdir http-client-demo
$ cd http-client-demo
```

Then create the module
```bash
$ go mod init github.com/favtuts/http-client
```

# The HTTP Client Library

Most of the functionality for making HTTP requests is provided by the [net/http](https://pkg.go.dev/net/http) package. This package provides a [Client type](https://pkg.go.dev/net/http#Client) that we can use to make HTTP requests.

**Requests and Responses**

Whenever we make an HTTP request, we need to specify the HTTP method (GET, POST, etc), the URL, and the request body (if any).

In return, we get an HTTP response. The response contains the response body, the status code, and some other metadata. In Go, the response is represented by the [Response type](https://pkg.go.dev/net/http#Response).

![request-response](./images/request-response.png)

# Making a GET Request

For example, to make a GET request, we can use the [http.Get](https://pkg.go.dev/net/http#Get) method.

```go
func main() {
	url := "http://www.example.com"
	resp, err := http.Get(url)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// print the response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
```

Run the code:
```bash
$ go run get_request.go 
<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
        
    }

    ....
    </style>    
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission.</p>
    <p><a href="https://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>
```