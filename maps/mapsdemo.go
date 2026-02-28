package maps

import "fmt"

func Demo() {
	// key - value
	sozluk := make(map[string]string)

	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	delete(sozluk, "book") // silme
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	deger, varMi := sozluk["pencil"] // listede mevcutluk durumu
	fmt.Println(deger)
	fmt.Println("Listede olma durumu: ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)
}
