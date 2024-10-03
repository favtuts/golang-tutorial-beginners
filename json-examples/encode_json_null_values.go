package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description *string
}

func main() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: nil,
	}

	data, _ := json.Marshal(pigeon)
	fmt.Println(string(data))
	// {"Species":"Pigeon","Description":null}

	birdData := map[string]any{
		"total birds": 2,
		"nullValue":   nil,
	}

	data, _ = json.Marshal(birdData)
	fmt.Println(string(data))
	// {"nullValue":null,"total birds":2}
}
