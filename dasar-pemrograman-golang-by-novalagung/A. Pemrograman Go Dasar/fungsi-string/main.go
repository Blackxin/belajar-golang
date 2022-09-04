package main

import (
	"fmt"
	"strings"
)

func main() {
	isExists := strings.Contains("john wick", "wick")
	fmt.Println(isExists)

	isPrefix1 := strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1)
	isPrefix2 := strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2)

	isSuffix1 := strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1)
	isSuffix2 := strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2)

	howMany := strings.Count("ethan hunt", "t")
	fmt.Println(howMany)

	index1 := strings.Index("ethan hunt", "ha")
	fmt.Println(index1)
	index2 := strings.Index("ethan hunt", "n")
	fmt.Println(index2)

	text := "banana"
	find := "a"
	replaceWith := "o"
	newText1 := strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)
	newText2 := strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)
	newText3 := strings.Replace(text, find, replaceWith, 3)
	fmt.Println(newText3)

	str := strings.Repeat("na", 4)
	fmt.Println(str)

	string1 := strings.Split("the dark knight", " ")
	fmt.Println(string1)
	string2 := strings.Split("batman", "")
	fmt.Println(string2)

	data := []string{"banana", "papaya", "tomato"}
	str = strings.Join(data, "-")
	fmt.Println(str)

	str = strings.ToLower("aLaY")
	fmt.Println(str)

	str = strings.ToUpper("eat!")
	fmt.Println(str)
}
