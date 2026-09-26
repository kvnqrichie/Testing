package piscine

// Join concatenates all strings in the slice 'strs'
// using the separator 'sep' between each element.
func Join(strs []string, sep string) string {
	// If the slice is empty, return an empty string immediately.
	if len(strs) == 0 {
		return ""
	}

	// Start the result with the first element of the slice.
	// This avoids adding a separator at the beginning.
	result := strs[0]

	// Loop through the slice starting from index 1
	// because index 0 is already added to result.
	for i := 1; i < len(strs); i++ {

		// Add the separator first
		result += sep

		// Then add the current string
		result += strs[i]
	}

	// Return the final concatenated string
	return result
}
