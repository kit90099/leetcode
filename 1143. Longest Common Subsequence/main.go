package main

import "fmt"

func main() {
	ans := longestCommonSubsequence("bsbininm", "jmjkbkjkv")
	fmt.Print(ans)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)
	fmt.Printf("n1:%d, n2:%d\n", n1, n2)
	memo := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		memo[i] = make([]int, n2+1)
	}

	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			if text1[i-1] == text2[j-1] {
				memo[i][j] = 1 + memo[i-1][j-1]
			} else {
				memo[i][j] = max(
					memo[i-1][j],
					memo[i][j-1],
				)
			}

			fmt.Printf("memo[%d][%d] %d\n", i, j, memo[i][j])
		}
	}

	return memo[n1][n2]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
