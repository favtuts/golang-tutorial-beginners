package main

import (
	"fmt"
	"strings"
)

type myString string

func (m myString) capitalize() myString {
	capStr := strings.ToUpper(string(m))

	return myString(capStr)
}

func TypeConvertionExample1() {
	fmt.Println("Hello World!")

	var m myString = "test"

	fmt.Println(m.capitalize())
}

func TypeConvertionExample2() {
	var simpleInt int = 3

	var simpleFloat float64 = 4.5

	// fmt.Println("This will throw an error: ", simpleInt + simpleFloat)

	fmt.Println("This will work correctly: ", simpleInt+int(simpleFloat))

	fmt.Println("This will work correctly too: ", float64(simpleInt)+simpleFloat)
}
