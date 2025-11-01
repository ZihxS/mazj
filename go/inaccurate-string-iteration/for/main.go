package main

import "fmt"

func main() {
	str := []rune("こんにちは")
	for i := range str {
		fmt.Printf("%c", str[i])
	}
}
