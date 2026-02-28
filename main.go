// go run main.go
// go mod init golang

package main

import (
	"fmt"
	"golang/arrays"
	"golang/channels"
	"golang/conditionals"
	"golang/defer_statement"
	"golang/error_handling"
	"golang/for_range"
	"golang/functions"
	"golang/goroutines"
	"golang/interfaces"
	"golang/loops"
	"golang/maps"
	"golang/pointers"
	"golang/slices"
	"golang/string_functions"
	"golang/structs"
	"golang/variables"
	"time"
)

func main() {
	// module - modül

	variables.Demo()

	conditionals.Demo()
	conditionals.Demo2()
	conditionals.Workshop()

	loops.ForLoop()
	loops.ForLoopWorkshop()
	loops.PrimeNumber()
	loops.FriendNumbers()

	arrays.Demo()
	arrays.Cities()
	arrays.Numbers()
	arrays.MultiArrays()

	slices.Names()
	slices.Demo()

	functions.SelamVer("Atakan")
	var sonuc = functions.Topla(2, 6)
	fmt.Printf("Sonuç: %v \n", sonuc*10)

	toplam, cikarim, carpim, bolum := functions.DortIslem(10, 6)
	// toplam, cikarim, carpim, _ := functions.DortIslem(10, 6) -> _ tanımında bana o değeri verme demek oluyor
	fmt.Printf("Toplam: %v \n", toplam)
	fmt.Printf("Çıkarım: %v \n", cikarim)
	fmt.Printf("Çarpım: %v \n", carpim)
	fmt.Printf("Bölüm: %v \n", bolum)

	var sonucVari = functions.ToplaVariadic(1, 4, 6, 3, 10)
	fmt.Printf("Variadic Sonuc: %v \n", sonucVari)

	sayilarVari := []int{4, 6, 7, 11}
	fmt.Println(functions.ToplaVariadic(sayilarVari...)) // ... variadic

	maps.Demo()

	for_range.Cities()
	for_range.Numbers()
	for_range.Demo()

	sayi := 20
	pointers.Demo(&sayi)
	fmt.Printf("Maindeki sayı: %v \n", sayi)

	sayilar := []int{1, 2, 3}
	pointers.Numbers(sayilar)
	fmt.Printf("Maindeki sayı: %v \n", sayilar[0])

	structs.Products()
	structs.Customers()

	/* goroutines */
	go goroutines.CiftSayilar()
	go goroutines.TekSayilar()

	fmt.Println("Main Bitti")
	time.Sleep(5 * time.Second)

	/* channels */
	ciftSayiToplamCn := make(chan int)
	tekSayiToplamCn := make(chan int)
	go channels.CiftSayilar(ciftSayiToplamCn)
	go channels.TekSayilar(tekSayiToplamCn)

	ciftSayiToplamS, tekSayiToplamS := <-ciftSayiToplamCn, <-tekSayiToplamCn

	carpimCn := ciftSayiToplamS * tekSayiToplamS
	fmt.Printf("Çarpım: %v \n", carpimCn)

	interfaces.Demo()
	interfaces.Bank()

	defer_statement.B()
	defer_statement.Test()
	defer_statement.Products()

	error_handling.Demo()
	interfaces.TypeAssertion()
	error_handling.Numbers()

	fmt.Println(error_handling.TahminEtException(89))

	string_functions.Demo()

	fmt.Println()
}
