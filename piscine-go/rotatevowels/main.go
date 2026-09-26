package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func toLower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	// collect vowels
	var vowels []rune

	for _, arg := range args {
		for _, c := range arg {
			if isVowel(c) {
				vowels = append(vowels, c)
			}
		}
	}

	// if no vowels → print original args
	if len(vowels) == 0 {
		for i, arg := range args {
			for _, c := range arg {
				z01.PrintRune(c)
			}
			if i != len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// mirror vowels
	i, j := 0, len(vowels)-1

	for k := 0; k < len(vowels)/2; k++ {
		vowels[i], vowels[j] = vowels[j], vowels[i]
		i++
		j--
	}

	// rebuild strings
	v := 0

	for ai, arg := range args {
		for _, c := range arg {
			if isVowel(c) {
				z01.PrintRune(vowels[v])
				v++
			} else {
				z01.PrintRune(c)
			}
		}
		if ai != len(args)-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
