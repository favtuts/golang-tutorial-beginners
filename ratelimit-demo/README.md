# How to Rate Limit HTTP Requests
* https://tuts.heomi.net/how-to-rate-limit-http-requests-in-go-application/


# Global rate limiter

Create a demo directory containing two files, `limit.go` and `main.go`, and initialize a new Go module
```bash
$ mkdir ratelimit-demo
$ cd ratelimit-demo
$ touch limit.go main.go
$ go mod init example.com/ratelimit-demo
```

We’ll use Go’s low-level implementation with the [x/time/rate](https://godoc.org/golang.org/x/time/rate) package. Install it with the following terminal command:
```bash
$ go get golang.org/x/time/rate
```

Let's start by making a global rate limiter which acts on *all the requests* that a HTTP server receives.

Open up the `limit.go` file and add the following code:

**ratelimit-demo/limit.go**

```go
package main

import (
    "net/http"

    "golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func limit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if limiter.Allow() == false {
            http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

Go ahead and run the application…
```bash
$ go run .
```

Make enough requests in quick succession
```bash
$ for i in {1..6}; do curl -i http://localhost:4000/; done
HTTP/1.1 200 OK
Date: Thu, 29 Aug 2024 02:45:43 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OKHTTP/1.1 200 OK
Date: Thu, 29 Aug 2024 02:45:43 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OKHTTP/1.1 200 OK
Date: Thu, 29 Aug 2024 02:45:43 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OKHTTP/1.1 429 Too Many Requests
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 29 Aug 2024 02:45:43 GMT
Content-Length: 18

Too Many Requests
HTTP/1.1 429 Too Many Requests
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 29 Aug 2024 02:45:43 GMT
Content-Length: 18

Too Many Requests
HTTP/1.1 429 Too Many Requests
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 29 Aug 2024 02:45:43 GMT
Content-Length: 18

Too Many Requests
```

# Rate limiting per user