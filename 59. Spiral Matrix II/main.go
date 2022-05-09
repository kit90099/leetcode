package main

import "fmt"

func main() {
	ans := generateMatrix(3)
	fmt.Println(ans)
}

func generateMatrix(n int) [][]int {
	upper, lower, left, right := 0, n-1, 0, n-1

	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, n))
	}

	count := 1
	for count <= n*n {
		if upper <= lower {
			for i := left; i <= right; i++ {
				ans[upper][i] = count
				count++
			}
			upper++
		}

		if left <= right {
			for i := upper; i <= lower; i++ {
				ans[i][right] = count
				count++
			}
			right--
		}

		if upper <= lower {
			for i := right; i >= left; i-- {
				ans[lower][i] = count
				count++
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= upper; i-- {
				ans[i][left] = count
				count++
			}
			left++
		}
	}

	return ans
}
