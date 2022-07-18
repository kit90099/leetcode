package main

import "fmt"

func main() {
	ans := canPartition([]int{1, 5, 11, 5})
	fmt.Println(ans)
}

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2

	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][sum]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
