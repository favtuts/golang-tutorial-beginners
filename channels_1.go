package main

//The main() finished the execution and did exit before the goroutine executes.
//So the print inside the display() didn’t get executed.

import (
	"fmt"
	"time"
)

func display() {
	time.Sleep(5 * time.Second)
	fmt.Println("Inside display()")
}

func main() {
	go display()
	fmt.Println("Inside main()")
}
