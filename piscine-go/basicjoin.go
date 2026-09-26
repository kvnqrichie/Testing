package piscine // declares the package name so this function can be imported and used in main

// BasicJoin takes a slice of strings and returns them concatenated into one string
func BasicJoin(elems []string) string { // function definition with a string slice parameter and string return type

	result := "" // creates an empty string that will store the final concatenated result

	for i := 0; i < len(elems); i++ { // loop through each index of the slice from 0 to length-1
		result += elems[i] // add the current string element to the result string (concatenation)
	}

	return result // return the final combined string after the loop finishes
}
