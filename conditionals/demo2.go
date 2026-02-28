package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz.")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız hazırlanıyor.")
		hesap = hesap - cekilmekIstenen
	} else {
		fmt.Println("Paranız hazırlanıyor.")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Println("İşleminiz tamamlanı. Hesaptaki para: " + fmt.Sprintf("%v", hesap))
}
