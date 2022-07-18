package main

func main() {
	/* ans := longestPalindromeSubseq("aaab")
	fmt.Println(ans) */
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minInsertions(s string) int {

	n := len(s)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				memo[i][j] = memo[i+1][j-1]
			} else {
				memo[i][j] = 1 + max(
					memo[i+1][j],
					memo[i][j-1],
				)
			}
		}
	}

	return memo[0][n-1]
}
