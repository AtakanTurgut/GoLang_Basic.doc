package slices

import "fmt"

func Names() {
	// slices : parçalar

	names := make([]string, 3) // veri türü, kaç eleman olacağı

	fmt.Println(names)

	names[0] = "atakan"
	names[1] = "yağmur"
	names[2] = "furkan"

	names = append(names, "büşra") // yeni slice oluştur - yeni eleman ekle

	fmt.Println(names)
	fmt.Println(len(names))
}
