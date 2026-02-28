package interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int)
	if !ok {
		fmt.Printf("Sayı değil: %v %v \n", i, ok)
	}
	fmt.Printf("Sayı: %v %v \n", sayi, ok)
}

func TypeAssertion() {
	var sayi1 interface{} = "Atakan"
	dogrula(sayi1)

	var sayi2 interface{} = 23
	dogrula(sayi2)
}
