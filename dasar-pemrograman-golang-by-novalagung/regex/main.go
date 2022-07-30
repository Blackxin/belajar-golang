package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "banana burger soup"
	regex, err := regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// res1 := regex.FindAllString(text, 2)
	// fmt.Printf("%#v \n", res1)

	// res2 := regex.FindAllString(text, -1)
	// fmt.Printf("%#v \n", res2)

	// isMatch := regex.MatchString(text)
	// fmt.Println(isMatch)

	// str := regex.FindString(text)
	// fmt.Println(str)

	// idx := regex.FindStringIndex(text)
	// fmt.Println(idx)

	// str = text[0:6]
	// fmt.Println(str)

	// str1 := regex.FindAllString(text, -1)
	// fmt.Println(str1)

	// str2 := regex.FindAllString(text, 1)
	// fmt.Println(str2)

	// str = regex.ReplaceAllString(text, "potato")
	// fmt.Println(str)

	// str = regex.ReplaceAllStringFunc(text, func(each string) string {
	// 	if each == "burger" {
	// 		return "potato"
	// 	}
	// 	return each
	// })
	// fmt.Println(str)

	str := regex.Split(text, -1)
	fmt.Printf("%#v \n", str)
}
