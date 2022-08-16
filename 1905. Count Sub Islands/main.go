package main

var m, n int

/* func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n = len(grid1), len(grid1[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(grid2, i, j)
			}
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(grid2, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
*/

/* // ver 2
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n = len(grid1), len(grid1[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				t := dfs(grid1, grid2, i, j)
				if t != -1 {
					res++
				}
			}
		}
	}
	return res
}

func dfs(grid1 [][]int, grid2 [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}

	if grid2[i][j] == 0 {
		return 0
	}

	if grid2[i][j] == 1 && grid1[i][j] == 0 {
		return -1
	}
	grid2[i][j] = 0
	if dfs(grid1, grid2, i+1, j) == -1 {
		return -1
	}
	if dfs(grid1, grid2, i-1, j) == -1 {
		return -1
	}
	if dfs(grid1, grid2, i, j-1) == -1 {
		return -1
	}
	if dfs(grid1, grid2, i, j+1) == -1 {
		return -1
	}
	return 0
}
*/
