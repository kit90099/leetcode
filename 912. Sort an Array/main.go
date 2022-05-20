package main

import "fmt"

func main() {
	ans := sortArray([]int{5, 2, 3, 1})
	fmt.Println(ans)
}

var temp []int

func sortArray(nums []int) []int {
	temp = make([]int, len(nums))
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, left int, right int) {
	if left == right {
		return
	}

	mid := (left + right) / 2
	sort(nums, left, mid)
	sort(nums, mid+1, right)

	merge(&nums, left, mid, right)
}

func merge(nums *[]int, left int, mid int, right int) {
	copy(temp, *nums)

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
