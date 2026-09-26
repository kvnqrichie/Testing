package piscine

func Atoi(s string) int {
	result := 0
	negative := false
	start := 0

	if len(s) == 0 {
		return 0
	}

	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		negative = true
		start = 1
	}

	for _, c := range s[start:] {
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}

	if negative {
		return -result
	}
	return result
}
