package main

import "math"

func main() {
	preimageSizeFZF(1000000000)
}

func preimageSizeFZF(k int) int {
	return int(right(k)-left(k)) + 1
}

func left(target int) int64 {
	var l, r int64 = 0, math.MaxInt64
	t := int64(target)

	for l < r {
		mid := l + (r-l)/2
		z := zeros(mid)

		if z < t {
			l = mid + 1
		} else if z > t {
			r = mid
		} else {
			r = mid
		}
	}
	return l
}

func right(target int) int64 {
	var l, r int64 = 0, math.MaxInt64
	t := int64(target)

	for l < r {
		mid := l + (r-l)/2
		z := zeros(mid)

		if z < t {
			l = mid + 1
		} else if z > t {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

func zeros(n int64) int64 {
	var res int64 = 0
	for d := n; d/5 > 0; d /= 5 {
		res += d / 5
	}
	return res
}

//1000000000
