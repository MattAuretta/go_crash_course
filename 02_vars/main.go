package main

import "fmt"

func main() {
	//Using var

	var age int32 = 27
	const isCool = true

	//Shorthand var
	// name := "Matt"
	// email := "matt@gmail.com"

	name, email := "Matt", "matt@gmail.com"

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", email)
}
