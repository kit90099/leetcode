package main

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})

	left := intervals[0][0]
	right := intervals[0][1]
	res := 0
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[0] >= left && interval[1] <= right {
			res++
		} else if interval[0] < right {
			right = interval[1]
		} else if interval[0] > right {
			left = interval[0]
			right = interval[1]
		}
	}

	return len(intervals) - res
}
