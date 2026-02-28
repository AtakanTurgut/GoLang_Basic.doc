package arrays

import "fmt"

func Cities() {
	var cities [5]string
	cities[0] = "Ankara"
	cities[1] = "Denizli"
	cities[2] = "Kayseri"
	cities[3] = "Isparta"
	cities[4] = "Antalya"

	fmt.Println(cities)

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i])
	}
}
