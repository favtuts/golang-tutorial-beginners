package main

import "fmt"

func main() {
	fmt.Printf("Error: %+v\n", CustomErrorInstance())
	fmt.Printf("Error: %+v\n", NewCustomError())
	fmt.Printf("Error: %+v\n", CustomErrorPkgErrors())
}
