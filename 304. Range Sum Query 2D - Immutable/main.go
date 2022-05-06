package main

import "fmt"

func main() {
	numMartix := Constructor([][]int{[]int{3, 0, 1, 4, 2}, []int{5, 6, 3, 2, 1}, []int{1, 2, 0, 1, 5}, []int{4, 1, 0, 1, 7}, []int{1, 0, 3, 0, 5}})
	fmt.Print(numMartix.SumRegion(0, 1, 0, 3))
}

type NumMatrix struct {
	arr *[][]int
}

func Constructor(matrix [][]int) NumMatrix {
	arr := make([][]int, len(matrix)+1)
	arr[0] = make([]int, len(matrix[0])+1)

	for i := 1; i < len(matrix)+1; i++ {
		subArr := matrix[i-1]
		arr[i] = make([]int, len(subArr)+1)
		for j := 1; j < len(subArr)+1; j++ {
			arr[i][j] = matrix[i-1][j-1] + arr[i][j-1] + arr[i-1][j] - arr[i-1][j-1]
		}
	}

	return NumMatrix{
		arr: &arr,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return (*this.arr)[row2+1][col2+1] - (*this.arr)[row2+1][col1] - (*this.arr)[row1][col2+1] + (*this.arr)[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
