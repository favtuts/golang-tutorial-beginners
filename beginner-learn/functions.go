package main

import "fmt"

//calc is the function name which accepts two integers num1 and num2
//(int, int) says that the function returns two values, both of integer type.
func calc(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}

func main() {
	x, y := 15, 10

	//calls the function calc with x and y and gets sum, diff as output
	sum, diff := calc(x, y)
	fmt.Println("Sum:", sum)
	fmt.Println("Diff:", diff)
}
