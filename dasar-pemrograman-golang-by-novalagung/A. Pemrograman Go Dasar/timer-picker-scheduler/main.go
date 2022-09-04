package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// A.41.1. Fungsi time.Sleep()
	// fmt.Println("start")
	// time.Sleep(time.Second * 5)
	// fmt.Println("after 5 seconds")

	// A.41.2. Scheduler Menggunakan time.Sleep()
	// for {
	// 	fmt.Println("Hello!")
	// 	time.Sleep(1 * time.Second)
	// }

	// A.41.3. Fungsi time.NewTimer()
	// timer := time.NewTimer(4 * time.Second)
	// fmt.Println("start")
	// <-timer.C
	// fmt.Println("finish")

	// A.41.4. Fungsi time.AfterFunc()
	// ch := make(chan bool)

	// time.AfterFunc(4*time.Second, func() {
	// 	fmt.Println("expired")
	// 	ch <- true
	// })

	// fmt.Println("start")
	// <-ch
	// fmt.Println("finish")

	// A.41.5. Fungsi time.After()
	// <-time.After(4 * time.Second)
	// fmt.Println("expired")

	// A.41.6. Scheduler Menggunakan Ticker
	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second)

	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop()
	// 		return
	// 	case t := <-ticker.C:
	// 		fmt.Println("Hello!!", t)
	// 	}
	// }

	// A.41.7. Kombinasi Timer & Goroutine
	timeout := 5
	ch := make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("What is 725 / 25? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntimeout! no answer more than", timeout, "seconds")
	os.Exit(1)
}
