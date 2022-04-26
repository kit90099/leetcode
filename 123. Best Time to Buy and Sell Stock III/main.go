package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func maxProfit(nums []int) int {
	n := len(nums)

	dp := make([][2 + 1][2]int, n)

	for i := 0; i < n; i++ {
		for k := 1; k <= 2; k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -nums[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+nums[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-nums[i])
		}
	}

	return dp[n-1][2][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
