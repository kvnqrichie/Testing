package piscine

func IsAlpha(s string) bool {
	for _, c := range s {

		// Check lowercase letters
		if c >= 'a' && c <= 'z' {
			continue
		}

		// Check uppercase letters
		if c >= 'A' && c <= 'Z' {
			continue
		}

		// Check digits (0–9)
		if c >= '0' && c <= '9' {
			continue
		}

		// If it's not a letter or number, return false
		return false
	}

	// If we never found an invalid character, return true
	return true
}
