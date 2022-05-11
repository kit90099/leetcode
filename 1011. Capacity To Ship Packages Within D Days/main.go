package main

func main() {
	shipWithinDays([]int{1, 2, 3, 1, 1}, 4)
}

func shipWithinDays(weights []int, days int) int {

	sum := 0
	max := -1
	for _, weight := range weights {
		sum += weight
		if weight > max {
			max = weight
		}
	}

	l, r := max, sum
	for l < r {
		mid := (l + r) / 2

		timeNeeded := f(weights, mid)

		if timeNeeded <= days {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func f(weights []int, x int) int {
	t := 0
	weightSum := 0
	for i := 0; i < len(weights); i++ {
		if weightSum+weights[i] > x {
			t++
			weightSum = weights[i]
		} else {
			weightSum += weights[i]
		}
	}

	return t + 1
}
