package main

import "fmt"

func main() {
	ans := reversePairs([]int{2, 4, 3, 5, 1})
	fmt.Println(ans)
}

var temp []int

func reversePairs(nums []int) int {
	temp = make([]int, len(nums))
	count := 0
	sort(nums, 0, len(nums)-1, &count)
	return count
}

func sort(nums []int, left int, right int, count *int) {
	if left == right {
		return
	}

	mid := (left + right) / 2
	sort(nums, left, mid, count)
	sort(nums, mid+1, right, count)

	merge(&nums, left, mid, right, count)
}

func merge(nums *[]int, left int, mid int, right int, count *int) {
	/**
	dont init all every time!, just init needed
	**/
	for i := left; i <= right; i++ {
		temp[i] = (*nums)[i]
	}

	for i := left; i <= mid; i++ {
		ptr := mid + 1

		for ptr <= right {
			if (*nums)[i] <= 2*(*nums)[ptr] {
				break
			} else {
				ptr++
			}
		}

		(*count) += ptr - mid - 1
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
