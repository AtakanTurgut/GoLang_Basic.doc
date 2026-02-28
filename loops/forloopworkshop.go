package loops

import "fmt"

func ForLoopWorkshop() {
	// sayı tahmini oyunu
	// kaç tahmin
	// 1-3: Süper 4-6: Başarılı <6: Fena Değil

	tutulanSayi := 80
	tahminEdilenSayi := 0

	fmt.Print("Bir sayı tahmin ediniz: ")
	fmt.Scanln(&tahminEdilenSayi)
	fmt.Println(tahminEdilenSayi)
	tahminSayisi := 1

	for tahminEdilenSayi != tutulanSayi {
		tahminSayisi++
		if tahminEdilenSayi < tutulanSayi {
			fmt.Println("Daha büyük bir sayı giriniz")

			fmt.Print("Bir sayı tahmin ediniz: ")
			fmt.Scanln(&tahminEdilenSayi)
		} else if tahminEdilenSayi > tutulanSayi {
			fmt.Println("Daha küçük bir sayı giriniz")

			fmt.Print("Bir sayı tahmin ediniz: ")
			fmt.Scanln(&tahminEdilenSayi)
		}
	}

	basariDurumu := ""

	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Başarılı"
	} else {
		basariDurumu = "Fena Değil"
	}

	fmt.Println("Tebrikler sayıyı buldunuz.")
	fmt.Printf("Tutulan sayı: %v Tahmin sayısı: %v %v \n", tutulanSayi, tahminSayisi, basariDurumu)
}
