package main

import (
	"fmt"
	"log"
)

func TypeAssertionExample1() {
	var i interface{} = "a string"
	t := i.(string) // "a string"
	log.Println(t)
}

func TypeAssertionExample2() {
	var i interface{} = 1 //"a string"
	t, ok := i.(string)
	if ok {
		log.Println(t)
	} else {
		log.Fatal("Error on type assertion to string")
	}
}

func TypeAssertionExample3() {
	// var testVar interface{} = 1
	var testVar interface{} = 42.56

	switch t := testVar.(type) {
	case string:
		fmt.Println("Variable is of type string!")
		log.Println(t)
	case int:
		fmt.Println("Variable is of type int!")
		log.Println(t)
	case float32:
		fmt.Println("Variable is of type float32!")
		log.Println(t)
	default:
		fmt.Println("Variable type unknown")
		log.Println("Unknown type of testVar")
	}
}
