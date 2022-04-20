package main

var sets = [][]int{}

func combine(n int, k int) [][]int {
	sets = [][]int{}
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	backTrack(nums, []int{}, k)
	return sets
}

func backTrack(nums []int, track []int, k int) {
	if len(track) == k {
		temp := make([]int, len(track))
		copy(temp, track)
		sets = append(sets, temp)
		return
	}

	numsBack := nums
	trackBack := track

	for i := 0; i < len(nums); i++ {
		track = append(track, nums[i])

		nums = nums[i+1:]
		backTrack(nums, track, k)

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
