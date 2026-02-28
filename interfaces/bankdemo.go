package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	var paymentTotal float64

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}

	return paymentTotal
}

func Bank() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Isparta"}
	credit2 := Mortgage{rate: 6, creditPaymentTotal: 50000, address: "Antalya"}
	credit3 := Car{rate: 12, creditPaymentTotal: 60000, carInfo: "Opel Astra"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Printf("Toplam Ödeme: %v \n", total)
}
