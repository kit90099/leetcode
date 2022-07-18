package main

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i < n+1; i++ {
		temp := make([]int, amount+1)
		temp[0] = 1
		for j := 1; j < amount+1; j++ {
			if j-coins[i-1] >= 0 {
				temp[j] = dp[j] + temp[j-coins[i-1]]
			} else {
				temp[j] = dp[j]
			}
		}
		dp = temp
	}

	return dp[amount]
}
