package functions

import "fmt"

func Topla(sayi1 int, sayi2 int) int { // int return
	var toplam = sayi1 + sayi2

	return toplam
}

func SelamVer(kullaniciAdi string) {
	fmt.Printf("Merhaba, %v \n", kullaniciAdi)
}
