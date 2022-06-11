package main

import (
	"fmt"
	"public-and-private/library"
)

// . "public-and-private/library"

func main() {
	// library.SayHello()
	// library.Introduce("taufik")

	// s1 := Student{"ethan", 21}
	// fmt.Println("name  :", s1.Name)
	// fmt.Println("grade :", s1.Grade)
	// sayHello("Taufik")

	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}
