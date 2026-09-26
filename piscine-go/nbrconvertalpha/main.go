package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isNumber(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	upper := false
	start := 0

	if args[0] == "--upper" {
		upper = true
		start = 1
	}

	for i := start; i < len(args); i++ {
		n, ok := isNumber(args[i])

		if !ok || n < 1 || n > 26 {
			z01.PrintRune(' ')
		} else {
			if upper {
				z01.PrintRune(rune('A' + n - 1))
			} else {
				z01.PrintRune(rune('a' + n - 1))
			}
		}
	}

	z01.PrintRune('\n')
}
