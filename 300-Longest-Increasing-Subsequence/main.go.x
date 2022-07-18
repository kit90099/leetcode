package main

import "fmt"

func main() {
	ans := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Print(ans)
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			} else {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	max := -1
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
