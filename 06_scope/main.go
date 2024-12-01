package main

import "fmt"

var packVar = "Package Scope" // Package Global
//packVar := "Package Scope"
// Short declaration cannot be used in package variables

var funcVar = "Func(Package) Scope" // !!!

func main() {

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar) // Block Scope   => blockVarblockVar
	}

	//var funcVar = "Func Scope"
	funcVar := "Func Scope" // Short declaration - Block Local
	fmt.Println(funcVar)    // Func Scope  => funcVar

	fmt.Println(packVar) // Package Scope  => packVar

	myFunc()

	var name = "Atakan"
	//name := "Yağmur"  xx second declaration xx
	name, surname := "Yağmur", "Soydan" // Assign, Declaration
	fmt.Println(name, surname)          // Yağmur Soydan
}

func myFunc() {
	fmt.Println(packVar) // Package Scope  		 => packVar

	fmt.Println(funcVar) // Func(Package) Scope  => funcVar !!!
}
