package arrays

import "fmt"

func Demo() {
	// zero index
	var numbers [5]int
	numbers[2] = 50

	fmt.Println(numbers)
	fmt.Println(numbers[2])
}
