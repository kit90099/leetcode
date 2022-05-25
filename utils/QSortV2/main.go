package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ans := []int{5, 2, 4, 1}
	QSort(ans)
	fmt.Println(ans)
}

func QSort(arr []int) {
	shuffle(&arr)

	sort(&arr, 0, len(arr)-1)
}

func sort(nums *[]int, start int, end int) {
	if start >= end {
		return
	}

	pivotPos := partition(nums, start, end)

	sort(nums, start, pivotPos-1)
	sort(nums, pivotPos+1, end)
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
