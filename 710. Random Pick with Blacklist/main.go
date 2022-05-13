package main

import (
	"math"
	"math/rand"
)

type Solution struct {
	actualSize int
	mapping    map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	res := Solution{
		actualSize: n - len(blacklist),
		mapping:    make(map[int]int),
	}

	for _, num := range blacklist {
		res.mapping[num] = 99999
	}

	last := n - 1
	for _, num := range blacklist {
		for res.mapping[last] == 99999 {
			last--
		}

		if num >= res.actualSize {
			continue
		}

		res.mapping[num] = last
		last--
	}

	return res
}

func (this *Solution) Pick() int {
	rand := math.Floor(rand.Float64() * float64(this.actualSize))

	mapping, exist := this.mapping[int(rand)]
	if exist {
		return mapping
	} else {
		return int(rand)
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
