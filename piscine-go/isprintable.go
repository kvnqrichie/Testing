package piscine

// IsPrintable checks whether every character in the string is printable.
// It returns true if all characters are printable ASCII characters,
// otherwise it returns false.
func IsPrintable(s string) bool {
	// Loop through each character in the string.
	// 'c' is a rune (Unicode code point), not just a byte.
	for _, c := range s {
		// Check if the character is outside the printable ASCII range.
		// Printable ASCII characters range from:
		// 32  -> space (' ')
		// 126 -> tilde ('~')
		//
		// Anything less than 32 (like '\n', '\t') or greater than 126
		// is considered non-printable.
		if c < 32 || c > 126 {
			// If we find a non-printable character,
			// we immediately return false.
			return false
		}
	}

	// If we finish the loop without finding any invalid character,
	// it means all characters are printable.
	return true
}
