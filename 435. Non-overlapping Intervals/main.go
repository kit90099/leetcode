package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}})
	fmt.Println(ans)
}

func eraseOverlapIntervals(intervals [][]int) int {
	var t Interval = intervals
	sort.Sort(t)

	res := 1
	end := t[0][1]
	for i := 1; i < len(intervals); i++ {
		start := t[i][0]
		if start < end {
			res++
		} else {
			end = t[i][1]
		}
	}
	return res
}

type Interval [][]int

func (i Interval) Len() int {
	return len(i)
}

func (t Interval) Less(i, j int) bool {
	if t[i][1] == t[j][1] {
		return t[i][0] < t[j][0]
	}
	return t[i][1] < t[j][1]
}

func (t Interval) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
