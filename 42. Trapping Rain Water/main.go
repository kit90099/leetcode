package main

func trap(height []int) int {
	l, r := 0, len(height)-1
	lmax, rmax := 0, 0

	res := 0
	for l < r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])

		if lmax < rmax {
			res += lmax - height[l]
			l++
		} else {
			res += rmax - height[r]
			r--
		}
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
