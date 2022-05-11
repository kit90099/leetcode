package main

import (
	"math"
	"sort"
)

func main() {
	minEatingSpeed([]int{3}, 8)
}

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)

	l, r := 0, piles[len(piles)-1]
	for l < r {
		mid := l + (r-l)/2

		if f(piles, mid) <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func f(piles []int, x int) int {
	if x == 0 {
		return math.MaxInt
	}

	t := 0
	for i := 0; i < len(piles); i++ {
		pile := piles[i]
		t += int(math.Ceil(float64(pile) / float64(x)))
	}

	return t
}
