package pointers

import "fmt"

func Demo(sayi *int) {
	*sayi = *sayi + 1

	fmt.Printf("Demodaki sayı: %v \n", *sayi)
}
