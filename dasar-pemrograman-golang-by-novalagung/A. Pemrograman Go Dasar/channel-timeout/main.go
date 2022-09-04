package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	messages := make(chan int)

	go sendData(messages)
	retrieveData(messages)
}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Printf("receive data \"%v\"\n", data)
		case <-time.After(time.Second * 5):
			fmt.Println("timeout, no activities under 5 seconds")
			break loop
		}
	}
}
