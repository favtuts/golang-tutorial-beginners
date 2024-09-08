package main

import "fmt"

func A1() {
	defer fmt.Println("Keep calm!")
	B1()
}
func B1() {
	defer fmt.Println("Else...")
	C1()
}
func C1() {
	defer fmt.Println("Turn on the air conditioner...")
	D1()
}
func D1() {
	defer fmt.Println("If it's more than 30 degrees...")
}
