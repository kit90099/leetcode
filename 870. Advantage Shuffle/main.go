package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := advantageCount([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2})
	fmt.Println(ans)
}

func advantageCount(nums1 []int, nums2 []int) []int {
	n2Map := make(map[int]int)
	for i, num := range nums2 {
		n2Map[num] = i
	}

	sort.Ints(nums1)
	n2Sorted := make([]int, len(nums2))
	copy(n2Sorted, nums2)
	sort.Ints(n2Sorted)

	ans := make([]int, len(nums1))
	l, r := 0, len(nums1)-1
	n2ptr := len(n2Sorted) - 1
	for l <= r {
		pos := n2Map[n2Sorted[n2ptr]]
		if nums1[r] > n2Sorted[n2ptr] {
			ans[pos] = nums1[r]
			r--
		} else {
			ans[pos] = nums1[l]
			l++
		}
		n2ptr--
	}
	return ans
}
