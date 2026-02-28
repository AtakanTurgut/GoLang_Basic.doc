package pointers

import "fmt"

func Numbers(numbers []int) {
	numbers[0] = 100

	fmt.Println("Demodaki sayı: ", numbers[0])
}
