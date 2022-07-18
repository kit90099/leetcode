package main

import (
	"fmt"
	"math"
)

func main() {
	ans := findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1)
	fmt.Println(ans)
}

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for _, x := range nums {
		sum += x
	}

	if sum < int(math.Abs(float64(target))) || (sum+target)%2 == 1 {
		return 0
	}

	newTarget := (sum + target) / 2

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, newTarget+1)
	}
	dp[0][0] = 1

	for i := 1; i < n+1; i++ {
		for j := 0; j < newTarget+1; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][newTarget]
}
