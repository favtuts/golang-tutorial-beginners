package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	<-timer.C
	// Code to execute after 10 seconds
	now := time.Now()
	fmt.Println(now.String())
}
