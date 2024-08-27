package main

import (
	"calcapp/formatters"
	"flag"
	"fmt"

	"favtuts.com/operations"
)

func main() {

	isSubtraction := flag.Bool("sub", false, "subtraction operation")
	aValue := flag.Int("a", 0, "a value")
	bValue := flag.Int("b", 0, "b value")

	flag.Parse()

	if *isSubtraction {
		formatters.Red(
			fmt.Sprintf(
				"Subtraction: %d",
				operations.Sub(*aValue, *bValue),
			),
		)
	} else {
		formatters.Green(
			fmt.Sprintf(
				"Addition: %d",
				operations.Add(*aValue, *bValue),
			),
		)
	}

}
