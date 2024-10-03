package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2 Jan 06 03:04PM")) // Output: 10 Nov 09 11:00PM
}
