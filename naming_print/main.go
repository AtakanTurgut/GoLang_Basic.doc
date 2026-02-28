package main

import "fmt"

// Package Scope Global --> Upper Case - variables
var E = 25

func main() {
	/* Print - Println - Printf */

	fmt.Print("Hello")   // crude
	fmt.Println("Hello") // crude
	fmt.Printf("Hello")  // formatted

	fmt.Println("")

	name := "Atakan"
	fmt.Print("Benim Adım", name) // Benim AdımAtakan
	fmt.Println("")
	fmt.Println("Benim Adım", name) // Benim Adım Atakan
	fmt.Printf("Benim Adım", name)  // Benim Adım%!(EXTRA string=Atakan)
	fmt.Println("")
	fmt.Printf("Benim Adım %v \n", name)          // Benim Adım Atakan
	fmt.Printf("Benim Adım %T \n", name)          // Benim Adım string
	fmt.Printf("Benim Adım %v %T \n", name, name) // Benim Adım Atakan string
	fmt.Printf("Benim Adım %x %X \n", name, name) // Benim Adım 4174616b616e 4174616B616E

	x := 100
	y := 20
	z := 30
	fmt.Printf("%b \n", x)             // binary 100 => 1100100
	fmt.Printf("%b %d %o \n", x, y, z) // 1100100 20 36
	// binary, decimal, octal - ikilik, onluk, sekizlik tabanda

	surname, age := "Turgut", 23
	fmt.Print("Benim adım ", name, " ", surname, ". Yaşım ", age, ". \n")  // Benim adım Atakan Turgut. Yaşım 23.
	fmt.Println("Benim adım", name, surname, "ve ben", age, "yaşındayım.") // Benim adım Atakan Turgut ve ben 23 yaşındayım.
	fmt.Printf("Benim adım %v %v. %v yaşındayım. \n", name, surname, age)  // Benim adım Atakan Turgut. 23 yaşındayım.

	/* Visibility */
	q := 10

	fmt.Println(q) // 10
	fmt.Println(E) // 25

	// Package Scope Global --> Upper Case - variables
	myFunc()

	var c string        // count - customer - coin ???
	var customer string // customer

	// Block Scope Local --> camel Case - variables
	var customerName string
	var customerSurname string

	// Abbreviations Upper Case
	var URL string  // Url xx
	var HTTP string // Http || http xx --> xyzHTTP

	fmt.Println(c, customer, customerName, customerSurname, URL, HTTP)
}

//fmt.Println(q)	--> non-declaration statement outside function body

func myFunc() {
	fmt.Println(E) // 25
}
