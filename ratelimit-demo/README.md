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

Let’s update the `limit.go` file to contain a basic implementation

```go
package main

import (
    "log"
    "net"
    "net/http"
    "sync"

    "golang.org/x/time/rate"
)

// Create a map to hold the rate limiters for each visitor and a mutex.
var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

// Retrieve and return the rate limiter for the current visitor if it
// already exists. Otherwise create a new rate limiter and add it to
// the visitors map, using the IP address as the key.
func getVisitor(ip string) *rate.Limiter {
    mu.Lock()
    defer mu.Unlock()

    limiter, exists := visitors[ip]
    if !exists {
        limiter = rate.NewLimiter(1, 3)
        visitors[ip] = limiter
    }

    return limiter
}

func limit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the IP address for the current user.
        ip, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
            log.Print(err.Error())
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        // Call the getVisitor function to retreive the rate limiter for
        // the current user.
        limiter := getVisitor(ip)
        if limiter.Allow() == false {
            http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}    
```


# Removing old entries from the map

We are going to record the last seen time for each visitor and running a background goroutine to delete old entries from the map

**ratelimit-demo/limit.go**
```go
package main

import (
    "log"
    "net"
    "net/http"
    "sync"
    "time"

    "golang.org/x/time/rate"
)

// Create a custom visitor struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type visitor struct {
    limiter  *rate.Limiter
    lastSeen time.Time
}

// Change the map to hold values of the type visitor.
var visitors = make(map[string]*visitor)
var mu sync.Mutex

// Run a background goroutine to remove old entries from the visitors map.
func init() {
    go cleanupVisitors()
}

func getVisitor(ip string) *rate.Limiter {
    mu.Lock()
    defer mu.Unlock()

    v, exists := visitors[ip]
    if !exists {
        limiter := rate.NewLimiter(1, 3)
        // Include the current time when creating a new visitor.
        visitors[ip] = &visitor{limiter, time.Now()}
        return limiter
    }

    // Update the last seen time for the visitor.
    v.lastSeen = time.Now()
    return v.limiter
}

// Every minute check the map for visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func limit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
            log.Print(err.Error())
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        limiter := getVisitor(ip)
        if limiter.Allow() == false {
            http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}   
```