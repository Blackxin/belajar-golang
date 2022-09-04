package main

import (
	"fmt"
	"strings"
)

func main() {
	// Closure Disimpan Sebagai Variabel
	// getMinMax := func(n []int) (int, int) {
	// 	var min, max int
	// 	for i, e := range n {
	// 		switch {
	// 		case i == 0:
	// 			max, min = e, e
	// 		case e > max:
	// 			max = e
	// 		case e < min:
	// 			min = e
	// 		}
	// 	}
	// 	return min, max
	// }

	// numbers := []int{2, 3, 4, 3, 4, 2, 3}
	// min, max := getMinMax(numbers)
	// fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max)

	// Immediately-Invoked Function Expression (IIFE)
	// numbers := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	// newNumbers := func(min int) []int {
	// 	var r []int
	// 	for _, e := range numbers {
	// 		if e < min {
	// 			continue
	// 		}
	// 		r = append(r, e)
	// 	}
	// 	return r
	// }(3)

	// fmt.Println("original number:", numbers)
	// fmt.Println("filtered number:", newNumbers)

	// max := 3
	// numbers := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	// howMany, getNumbers := findMax(numbers, max)
	// theNumber := getNumbers()

	// fmt.Println("number\t:", numbers)
	// fmt.Printf("find\t: %d\n\n", max)

	// fmt.Println("found\t:", howMany)
	// fmt.Println("value\t:", theNumber)

	data := []string{"wick", "jason", "ethan"}
	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	dataLength5 := filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli\t\t:", data)
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
}

// Closure Sebagai Nilai Kembalian
// func findMax(numbers []int, max int) (int, func() []int) {
// 	var res []int
// 	for _, e := range numbers {
// 		if e <= max {
// 			res = append(res, e)
// 		}
// 	}
// 	return len(res), func() []int {
// 		return res
// 	}
// }

// penerapan fungsi sebagai parameter
type filteredCallback func(string) bool

func filter(data []string, callback filteredCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
