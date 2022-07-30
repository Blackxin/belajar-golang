package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
}

func main() {
	// A.53.1. Decode JSON Ke Variabel Objek Struct
	// jsonString := `{ "fullName": "Taufik Hidayat", "age": 20 }`
	// jsonData := []byte(jsonString)

	// var data User

	// err := json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("user \t:", data.FullName)
	// fmt.Println("age \t:", data.Age)

	// A.53.2. Decode JSON Ke map[string]interface{} & interface{}
	// var data1 map[string]any
	// json.Unmarshal(jsonData, &data1)

	// fmt.Println("user \t:", data1["fullName"])
	// fmt.Println("age \t:", data1["age"])

	// var data2 any
	// json.Unmarshal(jsonData, &data2)

	// decodedData := data2.(map[string]any)
	// fmt.Println("user\t:", decodedData["fullName"])
	// fmt.Println("age\t:", decodedData["age"])

	// A.53.3. Decode Array JSON Ke Array Objek
	// jsonString := `[
	// 	{ "fullName": "john wick", "age": 27 },
	// 	{ "fullName": "ethan hunt", "age": 32 }
	// ]`

	// var data []User

	// err := json.Unmarshal([]byte(jsonString), &data)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println("user 1\t:", data[0].FullName)
	// fmt.Println("user 2\t:", data[1].FullName)

	// A.53.4. Encode Objek Ke JSON String
	object := []User{
		{
			FullName: "john wick",
			Age:      27,
		},
		{
			FullName: "ethan hunt",
			Age:      32,
		},
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}

	jsonString := string(jsonData)
	fmt.Println(jsonString)
}
