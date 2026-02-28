package slices

import "fmt"

func Demo() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}

	fmt.Println(cities)

	citiesCopy := make([]string, len(cities))
	fmt.Println(citiesCopy)

	copy(citiesCopy, cities) // hedef, kaynak
	fmt.Println(citiesCopy)

	cities = append(cities, "Antalya")
	fmt.Println(cities)     // [Ankara İstanbul İzmir Antalya]
	fmt.Println(citiesCopy) // [Ankara İstanbul İzmir]

	// filtreleme
	fmt.Println(cities[1:3]) // [İstanbul İzmir]
	// 0 indeksten 1. indeksten 3. indekse (dahil değil)
	fmt.Println(cities[1:]) // [İstanbul İzmir Antalya]
	// 0 indeksten 1. indeksten sonuna kadar
	fmt.Println(cities[:2]) // [Ankara İstanbul]
	// 0 indeksten ilkten 2. indekse (dahil değil)
}
