package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	ans := [][]int{}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		nums := append(nums[:i], nums[i+1:]...)
		res := permute(nums)

		for _, r := range res {
			t := []int{}
			t = append(t, num)
			t = append(t, r...)
			ans = append(ans, t)
		}

		nums = append(nums[:i+1], nums[i:]...)
		nums[i] = num
	}

	return ans
}
