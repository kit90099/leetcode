package main

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	slow := -1
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			slow++
			nums[slow] = nums[fast]
		}
	}

	for i := slow + 1; i < len(nums); i++ {
		nums[i] = 0
	}
}
