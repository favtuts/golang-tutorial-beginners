package main

import "fmt"

func main() {
	// creates an empty map with string keys and int values.
	myMap := make(map[string]int)

	// Adding Key-Value Pairs to a Map
	myMap["the_answer"] = 42

	// Getting a Value from a Map
	value := myMap["the_answer"]
	fmt.Printf("The value of key `the_answer` is %v\n", value)

	// Checking if a Key Exists in a Map
	value, ok := myMap["the_new_answer"]
	if !ok {
		fmt.Println("Can not find the value of key `the_new_answer`")
	}

	// Iterating Over a Map
	for key, value := range myMap {
		// do something with key and value
		fmt.Printf("%v has value: %v\n", key, value)
	}

	// Deleting Key-Value Pairs
	delete(myMap, "the_answer")
}
