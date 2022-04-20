package main

import "fmt"

var sets = [][]int{}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ans := subsets(arr)
	fmt.Print(ans)
}

func subsets(nums []int) [][]int {
	sets = [][]int{}
	backTrack(nums, []int{})
	return sets
}

func backTrack(nums []int, track []int) {
	temp := make([]int, len(track))
	copy(temp, track)
	sets = append(sets, temp)

	numsBack := nums
	trackBack := track

	for i := 0; i < len(nums); i++ {
		track = append(track, nums[i])

		nums = nums[i+1:]
		backTrack(nums, track)

		// remove num from track
		track = trackBack
		// add num to nums
		nums = numsBack
	}

	// remove num from track
	track = trackBack
	// add num to nums
	nums = numsBack
}
