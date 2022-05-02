package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
}

func fourSum(nums []int, target int) [][]int {
	answer := [][]int{}
	sort.Ints(nums)

	i := 0
	for i < len(nums) {
		resultThreeSum := threeSum(nums, i+1, target-nums[i])
		if len(resultThreeSum) != 0 {
			for j := 0; j < len(resultThreeSum); j++ {
				resultThreeSum[j] = append(resultThreeSum[j], nums[i])
			}
			answer = append(answer, resultThreeSum...)
		}

		i++
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
	}

	return answer
}

func threeSum(nums []int, start int, target int) [][]int {
	answer := [][]int{}
	sort.Ints(nums)
	i := start
	for i < len(nums) {
		resultTwoSum := twoSum(nums, i+1, target-nums[i])
		if len(resultTwoSum) != 0 {
			for j := 0; j < len(resultTwoSum); j++ {
				resultTwoSum[j] = append(resultTwoSum[j], nums[i])
			}
			answer = append(answer, resultTwoSum...)
		}

		i++
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
	}

	return answer
}

func twoSum(nums []int, start int, target int) [][]int {
	n := len(nums)
	l, r := start, n-1
	ans := [][]int{}

	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			ans = append(ans, []int{nums[l], nums[r]})
		}

		if sum < target {
			l++
			for l < r && nums[l] == nums[l-1] {
				l++
			}
		} else {
			r--
			for l < r && nums[r] == nums[r+1] {
				r--
			}
		}
	}

	return ans

}
