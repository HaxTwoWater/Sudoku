package utils

func IsValidMove(board [][]int, row, col, val int) bool {
	if val < 1 || val > 9 {
		return false
	}
	for i := 0; i < 9; i++ {
		if board[row][col] == val {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if board[row][col] == val {
			return false
		}
	}
	bRow := (row / 3) * 3
	bCol := (col / 3) * 3
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[bRow+i][bCol+j] == val {
				return false
			}
		}
	}
	return true
}
