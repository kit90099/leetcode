package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	l, r := 0, 0
	max := -1

	windows := make(map[byte]int)

	for r < len(s) {
		char := s[r]

		windows[char] = windows[char] + 1
		if windows[char] <= 1 {
			if r-l+1 > max {
				max = r - l + 1
			}
			r++
			continue
		}

		for l <= r {
			lChar := s[l]
			l++

			windows[lChar] = windows[lChar] - 1
			if windows[lChar] == 1 {
				break
			}
		}
		r++
	}

	return max
}
