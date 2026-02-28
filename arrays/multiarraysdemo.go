package arrays

import "fmt"

func MultiArrays() {
	var numbers [2][3]int

	numbers[0][0] = 1
	numbers[0][1] = 2
	numbers[0][2] = 3
	numbers[1][0] = 4
	numbers[1][1] = 5
	numbers[1][2] = 6

	fmt.Println(numbers)
	fmt.Println(numbers[1][1]) // 5

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(numbers[i][j])
		}
	}
}
