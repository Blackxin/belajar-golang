package main

import (
	"fmt"
	"time"
)

func getAge(userBirthYear int) int {
	var currentYear int = time.Now().Year()
	fmt.Printf("Current year: %v", currentYear)
	var userAge int = int(currentYear) - userBirthYear
	return userAge
}

func main() {
	var birthYear int

	fmt.Print("Enter your birth year: ")
	fmt.Scan(&birthYear)

	var userAge int = getAge(birthYear)

	fmt.Printf("Your age is: %v", userAge)
}
