package piscine

func IsLower(s string) bool {
	// Loop through each character in the string
	for _, ch := range s {
		// Check if the character is NOT between 'a' and 'z'
		if ch < 'a' || ch > 'z' {
			// If any character is outside lowercase range, return false
			return false
		}
	}

	// If all characters are lowercase, return true
	return true
}
