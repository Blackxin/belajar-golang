package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	sync.Mutex
	val int
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				meter.add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(meter.value())
}

func (c *counter) add(x int) {
	c.Lock()
	defer c.Unlock()

	c.val++
}

func (c *counter) value() int {
	return c.val
}
