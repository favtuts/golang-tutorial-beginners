#  Rate limiting your Go application 
* https://tuts.heomi.net/rate-limiting-your-go-application/

# Create web application

Create the `go.mod` file
```bash
$ go mod init ratelimiting
```

Start by creating a file called `limit.go` and add the following code to it:
```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Message struct {
    Status string `json:"status"`
    Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    message := Message{
        Status: "Successful",
        Body:   "Hi! You've reached the API. How may I help you?",
    }
    err := json.NewEncoder(writer).Encode(&message)
    if err != nil {
        return
    }
}

func main() {    
    http.HandleFunc("/ping", endpointHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Println("There was an error listening on port :8080", err)
    }
}
```

Run the web application
```bash
$ go run .

2024/08/28 15:43:03 Starting the web application...
```

Test the endpoint
```bash
curl --location 'http://localhost:8080/ping'
```


# Implement the token bucket algorithm

We’ll use Go’s low-level implementation with the [x/time/rate](https://godoc.org/golang.org/x/time/rate) package. Install it with the following terminal command:
```bash
$ go get golang.org/x/time/rate
```

Then, add the following path to your import statements:
```go
"golang.org/x/time/rate"
```

Add the following function to `limit.go`:
```go
func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
    limiter := rate.NewLimiter(2, 4)
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            message := Message{
                Status: "Request Failed",
                Body:   "The API is at capacity, try again later.",
            }

            w.WriteHeader(http.StatusTooManyRequests)
            json.NewEncoder(w).Encode(&message)
            return
        } else {
            next(w, r)
        }
    })
}
```

According to [the documentation](https://pkg.go.dev/golang.org/x/time/rate?utm_source=godoc#NewLimiter), `NewLimiter(r, b)` returns a new limiter that allows events up to rate `r` and permits bursts of at most `b` tokens.

Basically, `NewLimiter()` returns a struct that will allow an average of `r` requests per second and bursts, which are strings of consecutive requests, of up to `b`.

The middleware creates the `limiter` struct and then returns a handler created from an anonymous struct. 

The anonymous function uses the `limiter` to check if this request is within the rate limits with `limiter.Allow()`. If it is, the anonymous function calls the next function in the chain. If the request is out of limits, the anonymous function returns an error message to the client.

Next, use the middleware by changing the line that registers your handler with the default `servemux`:
```go
func main() {
    http.Handle("/ping", rateLimiter(endpointHandler))
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Println("There was an error listening on port :8080", err)
    }
} 
```

Now, run `limit.go` and fire off six requests to the `ping` endpoint in succession with cURL:
```bash
$ for i in {1..6}; do curl -i http://localhost:8080/ping; done
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 28 Aug 2024 09:26:39 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 28 Aug 2024 09:26:39 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 28 Aug 2024 09:26:39 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 28 Aug 2024 09:26:39 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 429 Too Many Requests
Date: Wed, 28 Aug 2024 09:26:39 GMT
Content-Length: 78
Content-Type: text/plain; charset=utf-8

{"status":"Request Failed","body":"The API is at capacity, try again later."}
HTTP/1.1 429 Too Many Requests
Date: Wed, 28 Aug 2024 09:26:39 GMT
Content-Length: 78
Content-Type: text/plain; charset=utf-8

{"status":"Request Failed","body":"The API is at capacity, try again later."}
```

You should see that the burst of four requests was successful, but the last two exceeded the rate limit and were rejected by the application.
