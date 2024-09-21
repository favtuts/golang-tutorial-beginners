#  Implementing JWT based authentication in Golang 
* https://tuts.heomi.net/implementing-jwt-based-authentication-in-golang/

# Init a Go project

Locate the project directory
```bash
$ mkdir jwt-go-example
$ cd jwt-go-example
```

Then create the module
```bash
$ go mod init github.com/favtuts/jwt-go-example
```

# How the JWT Signature Works

Here is the flow of JWT generation:

![jwt_generation](./images/jwt-algo.png)

To verify a JWT, the server generates the signature once again using the header and payload from the incoming JWT, and its secret key. If the newly generated signature matches the one on the JWT, then the JWT is considered valid.

![jwt_verification](./images/jwt-verification.png)


# Creating the HTTP Server

Let’s start by initializing the HTTP server with the required routes:
```go
package main

import (
	"log"
	"net/http"
)

func main() {
	// we will implement these handlers
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)

	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```