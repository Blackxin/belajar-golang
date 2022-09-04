package main

import (
	"fmt"
	"runtime"
)

// func printMessage(what chan string) {
// 	fmt.Println(<-what)
// }

func main() {
	// penerapan channel
	// runtime.GOMAXPROCS(2)

	// var messages = make(chan string)

	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	messages <- data
	// }

	// go sayHelloTo("john wick")
	// go sayHelloTo("ethan hunt")
	// go sayHelloTo("jason bourne")

	// var message1 = <-messages
	// fmt.Println(message1)

	// var message2 = <-messages
	// fmt.Println(message2)

	// var message3 = <-messages
	// fmt.Println(message3)

	// channel sebagai tipe data parameter
	// for _, each := range []string{"wick", "hunt", "bourne"} {
	// 	go func(who string) {
	// 		data := fmt.Sprintf("hello ke %s", who)
	// 		messages <- data
	// 	}(each)
	// }

	// for i := 0; i < 3; i++ {
	// 	printMessage(messages)
	// }

	// penerapan buffered channel
	runtime.GOMAXPROCS(2)
	messages := make(chan int, 3)
	go func() {
		for {
			i := <-messages
			fmt.Println("receive data:", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data:", i)
		messages <- i
	}
}
