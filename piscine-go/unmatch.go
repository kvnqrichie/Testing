package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, v := range a {
		counts[v]++
	}

	for _, v := range a {
		if counts[v]%2 != 0 {
			return v
		}
	}

	return -1
}
