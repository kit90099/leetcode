package main

type NumArray struct {
	arr *[]int
}

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums)+1)
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}

	return NumArray{
		arr: &arr,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return (*this.arr)[right+1] - (*this.arr)[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
