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