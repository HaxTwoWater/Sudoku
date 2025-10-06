package utils

import "fmt"

func DisplayGrid(grid [][]int) {
	n := len(grid) // supposé 9
	for i := 0; i < n; i++ {
		if i%3 == 0 {
			fmt.Println("+-------+-------+-------+")
		}
		for j := 0; j < n; j++ {
			if j%3 == 0 {
				fmt.Print("| ")
			}
			v := grid[i][j]
			if v == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", v)
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+-------+-------+")
}
