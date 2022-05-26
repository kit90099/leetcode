package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	return QSearch(nums, k)
}

func QSearch(arr []int, k int) int {
	shuffle(&arr)

	return sort(&arr, 0, len(arr)-1, len(arr)-k)
}

func sort(nums *[]int, start int, end int, pos int) int {
	pivotPos := partition(nums, start, end)

	if pos == pivotPos {
		return (*nums)[pivotPos]
	} else if pos > pivotPos {
		return sort(nums, pivotPos+1, end, pos)
	} else {
		return sort(nums, start, pivotPos-1, pos)
	}
}

func partition(arr *[]int, start int, end int) int {
	pivot := (*arr)[start]

	i, j := start+1, end
	for i <= j {
		for i < end && (*arr)[i] <= pivot {
			i++
		}

		for j > start && (*arr)[j] > pivot {
			j--
		}

		if i >= j {
			break
		}
		swap(arr, i, j)
	}

	swap(arr, start, j)

	return j
}

func shuffle(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		swap(arr, i, r)
	}
}

func swap(arr *[]int, a int, b int) {
	temp := (*arr)[a]
	(*arr)[a] = (*arr)[b]
	(*arr)[b] = temp
}
