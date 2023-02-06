package main

import "fmt"

func main() {
	// declaring array
	a := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println("Array after creating:", a)

	var b []string = a[1:4] //created a slice named b
	fmt.Println("Slice after creation:", b)

	b[0] = "changed" // changed the slice data
	fmt.Println("Slice after modifying:", b)
	fmt.Println("Array after slice modification:", a)
}
