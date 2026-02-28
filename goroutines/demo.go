// asenkron
package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Printf("Çift Sayı: %v \n", i)
		time.Sleep(1 * time.Second)
	}
}

func TekSayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("Tek Sayı: %v \n", i)
		time.Sleep(1 * time.Second)
	}
}
