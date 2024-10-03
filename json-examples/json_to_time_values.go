package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bird struct {
	Species     string
	Description string
	CreatedAt   time.Time
}

func main() {
	dateJson := `"2021-10-18T11:08:47.577Z"`
	var date time.Time
	json.Unmarshal([]byte(dateJson), &date)

	fmt.Println(date)
	// 2021-10-18 11:08:47.577 +0000 UTC

	birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "createdAt": "2021-10-18T11:08:47.577Z"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
	// {pigeon likes to perch on rocks 2021-10-18 11:08:47.577 +0000 UTC}
}
