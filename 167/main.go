package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left != right {

		sum := numbers[right] + numbers[left]

		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum < target {
			left++
			continue
		}

		if sum > target {
			right--
			continue
		}
	}

	return []int{}
}
