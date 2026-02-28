package for_range

import "fmt"

// 1-10 arasındaki sayılardan tek olanları toplayan fonksiyon
func Numbers() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := 0

	for _, number := range numbers {
		if number%2 != 0 {
			total = total + number
		}
	}

	fmt.Println("Toplam: ", total)
}
