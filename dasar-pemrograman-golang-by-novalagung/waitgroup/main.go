package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Penerapan sync.WaitGroup
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		data := fmt.Sprintf("Data %d", i)

		wg.Add(1)
		go doPrint(&wg, data)
	}

	wg.Wait()
}

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}
