package main

import (
	"fmt"
	"os"
)

func isInsertFlag(s string) (string, bool) {
	if len(s) >= 9 && s[:9] == "--insert=" {
		return s[9:], true
	}
	if len(s) >= 3 && s[:3] == "-i=" {
		return s[3:], true
	}
	return "", false
}

func isOrderFlag(s string) bool {
	return s == "--order" || s == "-o"
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func sortRunes(s string) string {
	r := []rune(s)

	for i := 0; i < len(r)-1; i++ {
		for j := 0; j < len(r)-i-1; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}

	return string(r)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	var input string
	var insert string
	order := false

	for _, arg := range args {
		if val, ok := isInsertFlag(arg); ok {
			insert = val
		} else if isOrderFlag(arg) {
			order = true
		} else {
			input = arg
		}
	}

	result := input + insert

	if order {
		result = sortRunes(result)
	}

	fmt.Println(result)
}
