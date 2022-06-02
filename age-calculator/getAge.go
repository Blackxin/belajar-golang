package main

import (
	"time"
)

func GetAge(userBirthYear uint8) uint8 {
	var currentYear int = time.Now().Year()
	var userAge uint8 = uint8(currentYear) - userBirthYear
	return userAge
}
