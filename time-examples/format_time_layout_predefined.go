package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.UnixDate))
	fmt.Println(now.Format(time.RubyDate))
}
