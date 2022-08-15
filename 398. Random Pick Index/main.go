package main

import (
	"math/rand"
)

func main() {
	s := Constructor([]int{1, 2, 3, 3, 3})
	s.Pick(3)
}

type Solution struct {
	valToIdxMap map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		arr, exist := m[nums[i]]
		if !exist {
			arr = make([]int, 0)
		}
		arr = append(arr, i)
		m[nums[i]] = arr
	}

	return Solution{
		valToIdxMap: m,
	}
}

func (this *Solution) Pick(target int) int {
	idxs := this.valToIdxMap[target]

	return idxs[rand.Intn(len(idxs))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
