package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		memo[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		memo[0][i] = i
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				memo[i][j] = memo[i-1][j-1]
			} else {
				memo[i][j] = min3(
					memo[i][j-1]+1,   //insert
					memo[i-1][j]+1,   //remove
					memo[i-1][j-1]+1, //replace
				)
			}
		}
	}
	return memo[m][n]
}

func min3(a, b, c int) int {
	if a < b {
		return min(a, c)
	} else {
		return min(b, c)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
