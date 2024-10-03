package main

import (
	"fmt"
	"time"
)

func main() {
	myTime, err := time.Parse("2 Jan 06 03:04PM", "10 Nov 09 11:00PM")
	if err != nil {
		panic(err)
	}
	now := time.Now()
	difference := now.Sub(myTime)
	fmt.Println(difference)

	fmt.Println(difference.Hours())
	fmt.Println(difference.Minutes())
	fmt.Println(difference.Seconds())
	fmt.Println(difference.Milliseconds())

	later := now.Add(3 * time.Hour)
	fmt.Println("now: ", now, "\nlater: ", later)
}
