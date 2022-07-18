package main

import "fmt"

func main() {
	ans := minimumDeleteSum("sea", "eat")
	fmt.Print(ans)
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	memo = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)
	}

	return dp(s1, s2, 0, 0)
}

var memo [][]int

func dp(s1, s2 string, i, j int) int {
	res := 0
	if i == len(s1) {
		for x := j; x < len(s2); x++ {
			res += int(s2[x])
		}
		return res
	}
	if j == len(s2) {
		for x := i; x < len(s1); x++ {
			res += int(s1[x])
		}
		return res
	}

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	if s1[i] == s2[j] {
		memo[i][j] = dp(s1, s2, i+1, j+1)
	} else {
		memo[i][j] = min(
			int(s2[j])+dp(s1, s2, i, j+1),
			int(s1[i])+dp(s1, s2, i+1, j),
		)
	}
	return memo[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
