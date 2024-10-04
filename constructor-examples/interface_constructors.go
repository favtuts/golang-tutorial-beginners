package main

import "fmt"

type Pizza struct {
	slices   int
	toppings []string
	isBaked  bool
}

func (p Pizza) Bake() {
	p.isBaked = true
}

func (p Pizza) Display() {
	fmt.Println(p)
}

// Pizza implements Bakeable
type Bakeable interface {
	Bake()
	Display()
}

// this constructor will return a `Bakeable`
// and not a `Pizza`
func NewUnbakedPizza(toppings []string) Bakeable {
	return Pizza{
		slices:   6,
		toppings: toppings,
	}
}

func main() {
	my_pizza := NewUnbakedPizza([]string{})
	fmt.Println(my_pizza)
	my_pizza.Bake()
	fmt.Println(my_pizza)
	my_pizza.Display()
}
