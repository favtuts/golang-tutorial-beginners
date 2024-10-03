package main

import (
	"encoding/json"
	"fmt"
)

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species    string
	Dimensions Dimensions
}

// marshals a Dimensions struct into a JSON string
// with format "heightxwidth"
func (d Dimensions) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%dx%d"`, d.Height, d.Width)), nil
}

func main() {
	bird := Bird{
		Species: "pigeon",
		Dimensions: Dimensions{
			Height: 24,
			Width:  10,
		},
	}
	birdJson, _ := json.Marshal(bird)
	fmt.Println(string(birdJson))
	// {"Species":"pigeon","Dimensions":"24x10"}
}
