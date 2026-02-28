package defer_statement

import "fmt"

func Numbers(sayi int) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		return "sayı çifttir"
	}

	if sayi%2 != 0 {
		return "sayı tektir"
	}

	return "Belli değil"
}

func Test() {
	sonuc := Numbers(90)

	fmt.Printf("Sonuç: %v \n", sonuc)
}

func DeferFunc() {
	fmt.Println("DeferFunc çalıştı")
}
