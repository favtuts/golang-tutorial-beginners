package main

import "fmt"

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

// `NewHouse` is a constructor function for `*House`
func NewHouse() *House {
	const (
		defaultFloors       = 2
		defaultHasFireplace = true
		defaultMaterial     = "wood"
	)

	h := &House{
		Material:     defaultMaterial,
		HasFireplace: defaultHasFireplace,
		Floors:       defaultFloors,
	}

	return h
}

func main() {
	my_house := NewHouse()
	fmt.Println(my_house)
}
