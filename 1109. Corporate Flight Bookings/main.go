package main

import "fmt"

func main() {
	ans := corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5)
	fmt.Println(ans)
}

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)

	for i := 0; i < len(bookings); i++ {
		diff[bookings[i][0]-1] += bookings[i][2]
		if bookings[i][1] < n {
			diff[bookings[i][1]] -= bookings[i][2]
		}
	}

	num := make([]int, n)
	num[0] = diff[0]
	for i := 1; i < n; i++ {
		num[i] = diff[i] + num[i-1]
	}

	return num
}
