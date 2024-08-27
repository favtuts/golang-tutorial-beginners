package main

import (
	"fmt"
)

func main() {
	//Just use := to declare and assign value, cannot use := to assign value to a variable which is already declared
	a := 20
	fmt.Println(a)

	//gives error since a is already declared
	// a := 30
	// fmt.Println(a)
}
