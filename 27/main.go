package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	ans := removeElement(nums, 3)
	fmt.Print(ans)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := -1

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
