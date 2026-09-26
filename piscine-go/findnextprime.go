package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	// Start from nb, ensure we begin on an odd number
	if nb%2 == 0 {
		nb++
	}
	for !isPrime(nb) {
		nb += 2
	}
	return nb
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	// Check divisors of the form 6k ± 1 up to √n
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
