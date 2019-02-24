package main

import (
	"fmt"
	"strconv"
)

// Person Define a person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (Value Receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " amd I am " + strconv.Itoa(p.age)
}

//hasBirthday method (Pointer Receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried method (Pointer Receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		p.lastName = spouseLastName
	}
}

func main() {
	fmt.Println("Hello world")
	// Init Person using struct
	person1 := Person{firstName: "Amit", lastName: "Thapa", city: "Margao", gender: "M", age: 24}

	// Alternative
	person2 := Person{"Tanishka", "Amonkar", "Margao", "F", 21}

	fmt.Println(person1)
	fmt.Println(person2)

	person2.age++
	fmt.Println(person2)

	fmt.Println(person2.greet())
	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.getMarried("Amonkar")
	fmt.Println(person1.greet())
	person2.getMarried("Thapa")
	fmt.Println(person2.greet())
}
