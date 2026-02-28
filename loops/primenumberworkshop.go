package loops

import "fmt"

func PrimeNumber() {
	// Kullanıcıdan bir  sayı girilecek
	// Alas sayı mı

	sayi := 0

	fmt.Print("Bir sayı giriniz: ")
	fmt.Scanln(&sayi)

	asalMi := true

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi {
		fmt.Printf("%v Asaldır \n", sayi)
	} else {
		fmt.Printf("%v Asal Değildir \n", sayi)
	}

}
