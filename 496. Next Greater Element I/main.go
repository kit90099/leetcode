package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	nextMap := map[int]int{}
	nextMap[nums2[n-1]] = -1
	stk := []int{nums2[n-1]}
	for i := n - 2; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] <= nums2[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			nextMap[nums2[i]] = -1
		} else {
			nextMap[nums2[i]] = stk[len(stk)-1]
		}
		stk = append(stk, nums2[i])
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = nextMap[nums1[i]]
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
