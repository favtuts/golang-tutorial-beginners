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
	birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Birds : %+v", birds)
	//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
}
