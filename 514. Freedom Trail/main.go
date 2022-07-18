package main

import (
	"fmt"
	"math"
)

func main() {
	ans := findRotateSteps("godding", "gd")
	fmt.Println(ans)
}

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	memo = make([][]int, m)
	charMap = make(map[byte][]int, 0)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		val, exist := charMap[ring[i]]
		if exist {
			val = append(val, i)
			charMap[ring[i]] = val
		} else {
			list := []int{i}
			charMap[ring[i]] = list
		}
	}

	return dp(ring, key, 0, 0)
}

var memo [][]int
var charMap map[byte][]int

func dp(ring, key string, i, j int) int {
	if j == len(key) {
		return 0
	}

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	m := len(ring)

	res := math.MaxInt
	list := charMap[key[i]]
	for _, c := range list {
		delta := abs(i - c)
		delta = min(delta, m-delta)

		sub := dp(ring, key, c, j+1)
		res = min(res, 1+sub+delta)
	}
	memo[i][j] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
