package main

import (
	"fmt"
)

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// const (
// 	a2 = iota
// 	b2 = iota
// 	c2 = iota
// )

// const (
// 	_  = iota
// 	KB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// const myConst int = 42
	// const myConst float64 = math.Sin(1.57)
	// fmt.Printf("%v, %T", myConst, myConst)

	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)

	// fmt.Printf("%v\n", a2)
	// fmt.Printf("%v\n", b2)
	// fmt.Printf("%v\n", c2)

	// fileSize := 4000000000.
	// fmt.Printf("%.2fGB", fileSize/GB)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)

}
