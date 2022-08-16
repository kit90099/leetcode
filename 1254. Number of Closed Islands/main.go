package main

import "fmt"

func main() {
	ans := closedIsland([][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}})
	fmt.Println(ans)
}

func closedIsland(grid [][]int) int {
	m, n = len(grid), len(grid[0])
	idxs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	for i := 0; i < m; i++ {
		dfs(&grid, i, 0)
		dfs(&grid, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(&grid, 0, j)
		dfs(&grid, m-1, j)
	}

	res := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(&grid, i, j)
			}
		}
	}

	return res
}

var m, n int
var idxs [][]int

func dfs(grid *[][]int, i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if (*grid)[i][j] == 1 {
		return
	}

	(*grid)[i][j] = 1

	for _, idx := range idxs {
		dfs(grid, i+idx[0], j+idx[1])
	}
}
