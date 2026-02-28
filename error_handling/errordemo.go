package error_handling

import (
	"fmt"
	"os"
)

func Demo() {
	f, err := os.Open("error demo.txt")
	// nil -> null, boş, sıfır
	if err != nil {
		// Type Assertion
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı: ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
	defer f.Close()
}
