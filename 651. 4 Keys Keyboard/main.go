package main

import "fmt"

func main() {
	ans := maxA(7)
	fmt.Println(ans)
}

func maxA(n int) int {
	memo := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		memo[i] = memo[i-1] + 1
		for j := 2; j < i; j++ {
			memo[i] = max(memo[i], memo[j-2]*(i-j+1))
		}
	}
	return memo[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
