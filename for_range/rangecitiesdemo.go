package for_range

import "fmt"

func Cities() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for i, city := range cities {
		fmt.Printf("%v %v \n", i, city)
	}
}
