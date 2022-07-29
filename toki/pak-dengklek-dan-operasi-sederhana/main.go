package main

import (
	"fmt"
	"log"
)

func main() {
	var angka1, angka2 int
	fmt.Scan(&angka1, &angka2)

	if angka1 >= 1 && angka1 <= 100 {
		if angka2 >= 1 && angka2 <= 100 {
			fmt.Println(angka1 + angka2)
			fmt.Println(angka1 - angka2)
			fmt.Println(angka1 * angka2)
			fmt.Println(angka1 / angka2)
			fmt.Println(angka1 % angka2)
		} else {
			log.Fatalln("| Angka 2 is invalid")
		}
	} else {
		log.Fatalln("| Angka 1 is invalid")
	}
}
