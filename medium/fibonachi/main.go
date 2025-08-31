package main

import (
	"fmt"
)

func fibonachi(n int, ch chan int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go fibonachi(20, ch)

	for num := range ch {
		fmt.Println(num)
	}
}
