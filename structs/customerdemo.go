package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() { // struct metodu
	fmt.Printf("Müşteri Eklendi: %v %v \n", c.firstName, c.lastName)
}

func (c customer) update() {
	fmt.Printf("Müşteri Güncellendi: %v %v \n", c.firstName, c.lastName)
}

func Customers() {
	c := customer{firstName: "Atakan", lastName: "Turgut", age: 24}
	c.save()
	c.update()
}
