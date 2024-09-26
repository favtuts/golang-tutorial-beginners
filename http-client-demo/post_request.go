package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// we will run an HTTP server locally to test the POST request
	url := "http://localhost:3000/"

	// create post body
	body := strings.NewReader("This is the request body.")

	resp, err := http.Post(url, "text/plain", body)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)
}
