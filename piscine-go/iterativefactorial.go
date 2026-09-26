package piscine

func IterativeFactorial(n int) int {
	if n < 0 {
		return 0
	}

	result := 1
	for i := 2; i <= n; i++ {
		if result > (1<<63-1)/i {
			return 0
		}
		result *= i
	}
	return result
}
