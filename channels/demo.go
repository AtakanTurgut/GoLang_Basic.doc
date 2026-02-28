// asenkron
package channels

import (
	"fmt"
	"time"
)

func CiftSayilar(ciftSayiToplamCn chan int) {
	toplam := 0

	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Çift Sayı Toplama Çalışıyor")
		time.Sleep(1 * time.Second)
	}

	ciftSayiToplamCn <- toplam
}

func TekSayilar(tekSayiToplamCn chan int) {
	toplam := 0

	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Tek Sayı Toplama Çalışıyor")
		time.Sleep(1 * time.Second)
	}

	tekSayiToplamCn <- toplam
}
