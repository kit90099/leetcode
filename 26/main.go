package main

func removeDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
