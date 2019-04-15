package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	//p is similar to the "this" keyword in JS
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"John", "Doe", "New York", "m", 30}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person2.firstName)

	person1.age++
	fmt.Println(person1)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
