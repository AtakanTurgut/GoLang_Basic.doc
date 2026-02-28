package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 75
	var y float64
	y = float64(x) // type(value)

	fmt.Println(y) // 75

	/* ------------------- */
	a := 5
	b := 10
	fmt.Println("A:", a, "B:", b) // A: 5 B: 10

	a, b = b, a                   // a, b = 10, 5 --> a = b, b = a
	fmt.Println("A:", a, "B:", b) // A: 10 B: 5

	/* ------------------- */
	// UTF-8
	yaş := 23
	fmt.Println(yaş) // 23

	სიმაღლე := 174
	წონა := 69.5
	fmt.Println("Height:", სიმაღლე, "Weight:", წონა) // Height: 174 Weight: 69.5

	/* ------------------- */
	// Shadowing
	q := 5
	k := 5

	if true {
		q := 11 // Block Scope  -  Shadowing
		q++
		fmt.Println(q) // 12

		k = 10
		k++
		fmt.Println(k) // 11
	}

	fmt.Println(q) // 5 - Func Scope
	fmt.Println(k) // 11

	/* ------------------- */
	num := 65          // int
	str := string(num) // ASCII

	fmt.Printf("%v %T \n", num, num) // 65 int
	fmt.Printf("%v %T \n", str, str) // A string

	str = strconv.Itoa(num)          // string - convert int to string
	fmt.Printf("%v %T \n", str, str) // 65 string --> "65"
}
