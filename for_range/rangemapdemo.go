package for_range

import "fmt"

func Demo() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}

	for k, v := range sozluk {
		fmt.Printf("%v %v \n", k, v)
	}
}
