package variables

import "fmt"

func Demo() {
	fmt.Print("Merhaba ")
	fmt.Println("Dünya!")
	fmt.Println("Nasılsın?")

	var commit string = "Merhaba Türkiye!"
	fmt.Println(commit)

	// integer
	var kdv int = 18
	fmt.Println(kdv)
	fmt.Println(100 + (100 * kdv / 100))

	var extra float32 = 12.5
	fmt.Println(extra)
	fmt.Println(100 + (100 * extra / 100))

	number := 25.2
	//number = "Atakan"
	fmt.Println(number)
	fmt.Printf("Veri Tipi: %T \n", number) // f - format

	var status bool //= false)

	var name1 string = "Atakan"
	var name2 string = "Yağmur"

	status = name1 == name2 // !=
	fmt.Println(status)
	fmt.Println(!status)
}
