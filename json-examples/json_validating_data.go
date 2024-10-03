package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"`
	// birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	if !json.Valid([]byte(birdJson)) {
		// handle the error here
		fmt.Println("invalid JSON string:", birdJson)
		return
	}

	var result map[string]any
	json.Unmarshal([]byte(birdJson), &result)
	fmt.Println("animals", result["animals"].(string))
}
