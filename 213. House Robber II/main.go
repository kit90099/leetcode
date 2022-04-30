package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1}))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(r(nums[:n-1]), r(nums[1:]))
}

func r(nums []int) int {
	n := len(nums)

	dp_last, dp_lastLast := 0, 0

	for i := 0; i < n; i++ {
		temp := dp_last
		dp_last = max(dp_lastLast+nums[i], dp_last)
		dp_lastLast = temp
	}

	return dp_last
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
