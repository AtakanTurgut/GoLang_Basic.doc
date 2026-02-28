package conditionals

import "fmt"

func Demo() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	var durum bool
	durum = cekilmekIstenen > hesap

	if durum {
		fmt.Println("Hesaptaki para yetersiz.")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız hazırlanıyor.")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Println("İşleminiz tamamlanı. Hesaptaki para: " + fmt.Sprintf("%v", hesap))

}
