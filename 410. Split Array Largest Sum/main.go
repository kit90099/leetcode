package main

import "fmt"

func main() {
	ans := splitArray([]int{1, 4, 4}, 3)
	fmt.Println(ans)
}

func splitArray(nums []int, m int) int {
	sum := 0
	max := -1
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}

	l, r := max, sum

	for l < r {
		mid := (l + r) / 2

		if minimumSubarraysRequired(nums, mid) <= m {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// minimumSubarraysRequired
func minimumSubarraysRequired(arr []int, x int) int {
	currSum := 0
	splitRequired := 0

	for _, num := range arr {
		if currSum+num <= x {
			currSum += num
		} else {
			currSum = num
			splitRequired++
		}
	}

	return splitRequired + 1
}
