package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	a := 5
	b := &a

	fmt.Println(a, b)

	fmt.Printf("%T\n", b)

	fmt.Println(a, *b)

	// Updating a
	a = 7
	fmt.Println(a, *b)
}
