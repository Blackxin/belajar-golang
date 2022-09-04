package main

// type person struct {
// 	name string `tag1`
// 	age  uint16 `tag2`
// }

// type people = person

// type student struct {
// 	person struct {
// 		name string
// 		age  uint16
// 	}
// 	grade   uint8
// 	hobbies []string
// }

// type student struct {
// 	person
// 	age   uint16
// 	grade uint8
// }

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
	// s1 := student{
	// 	name:  "wick",
	// 	grade: 2,
	// }
	// var s2 *student = &s1
	// fmt.Println("student 1, name :", s1.name)
	// fmt.Println("student 2, name :", s2.name)

	// s2.name = "ethan"
	// fmt.Println("student 1, name :", s1.name)
	// fmt.Println("student 2, name :", s2.name)

	// embedded struct
	// s1 := student{}
	// s1.name = "Taufik"
	// s1.age = 20
	// s1.grade = 2

	// fmt.Println("name  :", s1.name)
	// fmt.Println("age   :", s1.age)
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("grade :", s1.grade)

	// Embedded Struct Dengan Nama Property Yang Sama
	// s2 := student{}
	// s2.name = "Hidayat"
	// s2.age = 20
	// s2.person.age = 21
	// s2.grade = 3

	// fmt.Println("name  :", s2.name)
	// fmt.Println("age   :", s2.age)
	// fmt.Println("age   :", s2.person.age)
	// fmt.Println("grade :", s2.grade)

	// Anonymous Struct
	// s1 := struct {
	// 	person
	// 	grade int
	// }{}
	// s1.person = person{"wick", 21}
	// s1.grade = 2

	// fmt.Println("name  :", s1.person.name)
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("grade :", s1.grade)

	// kombinasi slice & struct
	// students := []person{
	// 	{
	// 		name: "Taufik",
	// 		age:  20,
	// 	},
	// 	{
	// 		name: "Doni",
	// 		age:  20,
	// 	},
	// 	{
	// 		name: "Aulia",
	// 		age:  19,
	// 	},
	// }

	// for _, student := range students {
	// 	fmt.Println(student.name, "age is", student.age)
	// }

	// Inisialisasi Slice Anonymous Struct
	// students := []struct {
	// 	person
	// 	grade int
	// }{
	// 	{person: person{"wick", 21}, grade: 2},
	// 	{person: person{"ethan", 22}, grade: 3},
	// 	{person: person{"bond", 21}, grade: 3},
	// }
	// for _, student := range students {
	// 	fmt.Println(student)
	// }

	// Deklarasi Anonymous Struct Menggunakan Keyword var
	// var student struct {
	// 	person
	// 	grade int
	// }

	// student.person = person{"wick", 22}
	// student.grade = 2

	// nested struct

}
