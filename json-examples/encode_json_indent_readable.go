package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

func main() {

	bird := Bird{
		Species:     "pigeon",
		Description: "likes to eat seed",
	}

	// The second parameter is the prefix of each line, and the third parameter
	// is the indentation to use for each level
	data, _ := json.MarshalIndent(bird, "", "  ")
	fmt.Println(string(data))
}
