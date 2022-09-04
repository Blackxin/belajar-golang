package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println("halo")
	// fmt.Println("selamat datang")

	// orderSomeFood("pizza")
	// orderSomeFood("burger")

	// kombinasi defer dan IIFE
	// number := 3

	// if number == 3 {
	// 	fmt.Println("halo 1")
	// 	func() {
	// 		defer fmt.Println("halo 3")
	// 	}()
	// }

	// fmt.Println("halo 2")

	// penerapan fungsi on.Exit()
	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")
}

// func orderSomeFood(menu string) {
// 	defer fmt.Println("terima kasih, silakan menunggu")
// 	if menu == "pizza" {
// 		fmt.Print("pilihan tepat!  ")
// 		fmt.Println("Pizza di tempat kami paling enak")
// 		return
// 	}

// 	fmt.Println("Pesanan anda:", menu)
// }
