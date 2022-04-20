package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		min := amount + 1

		for _, j := range coins {
			if i-j < 0 {
				continue
			}

			temp := dp[i-j] + 1

			if temp < min {
				min = temp
			}
		}

		dp[i] = min
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func main() {
	fmt.Println(coinChange([]int{2}, 3))
}
