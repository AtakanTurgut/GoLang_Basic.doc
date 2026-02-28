package main

import "fmt"

func main() {
	x := 10
	y := 10.5

	fmt.Printf("%v %T\n", x, x) // 10 int
	fmt.Printf("%v %T\n", y, y) // 10.5 float64

	// Type Conversion: type(value) => int(y) == 10.0 = 10

	fmt.Println(float64(x) + y) // 20.5
	fmt.Printf("%v %T\n", x, x) // 10 int

	/* ------------------- */
	var w int8 = 10
	var a int16 = 15

	fmt.Println(int16(w) + a) // 25

	/* ------------------- */
	var s int8 = 127
	var d int16

	d = int16(s) // type(value)

	fmt.Println(d) // 127

	/* ------------------- */
	var q int16 = 128 // 127 => 127
	var z int8        // 128 => -128

	z = int8(q) // type(value)

	fmt.Println(z)

	/* ------------------- */
	c := 10
	v := "10"

	fmt.Printf("%v %T\n", c, c) // 10 int
	fmt.Printf("%v %T\n", v, v) // 10 string

	//fmt.Println(c + int(v)) xx

	/* ------------------- */
	num1 := 106
	str1 := string(num1) // ASCII

	fmt.Printf("%v %T\n", num1, num1) // 106 int
	fmt.Printf("%v %T\n", str1, str1) // j string
}
