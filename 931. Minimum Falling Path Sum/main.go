package main

import "math"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = 0
	}

	for i := 0; i < n; i++ {
		temp := make([]int, n)
		for j := 0; j < n; j++ {
			if j == 0 {
				temp[j] = min(memo[j], memo[j+1]) + matrix[i][j]
			} else if j == n-1 {
				temp[j] = min(memo[j], memo[j-1]) + matrix[i][j]
			} else {
				temp[j] = min(
					memo[j-1],
					min(
						memo[j],
						memo[j+1],
					),
				) + matrix[i][j]
			}

		}
		memo = temp
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		if memo[i] < res {
			res = memo[i]
		}
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
