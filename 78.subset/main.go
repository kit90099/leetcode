package main

import "fmt"

func main() {
	res := subsets([]int{1, 2, 3, 4, 5})
	fmt.Print(res)
}

var sets = [][]int{}

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
