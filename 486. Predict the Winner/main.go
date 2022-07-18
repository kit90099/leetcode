package main

import "fmt"

func main() {
	ans := PredictTheWinner([]int{0})
	fmt.Println(ans)
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	memo := make([][]inner, n)

	for i := 0; i < n; i++ {
		memo[i] = make([]inner, n)
		memo[i][i] = inner{
			first:  nums[i],
			second: 0,
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			left := nums[i] + memo[i+1][j].second
			right := nums[j] + memo[i][j-1].second

			if left > right {
				memo[i][j].first = left
				memo[i][j].second = memo[i+1][j].first
			} else {
				memo[i][j].first = right
				memo[i][j].second = memo[i][j-1].first
			}
		}
	}
	return memo[0][n-1].first > memo[0][n-1].second
}

type inner struct {
	first  int
	second int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
