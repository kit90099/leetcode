package main

import (
	"fmt"
)

func main() {
	ans := superEggDrop(2, 50)
	fmt.Println(ans)
}

var memo [][]int

/* func superEggDrop(k int, n int) int {
	memo = make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = -1
		}
	}

	return dp(k, n)
}

func dp(k int, n int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}

	if memo[k][n] != -1 {
		return memo[k][n]
	}

	res := math.MaxInt
	lo, hi := 1, n
	for lo <= hi {
		mid := (lo + hi) / 2
		broken := dp(k-1, mid-1)
		notBroken := dp(k, n-mid)
		if broken > notBroken {
			hi = mid - 1
			res = min(res, broken+1)
		} else {
			lo = mid + 1
			res = min(res, notBroken+1)
		}
	}
	memo[k][n] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
*/

func superEggDrop(k int, n int) int {
	memo := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		memo[i] = make([]int, n+1)
	}

	m := 0
	for memo[k][m] < n {
		m++
		for x := 1; x <= k; x++ {
			memo[k][x] = memo[k-1][m-1] + memo[k][m-1] + 1
		}
	}

	return memo[k][n]
}
