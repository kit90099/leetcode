package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	stk := []int{}
	res := 0

	// construct stack
	for i, h := range heights {
		for len(stk) > 0 && h < heights[stk[len(stk)-1]] {
			curr := heights[stk[len(stk)-1]]
			stk = stk[:len(stk)-1]
			for len(stk) > 0 && heights[stk[len(stk)-1]] == curr {
				stk = stk[:len(stk)-1]
			}
			width := i
			if len(stk) != 0 {
				width = i - stk[len(stk)-1] - 1
			}
			res = max(res, width*curr)
		}
		stk = append(stk, i)
	}

	for len(stk) > 0 {
		curr := heights[stk[len(stk)-1]]
		stk = stk[:len(stk)-1]
		for len(stk) > 1 && heights[stk[len(stk)-1]] == curr {
			stk = stk[:len(stk)-1]
		}
		width := n
		if len(stk) != 0 {
			width = n - stk[len(stk)-1] - 1
		}
		res = max(res, width*curr)
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

func main() {
	largestRectangleArea([]int{2, 1, 0, 2})
}
