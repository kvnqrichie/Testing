package piscine // defines the package name so it can be used in the test program

func IsNumeric(s string) bool { // function that takes a string and returns true/false

	for _, ch := range s { // loop through each character in the string

		if ch < '0' || ch > '9' { // check if character is NOT between '0' and '9'
			return false // if any non-numeric character is found, return false immediately
		}
	}

	return true // if all characters are digits, return true
}
