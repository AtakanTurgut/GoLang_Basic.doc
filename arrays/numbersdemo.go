package arrays

import "fmt"

func Numbers() {
	numbers := [5]int{20, 30, 50, 10, 2}

	fmt.Println(numbers)

	bigger := numbers[0]

	// length - uzunluk
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > bigger {
			bigger = numbers[i]
		}

		fmt.Println(numbers[i])
	}

	fmt.Printf("En büyük sayı: %v \n", bigger)
}
