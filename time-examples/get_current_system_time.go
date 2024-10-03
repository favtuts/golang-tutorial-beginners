package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.String()) // Output: 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
}
