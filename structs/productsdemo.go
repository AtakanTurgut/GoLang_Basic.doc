package structs

import "fmt"

func Products() {
	x := product{"Klavye", 1000, "ABC"}
	fmt.Println(x)

	fmt.Println(product{"Bilgisayar", 50000, "XYZ"})
	fmt.Println(product{name: "Bilgisayar", unitPrice: 50000})

}

type product struct {
	name      string
	unitPrice float64
	brand     string
	//discountRate int
}
