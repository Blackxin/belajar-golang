package main

import (
	"fmt"
	"math"
)

func main() {
	// names := []string{"John", "Wick"}
	// fmt.Println("halo", names)
	// rand.Seed(time.Now().Unix())
	// var randomValue int

	// randomValue = randomWithRange(2, 10)
	// fmt.Println("Random number:", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("Random number:", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("Random number:", randomValue)

	// divideNumber(10, 2)
	// divideNumber(4, 0)
	// divideNumber(8, -4)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f\n", area)
	fmt.Printf("keliling lingkaran\t: %.2f\n", circumference)
}

// func printMessage(message string, arr []string) {
// 	nameString := strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

// Fungsi Dengan Return Value / Nilai Balik
// func randomWithRange(min, max int) int {
// 	value := rand.Int()%(max-min+1) + min
// 	return value
// }

// Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi
// func divideNumber(m, n int) {
// 	if n == 0 {
// 		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
// 		return
// 	}

// 	var res = m / n
// 	fmt.Printf("%d / %d = %d\n", m, n, res)
// }

// Penerapan Fungsi Multiple Return
func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}
