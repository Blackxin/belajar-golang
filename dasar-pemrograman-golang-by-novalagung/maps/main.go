package main

import "fmt"

func main() {
	// chicken := make(map[string]int)
	// chicken["januari"] = 50
	// chicken["februari"] = 40

	// chicken1 := map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// }

	// fmt.Println("januari", chicken["januari"])
	// fmt.Println("mei", chicken["mei"])

	// fmt.Println("Chicken 1:", chicken1)

	// keputusan := true
	// benarAtauSalah := map[bool]string{
	// 	true:  "benar",
	// 	false: "salah",
	// }

	// fmt.Printf("keputusannya adalah %v", benarAtauSalah[keputusan])

	// iterasi item map menggunakan for - range
	// chicken := map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }

	// for key, val := range chicken {
	// 	fmt.Printf("%v \t: %v\n", key, val)
	// 	// fmt.Println(key, " \t:", val)
	// }

	// menghapus item map
	// chicken := map[string]int{"januari": 50, "februari": 40}

	// fmt.Println(len(chicken))
	// fmt.Println(chicken)

	// delete(chicken, "januari")

	// fmt.Println(len(chicken))
	// fmt.Println(chicken)

	// deteksi keberadaan item dengan key tertentu
	// chicken := map[string]int{"januari": 50, "februari": 40}
	// value, isExist := chicken["mei"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("item is not exists")
	// }

	// kombinasi slice & map
	chickens := []map[string]string{
		{
			"name":   "chicken blue",
			"gender": "male",
		},
		{
			"name":   "chicken red",
			"gender": "male",
		},
		{
			"name":   "chicken yellow",
			"gender": "female",
		},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
