package main

import (
	"math"
	"sort"
)

func videoStitching(clips [][]int, time int) int {
	if time == 0 {
		return 0
	}

	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		} else {
			return clips[i][0] < clips[j][0]
		}
	})

	i, n, res := 0, len(clips), 0
	currEnd, nextEnd := 0, 0

	for i < n && clips[i][0] <= currEnd {
		for i < n && clips[i][0] <= currEnd {
			nextEnd = int(math.Max(float64(nextEnd), float64(clips[i][1])))
			i++
		}

		res++
		currEnd = nextEnd
		if currEnd >= time {
			return res
		}
	}

	return -1
}
