// defer ifadesi, bir fonksiyonun sonunda çalışacak olan bir işlemi belirtmek için kullanılır. Bu, genellikle kaynakların serbest bırakılması veya temizlenmesi gibi durumlarda kullanılır. Defer ifadesi, fonksiyonun tamamlanmasından sonra çalışır ve genellikle kaynakları serbest bırakmak veya temizlemek için kullanılır.
package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}

func B() {
	defer A()
	// defer ifadesi, fonksiyonun sonunda çalışacak olan bir işlemi belirtir. Bu durumda, B fonksiyonu tamamlandıktan sonra A fonksiyonu çalışacaktır.
	defer C()
	defer D()
	fmt.Println("B fonksiyonu çalıştı")
	// D()
	// C()
	// A()
}

func C() {
	fmt.Println("C fonksiyonu çalıştı")
}

func D() {
	fmt.Println("D fonksiyonu çalıştı")
}
