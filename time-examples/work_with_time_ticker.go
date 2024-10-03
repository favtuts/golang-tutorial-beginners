package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop() // Ensure clean-up

	for {
		<-ticker.C
		// Code to execute every 5 seconds
		now := time.Now()
		fmt.Println(now.String())
	}
}
