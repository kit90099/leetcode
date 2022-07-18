package main

import "math"

var memo [][]int

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	return dp(dungeon, 0, 0)
}

func dp(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i == m-1 && j == n-1 {
		if grid[i][j] >= 0 {
			return 1
		} else {
			return -grid[i][j] + 1
		}
	}

	if i == m || j == n {
		return math.MaxInt
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	res := min(
		dp(grid, i+1, j),
		dp(grid, i, j+1),
	) - grid[i][j]

	if res <= 0 {
		memo[i][j] = 1
		return 1
	} else {
		memo[i][j] = res
		return res
	}

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
