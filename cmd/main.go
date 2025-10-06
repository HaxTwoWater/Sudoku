package main

import (
	"sudoku/pkg/utils"
)

func main() {
	m := [][]int{
		{1, 0, 0, 4, 0, 0, 0, 0, 9},
		{4, 0, 9, 1, 7, 8, 2, 3, 6},
		{8, 0, 7, 2, 0, 3, 1, 0, 0},
		{3, 0, 5, 6, 8, 1, 9, 0, 4},
		{0, 0, 1, 5, 0, 2, 3, 8, 7},
		{2, 8, 4, 7, 3, 9, 6, 0, 0},
		{9, 4, 0, 0, 1, 5, 7, 6, 0},
		{5, 0, 6, 0, 0, 7, 4, 9, 8},
		{7, 0, 8, 0, 0, 4, 5, 1, 0},
	}

	utils.PrintPolish()
	utils.DisplayGrid(m)
	utils.PrintPolish()

}
