package main

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}

	memo = make(map[int]bool)

	target := sum / k

	return backtrack(nums, k, 0, 0, 0, target)
}

var memo map[int]bool

func backtrack(nums []int, k int, bucket int, start int, used int, target int) bool {
	if k == 0 {
		return true
	}

	if bucket == target {
		res := backtrack(nums, k-1, 0, 0, used, target)
		memo[used] = res
		return res
	}

	val, exist := memo[used]
	if exist {
		return val
	}

	for i := start; i < len(nums); i++ {
		// move right by i position and doing and for 00000000001
		if (used>>i)&1 == 1 {
			continue
		}

		if bucket+nums[i] > target {
			continue
		}

		used |= 1 << i
		if backtrack(nums, k, bucket+nums[i], i+1, used, target) {
			return true
		}
		used ^= 1 << i
	}

	return false
}
