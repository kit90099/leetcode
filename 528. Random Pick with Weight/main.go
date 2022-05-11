package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	obj := Constructor([]int{3, 14, 1, 7})
	for i := 0; i < 20; i++ {
		ans := obj.PickIndex()
		fmt.Println(ans)
	}
}

type Solution struct {
	preSum []int
	arr    []int
}

func Constructor(w []int) Solution {
	preSum := make([]int, len(w)+1)
	for i, num := range w {
		preSum[i+1] = preSum[i] + num
	}

	s := Solution{
		arr:    w,
		preSum: preSum,
	}

	return s
}

func (this *Solution) PickIndex() int {
	rand := math.Floor(rand.Float64()*float64(this.preSum[len(this.preSum)-1])) + 1

	l, r := 0, len(this.preSum)-1
	for l < r {
		idx := (l + r) / 2

		if this.preSum[idx] > int(rand) {
			r = idx
		} else if this.preSum[idx] < int(rand) {
			l = idx + 1
		} else {
			r = idx
		}
	}
	return l - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
