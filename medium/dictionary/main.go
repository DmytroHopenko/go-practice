package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter string: ")
	str, _ := reader.ReadString('\n')

	dict := map[rune]int{}

	for _, char := range str {
		if char == ' ' || char == '\n' {
			continue
		}
		dict[char]++
	}

	for char, count := range dict {
		fmt.Printf("%c -> %d\n", char, count)
	}
}
