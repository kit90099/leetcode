package main

import "fmt"

func longestWPI(hours []int) int {

	for i, h := range hours {
		if h > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}

	preSum := []int{0}
	for _, h := range hours {
		preSum = append(preSum, preSum[len(preSum)-1]+h)
	}

	stk := []int{}
	for i, p := range preSum {
		if len(stk) == 0 || preSum[stk[len(stk)-1]] > p {
			stk = append(stk, i)
		}
	}

	res := 0
	for j := len(preSum) - 1; j >= 0; j-- {
		for len(stk) > 0 && preSum[stk[len(stk)-1]] < preSum[j] {
			res = max(res, j-stk[len(stk)-1])
			stk = stk[:len(stk)-1]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	fmt.Println(longestWPI([]int{0, 0, 0, 9, 9, 9, 0}))
}
