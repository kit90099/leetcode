package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	for i := 0; i < n; i++ {
		l, r := 0, n-1
		for l < r {
			temp := matrix[i][l]
			matrix[i][l] = matrix[i][r]
			matrix[i][r] = temp

			l++
			r--
		}
	}
}
