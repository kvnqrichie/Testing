package piscine // Defines the package name (as required by the exercise environment)

// AtoiBase converts a string number `s` from a given `base` into an integer
func AtoiBase(s string, base string) int {
	// Step 1: Validate base length (must be at least 2 characters)
	if len(base) < 2 {
		return 0 // invalid base → return 0
	}

	// Step 2: Check for invalid characters and duplicates in base
	baseMap := make(map[rune]int) // map to store each base character and its value index

	for i, ch := range base { // iterate over each character in base string
		if ch == '+' || ch == '-' { // base cannot contain '+' or '-'
			return 0 // invalid base
		}
		if _, exists := baseMap[ch]; exists { // check if character already exists (duplicate check)
			return 0 // invalid base due to repeated character
		}
		baseMap[ch] = i // store character with its numeric value (position in base)
	}

	baseLen := len(base) // store base length for later calculations

	// Step 3: Convert string `s` to integer using the base
	result := 0 // final numeric result
	sign := 1   // sign handling (not strictly needed per problem, but safe practice)

	for _, ch := range s { // iterate through each character of input string
		value, exists := baseMap[ch] // get numeric value of character in base
		if !exists {                 // if character is not in base, input is invalid
			return 0 // invalid string number
		}
		result = result*baseLen + value // accumulate result in positional numeral system
		sign = 1                        // keep sign positive (problem states no negatives)
	}

	return result * sign // return final computed integer
}
