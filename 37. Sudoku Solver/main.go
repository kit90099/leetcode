package main

import "fmt"

func main() {
	solveSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}})
	fmt.Println("test")
}

func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}

func backtrack(board [][]byte, i, j int) bool {
	m, n := 9, 9

	if j == n {
		return backtrack(board, i+1, 0)
	}

	if i == m {
		return true
	}

	if board[i][j] != '.' {
		return backtrack(board, i, j+1)
	}

	var c byte = '1'
	for ; c <= '9'; c++ {
		if !valid(board, i, j, c) {
			continue
		}

		board[i][j] = c
		if backtrack(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func valid(board [][]byte, i, j int, c byte) bool {
	for x := 0; x < 9; x++ {
		if board[i][x] == c {
			return false
		}

		if board[x][j] == c {
			return false
		}

		if board[(i/3)*3+x/3][(j/3)*3+x%3] == c {
			return false
		}
	}
	return true
}
