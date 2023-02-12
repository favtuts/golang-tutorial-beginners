package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//reads the data and returns a byte sequence
	data, err := ioutil.ReadFile("filedata.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file: ", string(data))
}
