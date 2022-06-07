package main

import "fmt"

func main() {
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"

	// fmt.Println(names)

	// inisialisasi di awal
	// var fruits [4]string
	// horizontal
	// fruits = [4]string{"Apple", "Grape", "Banana", "Melon"}
	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)
	// vertical
	// fruits = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }

	// inisialisasi nilai awal array tanpa jumlah elemen
	// var numbers = [...]int{2, 3, 2, 4, 3}

	// fmt.Println("Data array \t:", numbers)
	// fmt.Println("Jumlah elemen \t:", len(numbers))

	// array multidimensi
	// var numbers1 = [2][3]int{
	// 	[3]int{3, 2, 3},
	// 	[3]int{3, 4, 5},
	// }
	// var numbers2 = [2][3]int{
	// 	{3, 2, 3},
	// 	{3, 4, 5},
	// }

	// fmt.Println("Numbers1", numbers1)
	// fmt.Println("Numbers2", numbers2)

	// perulangan elemen array menggunakan keyword for
	// fruits := [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d: %s\n", i, fruits[i])
	// }

	// perulangan elemen array menggunakan keyword for range
	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d: %s\n", i, fruit)
	// }

	// Penggunaan Variabel Underscore _ Dalam for - range
	// for _, fruit := range fruits {
	// 	fmt.Printf("nama buah: %v\n", fruit)
	// }

	// Alokasi Elemen Array Menggunakan Keyword make
	fruits := make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "mango"
	fmt.Println(fruits)
}
