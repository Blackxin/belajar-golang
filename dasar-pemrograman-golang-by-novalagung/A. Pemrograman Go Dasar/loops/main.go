package main

import "fmt"

func main() {
	// classic for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// while do style
	// i := 0
	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// break & continue
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// nested loop
	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

	// label
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Printf("Matriks [%v][%v]\n", i, j)
		}
	}
}
