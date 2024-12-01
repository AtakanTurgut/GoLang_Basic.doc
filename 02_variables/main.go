// Package Clause
package main

// Import Statement
import "fmt"

// Code Field
func main() {
	var name string = "Atakan"
	fmt.Println("Hello", name)

	var age int
	age = 24
	fmt.Println(age)

	var isMarried bool
	isMarried = false
	fmt.Println(isMarried)

	var firstName, lastName string = "Atakan", "Turgut"
	fmt.Println(firstName, lastName)

	var year = "2001"
	fmt.Println(year)

	userName := "Yağmur" // Short declaration
	userAge := 21
	userAge = 22 // x"Soydan"
	fmt.Println(userName, userAge)
}
