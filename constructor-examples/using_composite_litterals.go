package main

import "fmt"

type Book struct {
	title string
	pages int
}

type Pizza struct {
	slices   int
	toppings []string
}

func main() {
	// creating a new struct instance
	b := Book{}

	fmt.Println(b) // { 0}

	// creating a pointer to a struct instance
	bp := &Book{}

	fmt.Println(bp)  // &{ 0}
	fmt.Println(*bp) // { 0}

	// creating an empty value
	nothing := struct{}{}
	fmt.Println(nothing) // {}

	b1 := Book{
		title: "Julius Caesar",
		pages: 322,
	}
	fmt.Println(b1)

	somePizza := Pizza{
		slices:   6,
		toppings: []string{"pepperoni"},
	}
	fmt.Println(somePizza)

	otherPizza := Pizza{
		slices:   6,
		toppings: []string{"onion", "pineapple"},
	}
	fmt.Println(otherPizza)
}
