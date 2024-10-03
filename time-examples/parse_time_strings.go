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
	fmt.Println(myTime)
	fmt.Println(myTime.String())
	fmt.Println(myTime.Local().String())
}
