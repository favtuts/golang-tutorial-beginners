package main

import "fmt"

func sample() {
	fmt.Println("Inside the sample()")
}

func main() {
	//sample() will be invoked only after executing the statement of main()
	defer sample()
	fmt.Println("Inside the main()")
}
