package main

import "fmt"

func main() {
	arr := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []int) {
	sortInRange(arr, 0, len(arr)-1)
}

func sortInRange(arr []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := arr[start]

	l, r := start+1, end
	for l <= r {
		count := 0
		if arr[l] <= pivot {
			l++
		} else {
			count++
		}
		if arr[r] >= pivot {
			r--
		} else {
			count++
		}

		if count == 2 {
			swap(arr, l, r)
		}

	}

	swap(arr, start, l-1)

	sortInRange(arr, start, l-2)
	sortInRange(arr, l, end)
}

func swap(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
