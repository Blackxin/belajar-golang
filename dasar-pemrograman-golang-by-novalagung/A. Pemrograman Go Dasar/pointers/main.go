package main

import "fmt"

func main() {
	// var number *int
	// var name *string

	// penerapan pointer
	// numberA := 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value)\t\t:", numberA)
	// fmt.Println("numberA (address)\t:", &numberA)

	// fmt.Println("numberB (value)\t\t:", *numberB)
	// fmt.Println("numberB (address)\t:", numberB)

	// efek penerapan nilai pointer
	// numberA := 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value)    :", numberA)
	// fmt.Println("numberA (address)  :", &numberA)
	// fmt.Println("numberB (value)    :", *numberB)
	// fmt.Println("numberB (address)  :", numberB)

	// fmt.Println()

	// numberA = 5

	// fmt.Println("numberA (value)    :", numberA)
	// fmt.Println("numberA (address)  :", &numberA)
	// fmt.Println("numberB (value)    :", *numberB)
	// fmt.Println("numberB (address)  :", numberB)

	number := 4
	fmt.Println("before :", number)

	change(&number, 10)
	fmt.Println("after  :", number)
}

// parameter pointer
func change(original *int, value int) {
	*original = value
}
