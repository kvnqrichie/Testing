package piscine

// FirstRune returns the first rune (character) of the string.
func FirstRune(s string) rune {
	// range over the string gives each rune one by one
	for _, r := range s {
		// return the first rune found
		return r
	}

	// return 0 if the string is empty
	return 0
}
