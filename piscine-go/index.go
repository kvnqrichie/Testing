package piscine

func Index(s string, toFind string) int {
	// Loop through the string s
	for i := 0; i <= len(s)-len(toFind); i++ {
		// Check if the part of s matches toFind
		if s[i:i+len(toFind)] == toFind {
			// Return the index where the match starts
			return i
		}
	}

	// Return -1 if no match is found
	return -1
}
