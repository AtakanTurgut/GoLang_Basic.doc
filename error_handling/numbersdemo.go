package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {

	aklimdakiSayi := 89

	if tahmin < 0 || tahmin > 100 {
		return "", errors.New("tahmin 0 ile 100 arasında olmalıdır")
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı girin", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı girin", nil
	}

	return "tahmin başarılı", nil
}

func Numbers() {
	mesaj, _ := TahminEt(17)
	fmt.Println(mesaj)

	fmt.Println(TahminEt(99))
	fmt.Println(TahminEt(89))

	fmt.Println(TahminEt(110))
	fmt.Println(TahminEt(-5))
}
