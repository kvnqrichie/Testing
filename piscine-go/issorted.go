package piscine

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	order := 0
	for i := 1; i < len(a); i++ {
		cmp := f(a[i-1], a[i])
		if order == 0 {
			order = cmp
		}
		if cmp != 0 && (cmp > 0) != (order > 0) {
			return false
		}
	}
	return true
}
