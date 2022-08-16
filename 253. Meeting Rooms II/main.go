package main

import "sort"

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	begin := []int{}
	end := []int{}
	for _, i := range intervals {
		begin = append(begin, i[0])
		end = append(end, i[1])
	}

	less := func(a, b int) bool {
		return a < b
	}

	sort.Slice(begin, func(a, b int) bool {
		return begin[a] < begin[b]
	})
	sort.Slice(end, less)

	count := 0
	res, i, j := 0, 0, 0
	for i < n && j < n {
		if begin[i] < end[j] {
			count++
			i++
		} else {
			count--
			j++
		}
		res = max(res, count)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func less(a, b int) bool {
	return a < b
}
