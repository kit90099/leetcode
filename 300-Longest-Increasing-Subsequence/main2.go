package main

func lengthOfLIS(nums []int) int {
	top := make([]int, len(nums))
	piles := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] >= num {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == piles {
			piles++
		}
		top[left] = num
	}
	return piles
}
