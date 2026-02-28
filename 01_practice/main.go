package main

import "fmt"

func main() {
	var studentName1 string = "Atakan Turgut"
	var grade1 int = 79
	var isPassed1 bool = true

	fmt.Println(studentName1, grade1, isPassed1) // Atakan Turgut 79 true

	var studentName2 = "Fatih Necali"
	var grade2 = 80
	var isPassed2 = true

	fmt.Println(studentName2, grade2, isPassed2) // Fatih Necali 80 true

	studentName3 := "Yağmur Soydan"
	grade3 := 81
	isPassed3 := true

	fmt.Println(studentName3, grade3, isPassed3) // Yağmur Soydan 81 true

	var studentName4, grade4, isPassed4 = "Hasan Demir", 82, true

	fmt.Println(studentName4, grade4, isPassed4) // Hasan Demir 82 true

	studentName5, grade5, isPassed5 := "Ali Sabit", 83, true

	fmt.Println(studentName5, grade5, isPassed5) // Ali Sabit 83 true

	/* 	Declaration - Assign - Initialization - Initial Value 	*/
	/* 	Beyan - Atama - Başlatma - Başlangıç ​​Değeri	 */

	var studentName6 string = "Atilla Deniz" // Initialization
	// Declaration --> ""   = 	Assign
	fmt.Println(studentName6) // Initial Value
}
