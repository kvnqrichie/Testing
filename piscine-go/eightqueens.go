package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var isSafe func(col int, row int) bool
	isSafe = func(col int, row int) bool {
		for i := 0; i < col; i++ {
			if board[i] == row || abs(board[i]-row) == abs(i-col) {
				return false
			}
		}
		return true
	}

	printSolution := func() {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(board[i] + '1'))
		}
		z01.PrintRune('\n')
	}

	var solve func(col int)
	solve = func(col int) {
		if col == 8 {
			printSolution()
			return
		}

		for row := 0; row < 8; row++ {
			if isSafe(col, row) {
				board[col] = row
				solve(col + 1)
			}
		}
	}

	solve(0)
}
