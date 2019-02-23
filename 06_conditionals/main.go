package main

import "fmt"

func main() {
	x := 10
	y := 10

	if x < y {
		fmt.Println(x, "is smaller than", y)
	} else if x == y {
		fmt.Println("Both are equal to", x)
	} else {
		fmt.Println(y, "is smaller than", x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	switch color {
	case "red":
		fmt.Println("Color is read")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}
}
