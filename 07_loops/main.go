package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// // SHort hand
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Number", i)
	// }

	// FizzBuzz
	for i := 1; i <= 25; i++ {
		if i%3 == 15 {
			fmt.Printf("FizzBuzz  ")
		} else if i%3 == 0 {
			fmt.Printf("Fizz  ")
		} else if i%5 == 0 {
			fmt.Printf("Buzz  ")
		} else {
			fmt.Printf("%d  ", i)
		}
	}
}
