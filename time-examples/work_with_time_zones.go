package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	now := time.Now().In(loc)
	fmt.Println("Current time in New York:", now)

	utcTime := now.UTC()
	fmt.Println(utcTime)
}
