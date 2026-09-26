package piscine

func LastRune(s string) rune {
	// Convert the string into a slice of runes
	// so each Unicode character is handled correctly
	r := []rune(s)

	// Return the last rune in the slice
	return r[len(r)-1]
}
