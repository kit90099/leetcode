package main

import (
	"math"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[a][1] < points[b][1]
	})

	end := points[0][1]
	res := 1
	for _, p := range points {
		if p[0] > end {
			res++
			end = p[1]
		}
	}

	return res
}

type Interval [][]int

func (i Interval) Len() int {
	return len(i)
}

func (i Interval) Less(a, b int) bool {
	return i[a][1] < i[b][1]
}

func (i Interval) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func findMinArrowShots(points [][]int) int {
	sort.Sort(pointSlice(points))
	i := 0 // end
	j := 1 // i
	result := 0
	for i < len(points) && j < len(points) {
		if points[i][1] >= points[j][0] {
			points[i][1] = int(math.Min(float64(points[i][1]), float64(points[j][1])))
			j++
		} else {
			result++
			i = j
			j++
		}
	}
	result++
	return result
}
