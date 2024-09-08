package main

import (
	"errors"
	"fmt"
)

func A2() {
	defer fmt.Println("Then we can't save the earth!")
	B2()
}
func B2() {
	defer fmt.Println("And if it keeps getting hotter...")
	C2()
}
func C2() {
	defer fmt.Println("Turn on the air conditioner...")
	Break()
}
func Break() {
	defer fmt.Println("If it's more than 30 degrees...")
	panic(errors.New("Global Warming!!!"))
}
