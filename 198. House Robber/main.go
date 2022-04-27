package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2}))
}

// ver 1
/* func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)

	for i := 2; i < n+2; i++ {
		dp[i] = max(dp[i-2]+nums[i-2], dp[i-1])
	}

	return dp[n+1]
} */

// ver 2
/* func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)

	for i := n - 1; i >= 0; i-- {
		dp[i] = max(dp[i+2]+nums[i], dp[i+1])
	}

	return dp[0]
} */

//version 3
func rob(nums []int) int {
	n := len(nums)
	revenueLast, revenueLastLast := 0, 0

	for i := 0; i < n; i++ {
		temp := revenueLast
		revenueLast = max(revenueLastLast+nums[i], revenueLast)
		revenueLastLast = temp
	}

	return revenueLast
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
