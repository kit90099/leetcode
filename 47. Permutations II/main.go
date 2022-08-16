package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := permuteUnique([]int{1, 1, 2})
	fmt.Println(ans)
}

func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res = make([][]int, 0)
	backtrack(nums, []int{}, len(nums))

	return res
}

var res [][]int

func backtrack(nums []int, perm []int, targetSize int) {
	if len(perm) == targetSize {
		temp := make([]int, len(perm))
		copy(temp, perm)
		res = append(res, perm)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		tPerm := make([]int, len(perm))
		copy(tPerm, perm)
		tPerm = append(tPerm, nums[i])
		tNum := make([]int, len(nums))
		copy(tNum, nums)
		tNum = append(tNum[:i], tNum[i+1:]...)
		backtrack(tNum, tPerm, targetSize)
	}
}
