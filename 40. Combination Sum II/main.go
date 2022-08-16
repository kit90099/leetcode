package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := combinationSum2([]int{2, 5, 2, 1, 2}, 5)
	fmt.Println(ans)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(a, b int) bool {
		return candidates[a] < candidates[b]
	})

	res = make([][]int, 0)

	bactrack(candidates, []int{}, target, 0, 0)

	return res
}

var res [][]int

func bactrack(candidates []int, combintion []int, target int, pos int, sum int) {
	if sum > target {
		return
	}

	if sum == target {
		var temp []int = make([]int, len(combintion))
		copy(temp, combintion)
		res = append(res, temp)
		return
	}

	for i := pos; i < len(candidates); i++ {
		if i > pos && candidates[i] == candidates[i-1] {
			continue
		}

		combintion = append(combintion, candidates[i])
		bactrack(candidates, combintion, target, i+1, sum+candidates[i])
		combintion = combintion[:len(combintion)-1]
	}
}
