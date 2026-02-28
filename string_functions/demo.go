package string_functions

// alias - takma ad
import (
	"fmt"
	s "strings"
)

// ascii - American Standard Code for Information Interchange
func Demo() {
	name := "Atakan"

	fmt.Println(name)
	fmt.Println(s.Count(name, "a"))      // 2 a - case sensitive
	fmt.Println(s.Contains(name, "kan")) // true
	fmt.Println(s.Index(name, "k"))      // 3 - index 3'te k var - zero based index
	status := s.Index(name, "b")         // -1 - b yok

	if status != -1 {
		fmt.Println("b harfi var")
	} else {
		fmt.Println("b harfi yok")
	}

	fmt.Println(s.ToLower(name)) // atakan
	fmt.Println(s.ToUpper(name)) // ATAKAN

	fmt.Println(s.HasPrefix(name, "Ata")) // true - At ile başlıyor mu
	fmt.Println(s.HasSuffix(name, "in"))  // false - in ile bitiyor mu
	fmt.Println(s.Index(name, "aka"))     // 2 - index 2'de aka var

	letters := []string{"a", "t", "a", "k", "a", "n"}
	atakan := s.Join(letters, "-")

	fmt.Println(atakan) // a-t-a-k-a-n - letters dizisindeki elemanları - ile birleştir

	fmt.Println(s.Replace(atakan, "-", "*", -1)) // a*t*a*k*a*n - atakan stringindeki tüm - karakterlerini * ile değiştir
	fmt.Println(s.Replace(atakan, "-", "*", 2))  // a*t*a-k-a-n - atakan stringindeki ilk 2 - karakterini * ile değiştir

	// 12398765411,28022026...
	fmt.Println(s.Split(atakan, "-"))      // [a t a k a n] - atakan stringini - karakterinden böl ve slice olarak döndür
	fmt.Println(s.Split(atakan, "*"))      // [a-t-a-k-a-n] - atakan stringini * karakterinden böl ve slice olarak döndür
	fmt.Println(len(s.Split(atakan, "-"))) // 6 - atakan stringini - karakterinden böl ve slice olarak döndür, sonra slice'ın uzunluğunu yazdır

	fmt.Println(s.Repeat("Go ", 3)) // Go Go Go - Go stringini 3 kere tekrar et
}
