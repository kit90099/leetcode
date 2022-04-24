package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp_0, dp_1 := 0, math.MinInt
	for i := 0; i < n; i++ {
		dp_0 = max(dp_0, dp_1+prices[i])
		dp_1 = max(dp_1, -prices[i])
	}

	return dp_0
}

func max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}
