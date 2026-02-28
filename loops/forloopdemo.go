package loops

import "fmt"

func ForLoop() {
	var metin string = "Merhaba Dünya!"

	/*
		i := 1
		for i <= 5 {
			fmt.Println(metin)

			//infinite loop - sonsuz döngü
			i = i + 1
		}
	*/

	for i := 1; i <= 5; i++ {
		fmt.Println(metin + fmt.Sprintf(" %v", i))
	}
}
