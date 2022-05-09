package main

import "fmt"

func main() {
	ans := spiralOrder([][]int{})
	fmt.Println(ans)
}

func spiralOrder(matrix [][]int) []int {
	height, width := len(matrix), len(matrix[0])
	upper, lower, left, right := 0, height-1, 0, width-1
	ans := []int{}
	for len(ans) < height*width {

		// left to right
		if upper <= lower {
			for i := left; i <= right; i++ {
				ans = append(ans, matrix[upper][i])
			}
			upper++
		}

		// upper to lower
		if left <= right {
			for i := upper; i <= lower; i++ {
				ans = append(ans, matrix[i][right])
			}
			right--
		}

		// right to left
		if upper <= lower {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[lower][i])
			}
			lower--
		}

		// lower to upper
		if left <= right {
			for i := lower; i >= upper; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}

	return ans
}
