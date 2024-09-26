package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Person is a struct that represents the data we will send in the request body
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	url := "http://localhost:3000"

	// create post body using an instance of the Person struct
	p := Person{
		Name: "John Doe",
		Age:  25,
	}
	// convert p to JSON data
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	// We can set the content type here
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		log.Println("parsing received body to string...")
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)

		/*
			// parse the response
			log.Println("parsing received body to Person struct...")
			responsePerson := Person{}
			err = json.NewDecoder(resp.Body).Decode(&responsePerson)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("The person %v is %v ages", responsePerson.Name, responsePerson.Age)
		*/
	}

	fmt.Println("Status:", resp.Status)
}
