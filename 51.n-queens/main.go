package main

import (
	"fmt"
)

func main() {
	solveNQueens(5)
	fmt.Print(ans)
}

var ans [][]string = make([][]string, 0)

func solveNQueens(n int) [][]string {
	ans = make([][]string, 0)

	init := make([]string, 0)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			str += "."
		}
		init = append(init, str)
	}

	backtrack(&init, 0, n)

	return ans
}

func backtrack(arr *[]string, row int, size int) {

	if row == len(*arr) {
		temp := make([]string, len(*arr))
		copy(temp, *arr)
		ans = append(ans, temp)
		return
	}

	original := (*arr)[row]
	for i := 0; i < len(*arr); i++ {
		if row == 0 || isOk(arr, row, i, size) {
			(*arr)[row] = (*arr)[row][:i] + "Q" + (*arr)[row][i+1:]
			backtrack(arr, row+1, size)
			(*arr)[row] = original
		}
	}

}

func isOk(arr *[]string, row int, nextPos int, size int) bool {
	leftPtr := nextPos - 1
	rightPtr := nextPos + 1
	for i := row - 1; i >= 0; i-- {
		if (*arr)[i][nextPos] == 'Q' {
			return false
		}

		if leftPtr >= 0 && (*arr)[i][leftPtr] == 'Q' {
			return false
		}

		if rightPtr < size && (*arr)[i][rightPtr] == 'Q' {
			return false
		}

		leftPtr--
		rightPtr++
	}

	return true
}
