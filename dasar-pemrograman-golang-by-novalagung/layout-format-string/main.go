package main

import "fmt"

// A.38.1. Persiapan
type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "taufik",
	height:      174.5,
	age:         20,
	isGraduated: false,
	hobbies:     []string{"Listening to music", "Gaming", "Coding"},
}

// ===================

func main() {
	// A.38.2. Layout Format %b
	fmt.Printf("%b\n", data.age)

	// A.38.3. Layout Format %c
	fmt.Printf("%c\n", 1400)

	// A.38.4. Layout Format %d
	fmt.Printf("%d\n", data.age)

	// A.38.5. Layout Format %e atau %E
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)

	// A.38.6. Layout Format %f atau %F
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.9f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)

	// A.38.7. Layout Format %g atau %G
	fmt.Printf("%e\n", 0.123123123123)
	fmt.Printf("%f\n", 0.123123123123)
	fmt.Printf("%g\n", 0.123123123123)

	// A.38.8. Layout Format %o
	fmt.Printf("%o\n", data.age)

	// A.38.9. Layout Format %p
	fmt.Printf("%p\n", &data.name)

	// A.38.10. Layout Format %q
	fmt.Printf("%q\n", `" name \ height "`)

	// A.38.11. Layout Format %s
	fmt.Printf("%s\n", data.name)

	// A.38.12. Layout Format %t
	fmt.Printf("%t\n", data.isGraduated)

	// A.38.13. Layout Format %T
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.height)
	fmt.Printf("%T\n", data.age)
	fmt.Printf("%T\n", data.isGraduated)
	fmt.Printf("%T\n", data.hobbies)

	// A.38.14. Layout Format %v
	fmt.Printf("%v\n", data)

	// A.38.15. Layout Format %+v
	fmt.Printf("%+v\n", data)

	// A.38.16. Layout Format %#v
	fmt.Printf("%#v\n", data)

	// A.38.17. Layout Format %x atau %X
	fmt.Printf("%x\n", data.age)

	// A.38.18. Layout Format %%
	fmt.Printf("%%\n")
}
