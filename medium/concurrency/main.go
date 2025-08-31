package main

import (
	"fmt"
	"sync"
)

var counter int = 0

var (
	counter2 int
	mu       sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter (without mutex):", counter)

	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < 100; j++ {
				mu.Lock()
				counter2++
				mu.Unlock()
			}
		}()
	}

	wg2.Wait()
	fmt.Println("Final counter (with mutex):", counter2)
}
