package main

import (
	"fmt"
	"time"
)

func display(ch chan int) {
	fmt.Println("Func display() is executing in 5 seconds...")
	time.Sleep(5 * time.Second)
	fmt.Println("Inside display()")
	ch <- 1234
	fmt.Println("Func display() is executed")
}

func main() {
	ch := make(chan int)
	go display(ch)
	//This statement will wait for data on channel ch
	x := <-ch

	//The main() on receiving the data from the channel gets unblocked and continues its execution.
	fmt.Println("Inside main()")
	fmt.Println("Printing x in main() after taking from channel:", x)
}
