package main

import (
	"calculation1"
	"fmt"
)

//the package to be created

func main() {
	x, y := 15, 10
	//the package will have function Do_add()
	sum := calculation1.Do_add(x, y)
	fmt.Println("Sum", sum)
}
