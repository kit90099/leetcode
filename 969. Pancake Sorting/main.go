package main

import "fmt"

func main() {
	ans := pancakeSort([]int{3, 2, 4, 1})
	fmt.Println(ans)
}

func pancakeSort(arr []int) []int {
	var res []int = make([]int, 0)
	for i := len(arr); i > 0; i-- {
		idx := 0
		for _idx, num := range arr {
			if num == i {
				idx = _idx
			}
		}
		res = append(res, idx+1)
		res = append(res, i)

		flip(arr, idx)
		flip(arr, i-1)

	}
	return res
}

func flip(arr []int, end int) {
	i, j := 0, end
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
