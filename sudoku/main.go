package main

import (
	"fmt"
	"os"
)

func main() {
	// Must receive exactly 9 row arguments (plus program name = 10 total)
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	// Parse input into 9x9 grid
	grid := make([][]int, 9)
	for i := 0; i < 9; i++ {
		arg := os.Args[i+1]
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}
		grid[i] = make([]int, 9)
		for j, ch := range arg {
			if ch == '.' {
				grid[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				grid[i][j] = int(ch - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	// Reject grids with duplicate clues already present
	if !isValidInitial(grid) {
		fmt.Println("Error")
		return
	}

	// Solve and collect solutions (stop early if > 1)
	var solutions [][][]int
	solve(grid, &solutions)

	// A valid sudoku must have exactly ONE solution
	if len(solutions) != 1 {
		fmt.Println("Error")
		return
	}

	// Print the unique solution
	sol := solutions[0]
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(sol[i][j])
		}
		fmt.Println()
	}
}

// isValidInitial checks that all pre-filled numbers don't conflict with each other
func isValidInitial(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] != 0 {
				if !isValidPlacement(grid, i, j, grid[i][j]) {
					return false
				}
			}
		}
	}
	return true
}

// isValidPlacement checks if placing 'num' at (row, col) is legal
func isValidPlacement(grid [][]int, row, col, num int) bool {
	// Check row (skip current column)
	for j := 0; j < 9; j++ {
		if j != col && grid[row][j] == num {
			return false
		}
	}
	// Check column (skip current row)
	for i := 0; i < 9; i++ {
		if i != row && grid[i][col] == num {
			return false
		}
	}
	// Check 3x3 box
	boxRow, boxCol := (row/3)*3, (col/3)*3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {
			if (i != row || j != col) && grid[i][j] == num {
				return false
			}
		}
	}
	return true
}

// solve uses backtracking to find all solutions (stops early after finding 2)
func solve(grid [][]int, solutions *[][][]int) {
	if len(*solutions) > 1 {
		return
	}

	// Find first empty cell
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValidPlacement(grid, i, j, num) {
						grid[i][j] = num
						solve(grid, solutions)
						grid[i][j] = 0 // backtrack
					}
				}
				return // dead end
			}
		}
	}

	// No empty cells: complete solution found
	copyGrid := make([][]int, 9)
	for i := range grid {
		copyGrid[i] = make([]int, 9)
		copy(copyGrid[i], grid[i])
	}
	*solutions = append(*solutions, copyGrid)
}
