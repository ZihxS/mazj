package main

import "fmt"

func main() {
	str := "こんにちは"
	for i, char := range str {
		fmt.Printf("Index %d: %c\n", i, char)
	}
}
