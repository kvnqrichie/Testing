package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	for {
		// print current combination
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(arr[i] + '0'))
		}

		// check if last combination
		if isLast(arr, n) {
			break
		}

		z01.PrintRune(',')
		z01.PrintRune(' ')

		// generate next combination
		next(arr, n)
	}

	z01.PrintRune('\n')
}

func next(arr []int, n int) {
	i := n - 1

	for i >= 0 {
		if arr[i] < 10-n+i {
			arr[i]++
			for j := i + 1; j < n; j++ {
				arr[j] = arr[j-1] + 1
			}
			return
		}
		i--
	}
}

func isLast(arr []int, n int) bool {
	for i := 0; i < n; i++ {
		if arr[i] != 10-n+i {
			return false
		}
	}
	return true
}
