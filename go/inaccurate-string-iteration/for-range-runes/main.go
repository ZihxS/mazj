package main

import "fmt"

func main() {
	str := "こんにちは"
	runes := []rune(str)
	for i, char := range runes {
		fmt.Printf("Index %d: %c\n", i, char)
	}
}
