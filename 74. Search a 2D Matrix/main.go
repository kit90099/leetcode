package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	flatten := make([]int, m*n)
	for i := 0; i < m; i++ {
		copy(flatten[i*n:], matrix[i])
	}

	l, r := 0, len(flatten)
	for l <= r {
		mid := l + (r-l)/2
		if flatten[mid] == target {
			return true
		} else if flatten[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
