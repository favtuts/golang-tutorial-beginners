package main

import (
	"errors"
	"fmt"
)

func A3() {
	defer fmt.Println("Then we can't save the earth!")
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("Panic: %+v\n", x)
		}
	}()
	B3()
}
func B3() {
	defer fmt.Println("And if it keeps getting hotter...")
	C3()
}
func C3() {
	defer fmt.Println("Turn on the air conditioner...")
	Break3()
}
func Break3() {
	defer fmt.Println("If it's more than 30 degrees...")
	panic(errors.New("Global Warming!!!"))
}
