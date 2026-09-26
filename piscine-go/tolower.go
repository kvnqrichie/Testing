package piscine

// ToLower converts every uppercase ASCII letter in a string to lowercase
func ToLower(s string) string {
	// Convert string to a slice of runes/bytes so we can modify characters
	result := []rune(s)

	// Loop through each character in the string
	for i := 0; i < len(result); i++ {
		// If character is between 'A' and 'Z'
		if result[i] >= 'A' && result[i] <= 'Z' {
			// Add 32 to convert uppercase ASCII to lowercase
			result[i] = result[i] + 32
		}
	}

	// Convert rune slice back to string and return it
	return string(result)
}
