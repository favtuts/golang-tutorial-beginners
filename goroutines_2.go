package main

import (
	"fmt"
	"time"
)

//Now we modify the program to print the statements from display() as well.
//We add a time delay of 2 sec in the for loop of main() and a 1 sec delay in the for loop of the display().

func display() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("In display")
	}
}

func main() {
	//invoking the goroutine display()
	go display()

	//the main() continues without waiting for display()
	// the main() doesn’t wait for the display() to complete, and the main() completed its execution before the display() executed its code
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("In main")
	}
}
