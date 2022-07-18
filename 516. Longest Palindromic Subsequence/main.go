package main

import "fmt"

func longestPalindromeSubseq1(s string) int {
	return dp(s, 0, len(s)-1)
}

func dp(s string, i, j int) int {
	if i > j {
		return 0
	}

	if i == j {
		return 1
	}

	if s[i] == s[j] {
		return 2 + dp(s, i+1, j-1)
	} else {
		return max3(
			dp(s, i+1, j),
			dp(s, i, j-1),
			dp(s, i+1, j-1),
		)
	}
}

func max3(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	} else {
		return c
	}
}
func main() {
	ans := longestPalindromeSubseq("aaab")
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestPalindromeSubseq(s string) int {

	n := len(s)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][i] = 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				memo[i][j] = 2 + memo[i+1][j-1]
			} else {
				memo[i][j] = max(
					memo[i+1][j],
					memo[i][j-1],
				)
			}
		}
	}

	return memo[0][n-1]
}
