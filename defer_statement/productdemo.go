package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   float64
}

func (p product) save() {
	fmt.Printf("Veritabanına kaydedildi: %v \n", p.productName)
	defer Log()
}

func Products() {
	p1 := product{productName: "Laptop", unitPrice: 5000}
	defer p1.save()
	p1 = product{productName: "Telefon", unitPrice: 3000}

	fmt.Println("İşlem Başarılı")
	fmt.Println(p1.productName)
}

func Log() {
	fmt.Println("Loglandı")
}
