package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp_0, dp_1 := 0, -99999999

	for i := 0; i < n; i++ {
		temp := dp_0
		dp_0 = max(dp_0, dp_1+prices[i]-fee)
		dp_1 = max(dp_1, temp-prices[i])
	}
	return dp_0
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
