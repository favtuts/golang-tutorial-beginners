package main

import "fmt"

func main() {

	// Test custom error
	fmt.Printf("Error: %+v\n", CustomErrorInstance())
	fmt.Printf("Error: %+v\n", NewCustomError())
	fmt.Printf("Error: %+v\n", CustomErrorPkgErrors())

	// Test defer example
	//A1()

	// Test panic example
	//A2()

	// Test recover example
	A3()
}
