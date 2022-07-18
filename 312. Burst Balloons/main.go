package main

import (
	"fmt"
)

func main() {
	ans := maxCoins([]int{1, 5})
	fmt.Println(ans)
}

/* func maxCoins(nums []int) int {
	nNums := []int{1}
	nNums = append(nNums, nums...)
	nNums = append(nNums, 1)
	return dp(nNums)
}

func dp(nums []int) int {
	if len(nums) == 3 {
		return nums[1]
	}

	res := math.MinInt
	for i := 1; i < len(nums)-1; i++ {
		temp := []int{}
		temp = append(temp, nums[:i]...)
		temp = append(temp, nums[i+1:]...)
		res = max(res, dp(temp)+nums[i-1]*nums[i]*nums[i+1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
} */

func maxCoins(nums []int) int {
	n := len(nums)

	nNums := []int{1}
	nNums = append(nNums, nums...)
	nNums = append(nNums, 1)

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for i := n; i >= 0; i-- {
		for j := 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][k]+dp[k][j]+(nNums[i]*nNums[k]*nNums[j]), dp[i][j])
			}
		}
	}

	return dp[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
