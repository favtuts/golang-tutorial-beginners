package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

// unmarshals a JSON string with format
// "heightxwidth" into a Dimensions struct
func (d *Dimensions) UnmarshalJSON(data []byte) error {
	// the "data" parameter is expected to be JSON string as a byte slice
	// for example, `"20x30"`

	if len(data) < 2 {
		return fmt.Errorf("dimensions string too short")
	}
	// remove the quotes
	s := string(data)[1 : len(data)-1]
	// split the string into its two parts
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return fmt.Errorf("dimensions string must contain two parts")
	}
	// convert the two parts into ints
	height, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("dimension height must be an int")
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("dimension width must be an int")
	}
	// assign the two ints to the Dimensions struct
	d.Height = height
	d.Width = width
	return nil
}

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "dimensions":"20x30"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("%+v", bird)
	// {Species:pigeon Description:likes to perch on rocks Dimensions:{Height:20 Width:30}}
}
