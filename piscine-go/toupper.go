package piscine // declares the package name where this function belongs

// ToUpper takes a string and returns a new string with all lowercase letters converted to uppercase
func ToUpper(s string) string {
	result := "" // this will store the final uppercase version of the string

	for _, c := range s { // loop through each character (rune) in the string

		if c >= 'a' && c <= 'z' { // check if the character is a lowercase letter (a to z)

			c = c - 32 // convert lowercase ASCII letter to uppercase by subtracting 32
			// (e.g., 'a' (97) becomes 'A' (65))
		}

		result += string(c) // add the (possibly converted) character to the result string
	}

	return result // return the final uppercase string
}
