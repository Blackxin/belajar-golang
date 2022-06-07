package main

import "fmt"

func main() {
	// fruits := []string{"apple", "grape", "banana", "melon"}
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

	// fmt.Println(len(fruits))
	// fmt.Println(cap(fruits))

	// aFruits := fruits[0:3]
	// fmt.Println(len(aFruits))
	// fmt.Println(cap(aFruits))

	// bFruits := fruits[1:4]
	// fmt.Println(len(bFruits))
	// fmt.Println(cap(bFruits))

	// append
	// fruits := []string{"apple", "grape", "banana"}
	// cFruits := append(fruits, "papaya")

	// fmt.Println(fruits)
	// fmt.Println(cFruits)
	// bFruits := fruits[0:2]

	// fmt.Println(cap(bFruits))
	// fmt.Println(len(bFruits))

	// fmt.Println(fruits)
	// fmt.Println(bFruits)

	// cFruits := append(bFruits, "papaya")

	// fmt.Println(fruits)
	// fmt.Println(bFruits)
	// fmt.Println(cFruits)

	// copy
	// dst := make([]string, 3)
	// src := []string{"watermelon", "pineapple", "apple", "orange"}
	// n := copy(dst, src)

	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	// pengaksesan elemen slice dengan 3 indekx
	fruits := []string{"apple", "grape", "banana"}
	aFruits := fruits[0:2]
	bFruits := fruits[0:2:2]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
