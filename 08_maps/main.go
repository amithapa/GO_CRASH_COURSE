package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	// // Defin map
	// emails := make(map[string]string)

	// //Assign key values
	// emails["amit"] = "amitt@gmail.com"
	// emails["tanishka"] = "tanishka@gmail.com"
	// emails["bunny"] = "bunny@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"amit": "amit@gmail.com"}
	fmt.Println(len(emails))
	fmt.Println(emails)

	//Delete from map
	delete(emails, "bunny")
	fmt.Println(emails)
}
