package main

/**
first version
**/
/* func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	minSum := 0
	start := 0
	for i := 0; i < n; i++ {
		sum = sum + (gas[i] - cost[i])
		if sum < minSum {
			minSum = sum
			start = i + 1
		}
	}

	if sum < 0 {
		return -1
	}

	if start == n {
		return 0
	} else {
		return start
	}
} */

/**
ver2
**/
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	for i := 0; i < n; i++ {
		sum += (gas[i] - cost[i])
	}
	if sum < 0 {
		return -1
	}

	sum = 0
	start := 0
	for i := 0; i < n; i++ {
		sum = gas[i] - cost[i]

		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}

	if start == n {
		return 0
	} else {
		return start
	}
}
