package main

import (
	"fmt"
	"reflect"
)

// pengaksesan informasi property variabel objek
type student struct {
	Name  string
	Grade int
}

// func (s *student) getPropertyInfo() {
// 	reflectValue := reflect.ValueOf(s)
// 	if reflectValue.Kind() == reflect.Ptr {
// 		reflectValue = reflectValue.Elem()
// 	}

// 	reflectType := reflectValue.Type()

// 	for i := 0; i < reflectValue.NumField(); i++ {
// 		fmt.Println("name      :", reflectType.Field(i).Name)
// 		fmt.Println("tipe data :", reflectType.Field(i).Type)
// 		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
// 		fmt.Println("")
// 	}
// }

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	// number := 23
	// reflectValue := reflect.ValueOf(number)

	// fmt.Println("tipe variabel:", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variabel:", reflectValue.Int())
	// }

	// pengaksesan nilai dalmam bentuk interface{}
	// fmt.Println("tipe variabel  :", reflectValue.Type())
	// fmt.Println("nilai variabel :", reflectValue.Interface())

	// s1 := &student{Name: "wick", Grade: 2}
	// s1.getPropertyInfo()

	// pengaksesan informasi method variabel objek
	s1 := &student{Name: "john wick", Grade: 2}
	fmt.Println("name :", s1.Name)

	reflectValue := reflect.ValueOf(s1)
	method := reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("name :", s1.Name)
}
