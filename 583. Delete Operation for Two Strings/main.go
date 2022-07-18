package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
			} else {
				memo[i][j] = max(
					memo[i-1][j],
					memo[i][j-1],
				)
			}
		}
	}

	common := memo[m][n]
	return (m - common) + (n - common)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
