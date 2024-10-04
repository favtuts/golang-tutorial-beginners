package main

import "fmt"

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

// define a function type that accepts a pointer to a House:
type HouseOption func(*House)

// define functional option that modify the *House instance
func WithConcrete() HouseOption {
	return func(h *House) {
		h.Material = "concrete"
	}
}

// define functional option that modify the *House instance
func WithoutFireplace() HouseOption {
	return func(h *House) {
		h.HasFireplace = false
	}
}

// define functional option that modify the *House instance
func WithFloors(floors int) HouseOption {
	return func(h *House) {
		h.Floors = floors
	}
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

// NewHouse now takes a slice of option as the rest arguments
func NewHouseWithOptions(opts ...HouseOption) *House {
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

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *House as the argument
		opt(h)
	}

	// return the modified house instance
	return h
}

func main() {
	my_house := NewHouse()
	fmt.Println(my_house)

	my_opt_house := NewHouseWithOptions(
		WithConcrete(),
		WithoutFireplace(),
		WithFloors(3),
	)

	fmt.Printf("%+v", my_opt_house)
}
