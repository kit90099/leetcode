package main

func generateParenthesis(n int) []string {
	used = make([]int, 2)
	used[0], used[1] = 0, 0
	res = make([]string, 0)
	backtrack("", n)
	return res
}

var used []int
var res []string

func backtrack(curr string, n int) {
	if used[1] == n {
		res = append(res, curr)
	}

	// open
	if used[0] < n {
		used[0]++
		backtrack(curr+"(", n)
		used[0]--
	}

	// close
	if used[0] > used[1] {
		used[1]++
		backtrack(curr+")", n)
		used[1]--
	}
}
