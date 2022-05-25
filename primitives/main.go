package main

import "fmt"

func main() {
	// boolean
	// n := 1 == 1
	// m := 1 == 2
	// fmt.Printf("%v, %T\n", n, n)
	// fmt.Printf("%v, %T\n", m, m)

	// integer
	// n := 42
	// var n uint16 = 42
	// fmt.Printf("%v, %T\n", n, n)
	// a := 10 // 1010
	// b := 3  // 0011
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(float32(a) / float32(b))
	// fmt.Println(a % b)

	// bitwise
	// fmt.Println(a & b)  // 0010
	// fmt.Println(a | b)  // 1011
	// fmt.Println(a ^ b)  // 1001
	// fmt.Println(a &^ b) // 0100

	// bit shifting
	// a := 8              // 2^3
	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0

	// var n float64 = 3.14
	// n = 13.7e72
	// n = 2.1e14
	// fmt.Printf("%v, %T", n, n)

	// a := 10.2
	// b := 3.7
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)

	// complex number
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T", n, n)
}
