package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	start := 0

	for i, char := range name {
		if char == '/' {
			start = i + 1
		}
	}

	for _, char := range name[start:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
