package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var input string = "18"
	// fmt.Scan(&input)

	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalln("| Input should be number")
	}

	if num >= 0 && num <= 100 {
		fmt.Println("YA")
	} else {
		fmt.Println("TIDAK")
	}
}
