package main

import "fmt"

func main() {
	var name = "Atakan"
	fmt.Println(name)

	var age int16 = -256
	fmt.Println(age)

	var isMarried = false
	fmt.Println(isMarried)

	var weight float32 = 69.5
	fmt.Println(weight)

	fmt.Printf("%T\n", name)      // string
	fmt.Printf("%T\n", age)       // int16
	fmt.Printf("%T\n", isMarried) // bool
	fmt.Printf("%T\n", weight)    // float32
}
