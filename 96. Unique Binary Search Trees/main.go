package main

func numTrees(n int) int {
	memo = make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}
	return count(1, n)
}

var memo [][]int

func count(lo int, hi int) int {
	if lo > hi {
		return 1
	}

	if memo[lo-1][hi-1] != 0 {
		return memo[lo-1][hi-1]
	}

	res := 0
	for i := lo; i <= hi; i++ {
		left := count(lo, i-1)
		right := count(i+1, hi)

		res += left * right
	}

	memo[lo-1][hi-1] = res
	return res
}
