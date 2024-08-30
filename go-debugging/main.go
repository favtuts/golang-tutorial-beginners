package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Status   string `json:"status"`
	RegDate  string `json:"Reg_date"`
}

func main() {
	userinfos := []user{
		{
			FullName: "blessing james",
			Email:    "blessing@gmail.com",
			Gender:   "Male",
			Status:   "active",
			RegDate:  "20-01-2021",
		},
		{
			FullName: "matt john",
			Email:    "matt@gmail.com",
			Gender:   "Male",
			Status:   "active",
			RegDate:  "20-01-2021",
		},
		{
			FullName: "john peace",
			Email:    "peace@gmail.com",
			Gender:   "Midgard",
			Status:   "active",
			RegDate:  "20-01-2021",
		},
	}

	jsonBytes, err := json.Marshal(userinfos)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}
