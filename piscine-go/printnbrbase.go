package piscine // Defines the package name required by the 01-edu project

import "github.com/01-edu/z01" // Library used to print single characters (runes)

// PrintNbrBase prints an integer in a given base string
func PrintNbrBase(nbr int, base string) {
	// -----------------------------
	// STEP 1: Validate the base
	// -----------------------------

	if len(base) < 2 { // A valid base must have at least 2 characters
		printNV() // If invalid, print "NV"
		return    // Stop execution immediately
	}

	seen := make(map[rune]bool) // Map used to ensure all base characters are unique

	for _, ch := range base { // Loop through each character in the base string
		if ch == '+' || ch == '-' { // Base cannot contain '+' or '-'
			printNV() // Invalid base → print NV
			return
		}
		if seen[ch] { // If character already appeared
			printNV() // Duplicate character → invalid base
			return
		}
		seen[ch] = true // Mark character as seen
	}

	// -----------------------------
	// STEP 2: Handle special case (0)
	// -----------------------------

	if nbr == 0 { // If number is zero
		z01.PrintRune(rune(base[0])) // Print first character of base
		return                       // Done
	}

	// -----------------------------
	// STEP 3: Handle sign safely
	// -----------------------------

	var sign bool // Will store if number is negative
	var n uint64  // We use unsigned int to safely handle MinInt64

	if nbr < 0 {
		sign = true // Remember that number was negative

		// IMPORTANT FIX:
		// We cannot do: nbr = -nbr because MinInt64 overflows in Go
		// Instead we convert safely using uint64 trick:
		// -(nbr + 1) avoids overflow, then we correct by +1
		n = uint64(-(nbr + 1)) + 1
	} else {
		n = uint64(nbr) // Safe conversion for positive numbers
	}

	// -----------------------------
	// STEP 4: Convert number to the given base
	// -----------------------------

	baseLen := uint64(len(base)) // Length of base used for conversions
	var result []rune            // Stores digits in reverse order

	for n > 0 { // Repeat until number becomes 0
		remainder := n % baseLen                       // Get remainder → current digit in base
		result = append(result, rune(base[remainder])) // Convert digit to base character
		n = n / baseLen                                // Reduce number for next iteration
	}

	// -----------------------------
	// STEP 5: Print sign if needed
	// -----------------------------

	if sign {
		z01.PrintRune('-') // Print minus sign for negative numbers
	}

	// -----------------------------
	// STEP 6: Print result in correct order
	// -----------------------------

	for i := len(result) - 1; i >= 0; i-- { // Reverse loop (because digits were stored backwards)
		z01.PrintRune(result[i]) // Print each digit
	}
}

// -----------------------------
// Helper function: prints "NV"
// -----------------------------

func printNV() {
	z01.PrintRune('N') // Print 'N'
	z01.PrintRune('V') // Print 'V'
}
