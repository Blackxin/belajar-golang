package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

// func (s student) sayHello() {
// 	fmt.Println("halo", s.name)
// }

// func (s student) getNameAt(i int) string {
// 	return strings.Fields(s.name)[i-1]
// }

// method pointer
func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	// s1 := student{"john wick", 21}
	// s1.sayHello()

	// name := s1.getNameAt(2)
	// fmt.Println("nama panggilan :", name)

	// method pointer
	s1 := student{"john wick", 21}
	fmt.Println("s1 before:", s1.name)

	s1.changeName1("jason bourne")
	fmt.Println("s1 after changeName1:", s1.name)

	s1.changeName2("ethan hunt")
	fmt.Println("s1 after changeName2:", s1.name)
}
