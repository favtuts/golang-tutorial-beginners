package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://www.example.com"

	// set a timeout of 50 milliseconds. This means that if the server does not
	// respond within 50 milliseconds, the request will fail with a timeout error
	http.DefaultClient.Timeout = 50 * time.Millisecond
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)
}
