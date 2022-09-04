package main

import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

func main() {
	// var secret interface{}

	// secret = "ethan hunt"
	// fmt.Println(secret)

	// secret = []string{"apple", "mango", "banana"}
	// fmt.Println(secret)

	// secret = 12.4
	// fmt.Println(secret)

	// data := map[string]any{
	// 	"name":      "Taufik Hidayat",
	// 	"grade":     2,
	// 	"breakfast": []string{"apple", "mango", "banana"},
	// }

	// secret = 2
	// number := secret.(int) * 10
	// fmt.Println(secret, "multiplied by 10 is :", number)

	// secret = []string{"apple", "mango", "banana"}
	// fruits := strings.Join(secret.([]string), ", ")
	// fmt.Println(fruits, "is my favorite fruits")

	// casting variabel interface kosong ke object pointer
	// var secret any = &person{name: "Taufik", age: 27}
	// name := secret.(*person).name
	// fmt.Println(name)

	// kombinasi slice, map, dan interface{}
	person := []map[string]any{
		{"name": "taufik", "age": 20},
		{"name": "hidayat", "age": 20},
		{"name": "fauzi", "age": 9},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	fruits := []any{"bruh", true, 69}
	fmt.Println("Fruits:", fruits)
}
