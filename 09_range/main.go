package main

import (
	"fmt"
)

func main() {
	ids := []int{1, 34, 54, 1, 2, 5, 9, 7, 85}

	for _, id := range ids {
		// fmt.Printf("%d - ID: %d\n", i, id)
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println("The sum is", sum)

	// Range with map
	emails := map[string]string{"amit": "amit@gmail.com", "tanishka": "tanishka@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Just getting the key
	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}
