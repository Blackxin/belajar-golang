package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// A.37.1. Pemanfaatan Error
	// var input string
	// fmt.Print("Type some number: ")
	// fmt.Scanln(&input)

	// number, err := strconv.Atoi(input)

	// if err != nil {
	// 	fmt.Println(input, "is not number")
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(number, "is number")
	// }

	// A.37.2. Membuat Custom Error
	// var name string
	// fmt.Print("Type your name: ")
	// fmt.Scanln(&name)

	// if valid, err := validate(name); valid {
	// 	fmt.Println("Halo", name)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// A.37.3. Penggunaan panic
	// var name string
	// fmt.Print("Type your name: ")
	// fmt.Scanln(&name)

	// if valid, err := validate(name); valid {
	// 	fmt.Println("Halo", name)
	// } else {
	// 	panic(err.Error())
	// 	fmt.Println("end")
	// }

	// A.37.4. Penggunaan recover
	// defer catch()

	// var name string
	// fmt.Print("Type your name: ")
	// fmt.Scanln(&name)

	// if valid, err := validate(name); valid {
	// 	fmt.Println("Halo", name)
	// } else {
	// 	panic(err.Error())
	// }

	// A.37.5. Pemanfaatan recover pada IIFE
	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping:", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()

			panic("some error happen")
		}()
	}

}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
