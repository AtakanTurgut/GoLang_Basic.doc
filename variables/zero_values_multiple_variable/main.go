package main

import "fmt"

func main() {
	/*
		var (
			name      string = "Atakan"
			age       int    = 23
			isMarried bool   = false

			weight float32 = 69.5
			height float32 = 172.5
		)
	*/

	// var name, age, isMarried, weight, height = "Atakan", 23, false, 69.5, 172.5

	name, age, isMarried, weight, height := "Atakan", 23, false, 69.5, 172.5

	var surname string
	var year int
	var money float32
	var isAdmin bool

	fmt.Println(name)
	fmt.Println(surname) // string --> ""  :  null  - Zero Values
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(year) // int --> 0  :  Zero Values
	fmt.Println(weight)
	fmt.Println(height)
	fmt.Println(money)   // float32 --> 0  :  Zero Values
	fmt.Println(isAdmin) // bool --> false  :  Zero Values
}
