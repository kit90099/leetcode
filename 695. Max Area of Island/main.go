package main

var m, n int
var idxs [][]int = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func maxAreaOfIsland(grid [][]int) int {
	m, n = len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(dfs(grid, i, j), res)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}

	if grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	res := 0
	for _, idx := range idxs {
		res += dfs(grid, i+idx[0], j+idx[1])
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
