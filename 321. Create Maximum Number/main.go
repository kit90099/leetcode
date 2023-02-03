package main

import (
	"fmt"
)

func largestKDigit(num []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if k == len(num) {
		return num
	}
	stk := []int{}
	n := len(num) - k
	for i := range num {
		c := num[i]
		for n > 0 && len(stk) > 0 && stk[len(stk)-1] < c {
			stk = stk[:len(stk)-1]
			n--
		}
		stk = append(stk, c)
	}
	stk = stk[:len(stk)-n]
	return stk
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, k)
	for n1 := 0; n1 <= k; n1++ {
		n2 := k - n1
		if m < n1 || n < n2 {
			continue
		}

		sub1 := largestKDigit(nums1, n1)
		sub2 := largestKDigit(nums2, n2)

		merged := merge(sub1, sub2, k)
		if compare(res, merged, k) {
			res = merged
		}
	}
	return res
}

func merge(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)
	if m == 0 {
		return nums2
	}
	if n == 0 {
		return nums1
	}
	res := make([]int, k)
	for i := range res {
		if compare(nums1, nums2, k) {
			res[i] = nums2[0]
			nums2 = nums2[1:]
		} else {
			res[i] = nums1[0]
			nums1 = nums1[1:]
		}
	}
	return res
}

func compare(nums1 []int, nums2 []int, k int) bool {
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] < nums2[i]
		}
	}
	return len(nums1) < len(nums2)
}

/* func removeKdigits(num string, k int) string {
	stk := []byte{}
	for i := range num {
		c := num[i]
		for k > 0 && len(stk) > 0 && stk[len(stk)-1] > c {
			stk = stk[:len(stk)-1]
			k--
		}
		stk = append(stk, c)
	}
	stk = stk[:len(stk)-k]
	ans := strings.TrimLeft(string(stk), "0")
	if ans == "" {
		return "0"
	}
	return ans
} */

func main() {
	fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
}
