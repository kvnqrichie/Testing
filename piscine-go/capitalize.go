package piscine // package name (as required by the exercise environment)

// Capitalize takes a string and returns a new string
// where the first letter of each word is uppercase
// and all other letters are lowercase.
func Capitalize(s string) string {
	// Convert the string into a slice of runes
	// so we can properly handle Unicode characters
	runes := []rune(s)

	// This will store the final result
	result := []rune{}

	// This flag tells us if the next letter is the start of a new word
	newWord := true

	// Loop through each character (rune) in the input string
	for _, r := range runes {

		// Check if the character is a letter or digit (part of a word)
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {

			// If we are at the start of a word
			if newWord {
				// If it's a lowercase letter, convert it to uppercase
				if r >= 'a' && r <= 'z' {
					r = r - 32 // ASCII conversion from lowercase to uppercase
				}
			} else {
				// If it's not the start of a word and it's uppercase,
				// convert it to lowercase
				if r >= 'A' && r <= 'Z' {
					r = r + 32 // ASCII conversion from uppercase to lowercase
				}
			}

			// After processing, we are no longer at a new word
			newWord = false

		} else {
			// If the character is NOT alphanumeric (like space, punctuation, +, ?, etc.)
			// the next character will start a new word
			newWord = true
		}

		// Add the processed character to the result
		result = append(result, r)
	}

	// Convert rune slice back to string and return it
	return string(result)
}
