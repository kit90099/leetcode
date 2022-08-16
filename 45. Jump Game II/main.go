package main

import "fmt"

func main() {
	ans := jump([]int{2, 3, 1, 1, 4})
	fmt.Println(ans)
}

func jump(nums []int) int {
	n := len(nums)
	farest, end := 0, 0
	res := 0

	for i := 0; i < n-1; i++ {
		farest = max(farest, i+nums[i])
		if end == i {
			res++
			end = farest
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
