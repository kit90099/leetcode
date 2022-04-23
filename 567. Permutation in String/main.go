package main

import (
	"fmt"
)

func main() {
	fmt.Print(checkInclusion("abcdxabcde", "abcdeabcdx"))
}

func checkInclusion(s1 string, s2 string) bool {
	l, r := 0, 0

	needs := make(map[byte]int, 26)
	windows := make(map[byte]int, 26)

	for i := 0; i < len(s1); i++ {
		needs[s1[i]] = needs[s1[i]] + 1
	}

	valid := 0
	for r < len(s2) {
		char := s2[r]
		need, exist := needs[char]

		if !exist {
			r++
			continue
		}

		windows[char] = windows[char] + 1
		if windows[char] == need {
			valid++
		}

		for valid == len(needs) && l < r-len(s1)+1 {
			lChar := s2[l]
			need, exist = needs[lChar]

			if !exist {
				l++
				continue
			}

			windows[lChar] = windows[lChar] - 1
			if windows[lChar] < need {
				valid--
			}
			l++
		}
		if valid == len(needs) {
			return true
		}
		r++
	}

	return false
}
