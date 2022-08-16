package main

/* func canJump(nums []int) bool {

	n := len(nums)
	i := 0
	nextEnd := 0
	for i < n {
		nextStart := 0
		for j := i + 1; j < i+nums[i]; i++ {
			currMax := max(nextEnd, j+nums[j])
			if currMax >= nextEnd {
				nextStart = j
				nextEnd = currMax
			}
		}

		i = nextStart

		if i >
	}
} */

func canJump(nums []int) bool {

	n := len(nums)
	nextEnd := 0
	for i := 0; i < n-1; i++ {
		nextEnd = max(nextEnd, i+nums[i])

		if nextEnd <= i {
			return false
		}
	}
	return nextEnd >= n-1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
