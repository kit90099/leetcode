package main

func main() {
	findErrorNums([]int{1, 2, 2, 4})
}

/* func findErrorNums(nums []int) []int {
	arr := make([]int, len(nums)+1)

	ans1 := 0
	for i := 0; i < len(nums); i++ {
		if arr[nums[i]] == 0 {
			arr[nums[i]] = -nums[i]
		} else {
			ans1 = nums[i]
		}
	}

	ans2 := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == 0 {
			ans2 = i
		}
	}

	return []int{ans1, ans2}
}
*/

//ver2
func findErrorNums(nums []int) []int {
	arr := make([]int, len(nums)+1)
	allNums := 0
	for i := 0; i < len(nums); i++ {
		arr[nums[i]] = nums[i]
		allNums ^= nums[i]
	}

	ans2 := 0

	allIndex := 0
	for i := 1; i < len(arr); i++ {
		ans2 ^= arr[i] ^ i
		allIndex ^= i
	}
	return []int{allIndex ^ allNums ^ ans2, ans2}
}
