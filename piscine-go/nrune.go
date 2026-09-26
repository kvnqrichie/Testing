package piscine

func NRune(s string, n int) rune {
	// If n is less than or equal to 0,
	// it is not a valid position.
	if n <= 0 {
		return 0
	}

	// Convert the string into a slice of runes.
	// This allows correct handling of Unicode characters.
	runes := []rune(s)

	// Check if n is bigger than the number of runes.
	// If yes, return 0 because the position does not exist.
	if n > len(runes) {
		return 0
	}

	// Return the rune at position n-1.
	// We use n-1 because slice indexes start at 0.
	return runes[n-1]
}
