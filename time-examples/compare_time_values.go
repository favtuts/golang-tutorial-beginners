package main

import (
	"fmt"
	"time"
)

func main() {
	myTime, err := time.Parse("2 Jan 06 03:04PM", "11 Nov 09 11:00PM")
	if err != nil {
		panic(err)
	}
	now := time.Now()
	fmt.Println(now.Before(myTime))
	fmt.Println(now.After(myTime))
	fmt.Println(now.Equal(myTime))
}
