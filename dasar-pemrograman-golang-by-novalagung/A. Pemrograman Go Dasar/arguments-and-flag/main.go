package main

import (
	"flag"
	"fmt"
)

func main() {
	// A.48.1. Penggunaan Arguments
	// argsRaw := os.Args
	// fmt.Printf("-> %#v\n", argsRaw)

	// args := argsRaw[1:]
	// fmt.Printf("-> %#v\n", args)

	// A.48.2. Penggunaan Flag
	name := flag.String("name", "anonymous", "type your name")
	age := flag.Int64("age", 20, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)

	// A.48.3. Deklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	fmt.Printf("gender\t: %s\n", data2)
}
