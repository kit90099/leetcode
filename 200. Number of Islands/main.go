package main

func numIslands(grid [][]byte) int {
	idx = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	res := 0
	m, n = len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
				dfs(&grid, i, j)
			}
		}
	}
	return res
}

var idx [][]int
var m, n int

func dfs(nums *[][]byte, i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if (*nums)[i][j] == 0 {
		return
	}

	(*nums)[i][j] = 0
	for _, offset := range idx {
		dfs(nums, i+offset[0], j+offset[1])
	}
}
