package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	// var s1 student
	// s1.name = "taufik hidayat"
	// s1.grade = 100

	// fmt.Println("name  :", s1.name)
	// fmt.Println("grade :", s1.grade)

	// inisialisasi object struct
	// var s1 = student{}
	// s1.name = "wick"
	// s1.grade = 2

	// var s2 = student{"ethan", 2}

	// s3 := student{name: "jason"}

	// fmt.Println("student 1 :", s1.name)
	// fmt.Println("student 2 :", s2.name)
	// fmt.Println("student 3 :", s3.name)

	// variabel object pointer
	s1 := student{
		name:  "wick",
		grade: 2,
	}
	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 2, name :", s2.name)

	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 2, name :", s2.name)

	// embedded struct

}
