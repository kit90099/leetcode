package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int, len(nums))

	for i, num := range nums {
		mapNums[num] = i
	}
	for i, num := range nums {
		pos, exists := mapNums[target-num]

		if exists && i != pos {
			return []int{i, pos}
		}
	}
	return []int{}
}
