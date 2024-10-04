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

func getToppings(url string) ([]string, error) {
	// this should fetch some toppings from some remote URL
	// return []string{"peppers"}, nil
	return nil, fmt.Errorf("cannot fetch Piza info from url: %v", url)
}

func NewRemotePizza(url string) (Pizza, error) {
	// toppings are received from a remote URL, which may fail
	toppings, err := getToppings(url)
	if err != nil {
		// if an error occurs, return the wrapped error along with an empty
		// Pizza instance
		return Pizza{}, fmt.Errorf("could not construct new Pizza: %v", err)
	}
	return Pizza{
		slices:   6,
		toppings: toppings,
	}, nil
}

func main() {

	fmt.Println(NewPizza(nil)) // {6 []}

	somePizza := NewPizza([]string{"pepperoni"})
	fmt.Println(somePizza) // {6 [pepperoni]}

	otherPizza := NewPizza([]string{"onion", "pineapple"})
	fmt.Println(otherPizza) // {6 [onion pineapple]}

	remotePizza, err := NewRemotePizza("http://mypizza.org")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(remotePizza)
	}
}
