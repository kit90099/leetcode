package main

import (
	"fmt"
	"strconv"
)

var m, n int

func main() {

	ans := numDistinctIslands([][]int{{0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0}, {0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0}, {0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0}, {1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}})
	fmt.Println(ans)
}

func numDistinctIslands(grid [][]int) int {
	for _, i := range grid {
		for _, j := range i {
			fmt.Print(strconv.Itoa(j) + ",")
		}
		fmt.Println("")
	}

	m, n = len(grid), len(grid[0])

	iMap := make(map[string]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res := ""
				dfs(grid, i, j, 9, &res)
				fmt.Printf("i:%d, j:%d, res: %s\n", i, j, res)
				iMap[res] = true
			}
		}
	}

	return len(iMap)
}

/**
*	up: 1
*	down: 2
*	left: 3
*	right: 4
**/
func dfs(grid [][]int, i, j int, dir int, res *string) {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0
	(*res) += "," + strconv.Itoa(dir)
	dfs(grid, i-1, j, 1, res) // up
	dfs(grid, i+1, j, 2, res) // down
	dfs(grid, i, j-1, 3, res) // left
	dfs(grid, i, j+1, 4, res) // right

	(*res) += "," + strconv.Itoa(-dir)
}
