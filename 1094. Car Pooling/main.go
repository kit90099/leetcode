package main

import "fmt"

func main() {
	ans := carPooling([][]int{{9, 0, 1}, {3, 3, 7}}, 4)
	fmt.Println(ans)
}

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)

	last := -1

	for i := 0; i < len(trips); i++ {
		diff[trips[i][1]] += trips[i][0]
		diff[trips[i][2]] -= trips[i][0]

		if trips[i][2] > last {
			last = trips[i][2]
		}
	}

	nums := make([]int, last+2)
	for i := 1; i <= last; i++ {
		nums[i] = nums[i-1] + diff[i-1]
		if nums[i] > capacity {
			return false
		}
	}
	return true
}
