package main

import "sort"

type Doll [][]int

func (this Doll) Less(i, j int) bool {
	if this[i][0] == this[j][0] {
		return this[j][1] < this[i][1]
	} else {
		return this[i][0] < this[j][0]
	}
}

func (this Doll) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Doll) Len() int {
	return len(this)
}

func maxEnvelopes(envelopes [][]int) int {
	top := make([]int, len(envelopes))
	piles := 0
	var dolls Doll = envelopes
	sort.Sort(dolls)

	for i := 0; i < len(envelopes); i++ {
		d := dolls[i]
		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if d[1] > top[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == piles {
			piles++
		}
		top[left] = d[1]
	}
	return piles

}
