package main

import "fmt"

func main() {
	fruits := []string{"apple", "grape", "banana", "melon"}
	// newFruits := fruits[0:2]

	// fmt.Println("fruits:", fruits)
	// fmt.Println("new fruits:", newFruits)
	// aFruits := fruits[0:3]
	// bFruits := fruits[1:4]

	// aaFruits := aFruits[1:2]
	// baFruits := bFruits[0:1]

	// fmt.Println(fruits)
	// fmt.Println(aFruits)
	// fmt.Println(bFruits)
	// fmt.Println(aaFruits)
	// fmt.Println(baFruits)

	// baFruits[0] = "pineapple"

	// fmt.Println(fruits)
	// fmt.Println(aFruits)
	// fmt.Println(bFruits)
	// fmt.Println(aaFruits)
	// fmt.Println(baFruits)

	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	aFruits := fruits[0:3]
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	bFruits := fruits[1:4]
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
