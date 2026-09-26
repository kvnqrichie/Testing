package piscine

// Compare compares two strings lexicographically.
// It returns:
// 0  if both strings are equal
// -1 if a < b
// 1  if a > b
func Compare(a, b string) int {
	// Check if both strings are exactly the same
	if a == b {
		return 0
	}

	// If string a comes before string b alphabetically
	if a < b {
		return -1
	}

	// Otherwise, string a comes after string b
	return 1
}
