package utils

import "fmt"

func ScanUserInput() (int, int) {
	a, b := 0, 0
	PrintPolish()
	fmt.Println("Enter your first number for the row: (1-9)")
	fmt.Scan(&a)
	fmt.Println("Enter your second number for the column: (1-9)")
	fmt.Scan(&b)
	PrintPolish()
	return a - 1, b - 1
}
