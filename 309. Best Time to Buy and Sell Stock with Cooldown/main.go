package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{1}))
}

func maxProfit(prices []int) int {
	n := len(prices)

	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[0][0] = 0
			dp[0][1] = -prices[0]
			continue
		}

		if i-2 == -1 {
			dp[1][0] = max(0, dp[0][1]+prices[1])
			dp[1][1] = max(dp[0][1], dp[i-1][0]-prices[1])
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}

	return dp[n-1][0]
}

func max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}
