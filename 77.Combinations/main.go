package main

var ans [][]int

func combine(n int, k int) [][]int {
	ans = [][]int{}

	arr := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	traceBack(arr, []int{}, k)

	return ans
}

func traceBack(arr []int, track []int, k int) {
	if len(track) == k {
		ans = append(ans, track)
		return
	}

	arrBack := arr
	trackBack := track
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		track := append(track, num)
		arr := arr[i+1:]

		traceBack(arr, track, k)

		arr = arrBack
		track = trackBack
	}
}
