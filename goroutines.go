package main

import "fmt"

func display() {
	for i := 0; i < 5; i++ {
		fmt.Println("In display")
	}
}

func main() {
	//invoking the goroutine display()
	go display()

	//the main() continues without waiting for display()
	// the main() doesn’t wait for the display() to complete, and the main() completed its execution before the display() executed its code
	for i := 0; i < 5; i++ {
		fmt.Println("In main")
	}
}
