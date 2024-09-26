package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://www.example.com"

	// create a new context instance with a timeout of 50 milliseconds
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)

	// create a new request using the context instance
	// now the context timeout will be applied to this request as well
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resp.Status)
}
