package main

import (
	"fmt"
	"log"
)

func test_empty_interface() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	age, ok := person["age"].(int)
	if !ok {
		log.Fatal("could not assert value to int")
		return
	}

	person["age"] = age + 1

	fmt.Printf("%+v", person)
}
