package main

import (
	"fmt"
	"io"
)

// printer.go
//
//	func Print(text string) {
//		fmt.Println(text)
//	}

// func Print(text string) {
//     fmt.Fprintln(os.Stdout, text)
// }

func Print(text string, w io.Writer) {
	fmt.Fprintln(w, text)
}
