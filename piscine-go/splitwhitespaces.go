package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	word := ""

	for _, r := range str {
		if isSeparator(r) {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}

func isSeparator(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
