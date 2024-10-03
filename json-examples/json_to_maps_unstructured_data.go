package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]any
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]any type, and its type is asserted from
	// the `any` type
	birds := result["birds"].(map[string]any)

	for key, value := range birds {
		// Each value is an `any` type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
}
