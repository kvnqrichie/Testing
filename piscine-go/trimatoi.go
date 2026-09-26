package piscine

func TrimAtoi(s string) int {
	// This will store the final number
	result := 0

	// This flag determines if the number is negative
	negative := false

	// This flag tells us if we've started reading digits
	started := false

	for i := 0; i < len(s); i++ {
		c := s[i] // get current character

		// If we find a '-' before any number, mark as negative
		if c == '-' && !started {
			negative = true
			continue
		}

		// If we find a digit (ASCII '0' to '9')
		if c >= '0' && c <= '9' {
			started = true // we have started reading numbers

			// convert character digit to integer value
			digit := int(c - '0')

			// build the number step by step
			result = result*10 + digit
		}
		// If it's not a digit and not '-', we simply ignore it
	}

	// If no digits were found, return 0
	if !started {
		return 0
	}

	// Apply negative sign if needed
	if negative {
		return -result
	}

	return result
}
