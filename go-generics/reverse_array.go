package main

import "fmt"

// T is a type parameter that is used like normal type inside the function
// any is a constraint on type e.e T has to implement "any" interface
func reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, ele := range s {
		r[l-i-1] = ele
	}

	return r
}

func demo_reverse_int_array() {
	fmt.Println(reverse([]int{1, 2, 3, 4, 5}))
}
