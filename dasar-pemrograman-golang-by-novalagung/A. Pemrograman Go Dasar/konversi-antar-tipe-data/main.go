package main

import (
	"fmt"
)

func main() {
	// A.43.1. Konversi Menggunakan strconv
	// - Atoi
	// str := "124bruh"
	// num, err := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(num)

	// - Itoa
	// num := 124
	// str := strconv.Itoa(num)
	// fmt.Println(str)

	// - ParseInt
	// str := "124"
	// num, err := strconv.ParseInt(str, 10, 64)
	// if err == nil {
	// 	fmt.Println(num)
	// }

	// str := "1010"
	// num, err := strconv.ParseInt(str, 2, 8)
	// if err == nil {
	// 	fmt.Println(num)
	// }

	// - FormatInt
	// num := int64(24)
	// str := strconv.FormatInt(num, 8)
	// fmt.Println(str)

	// - ParseFloat
	// str := "24.12"
	// num, err := strconv.ParseFloat(str, 32)
	// if err == nil {
	// 	fmt.Println(num)
	// }

	// - ParseFloat
	// num := float64(24.12)
	// str := strconv.FormatFloat(num, 'f', 6, 64)
	// fmt.Println(str)

	// - ParseBool
	// str := "true"
	// bul, err := strconv.ParseBool(str)
	// if err == nil {
	// 	fmt.Println(bul)
	// }

	// - FormatBool
	// bul := true
	// str := strconv.FormatBool(bul)
	// fmt.Printf("%v: %T", str, str)

	// A.43.2. Konversi Data Menggunakan Teknik Casting
	// a := float64(24)
	// fmt.Println(a)

	// b := int32(24.00)
	// fmt.Println(b)

	// A.43.3. Casting string ↔ byte
	// text1 := "halo"
	// b := []byte(text1)
	// fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])

	// byte1 := []byte{104, 97, 108, 111}
	// s := string(byte1)
	// fmt.Printf("%s \n", s)

	// c := int64('h')
	// fmt.Println(c)
	// d := string(rune(104))
	// fmt.Println(d)

	// A.43.4. Type Assertions Pada Interface Kosong ( interface{} )
	data := map[string]any{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	// fmt.Println(data["nama"].(string))
	// fmt.Println(data["grade"].(int))
	// fmt.Println(data["height"].(float64))
	// fmt.Println(data["isMale"].(bool))
	// fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
