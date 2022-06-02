package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	// ===== maps =====
	// statePopulations := make(map[string]int)
	// statePopulations = map[string]int{
	// 	"California":   395,
	// 	"Texas":        278,
	// 	"Florida":      286,
	// 	"New York":     197,
	// 	"Pennsylvania": 128,
	// 	"Illinois":     128,
	// 	"Ohio":         116,
	// }
	// sp := statePopulations
	// delete(sp, "Ohio")
	// fmt.Println(sp)
	// // pop, ok := statePopulations["Oho"]
	// fmt.Println(statePopulations)

	// ===== structs =====
	// aDoctor := Doctor{
	// 	number:    3,
	// 	actorName: "Taufik",
	// 	companions: []string{
	// 		"Fauzi",
	// 		"Doni",
	// 		"Fahrul",
	// 	},
	// }
	// fmt.Println(aDoctor)

	aDoctor := struct {
		name string
	}{
		name: "Taufik Hidayat",
	}
	anotherDoctor := &aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
