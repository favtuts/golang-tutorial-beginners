package main

import "fmt"

type Pizza struct {
	slices   int
	toppings []string
}

func NewPizza(toppings []string) Pizza {
	if toppings == nil {
		toppings = []string{}
	}
	return Pizza{
		slices:   6,
		toppings: toppings,
	}
}

func main() {

	fmt.Println(NewPizza(nil)) // {6 []}

	somePizza := NewPizza([]string{"pepperoni"})
	fmt.Println(somePizza) // {6 [pepperoni]}

	otherPizza := NewPizza([]string{"onion", "pineapple"})
	fmt.Println(otherPizza) // {6 [onion pineapple]}
}
