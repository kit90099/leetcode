package main

import "fmt"

func main() {
	ans := countRangeSum([]int{-2, 5, -1}, -2, 2)
	fmt.Println(ans)
}

func countRangeSum(nums []int, lower int, upper int) int {
	temp = make([]int, len(nums)+1)
	count := 0

	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	sort(preSum, 0, len(preSum)-1, &count, upper, lower)
	return count
}

var temp []int

func sort(nums []int, left int, right int, count *int, upper int, lower int) {
	if left == right {
		return
	}

	mid := (left + right) / 2
	sort(nums, left, mid, count, upper, lower)
	sort(nums, mid+1, right, count, upper, lower)

	merge(&nums, left, mid, right, count, upper, lower)
}

func merge(nums *[]int, left int, mid int, right int, count *int, upper int, lower int) {
	for i := left; i <= right; i++ {
		temp[i] = (*nums)[i]
	}

	start, end := mid+1, mid+1
	for i := left; i <= mid; i++ {
		for start <= right && (*nums)[start]-(*nums)[i] < lower {
			start++
		}

		for end <= right && (*nums)[end]-(*nums)[i] <= upper {
			end++
		}

		(*count) += (end - start)
	}

	lPtr, rPtr := left, mid+1
	for i := left; i <= right; i++ {
		if lPtr == mid+1 {
			(*nums)[i] = temp[rPtr]
			rPtr++
		} else if rPtr == right+1 {
			(*nums)[i] = temp[lPtr]
			lPtr++
		} else if temp[lPtr] > temp[rPtr] {
			(*nums)[i] = temp[rPtr]
			rPtr++
		} else {
			(*nums)[i] = temp[lPtr]
			lPtr++
		}
	}
}
