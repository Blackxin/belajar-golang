package main

import (
	"fmt"
	"strconv"
)

// var (
// 	actorName    string = "Elisabeth Sladen"
// 	companion    string = "Sarah Jane Smith"
// 	doctorNumber int    = 3
// 	season       int    = 11
// )

// var I int = 27

func main() {
	// var seasonNumber int = 11
	// fmt.Println(seasonNumber)

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T", j, j)
}
