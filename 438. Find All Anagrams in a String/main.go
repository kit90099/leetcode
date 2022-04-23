package main

import (
	"fmt"
)

func main() {
	fmt.Print(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {
	l, r := 0, 0

	needs := make(map[byte]int, 26)
	windows := make(map[byte]int, 26)

	for i := 0; i < len(p); i++ {
		needs[p[i]] = needs[p[i]] + 1
	}

	valid := 0
	ans := []int{}

	for r < len(s) {
		char := s[r]
		need, exist := needs[char]

		if !exist {
			r++
			continue
		}

		windows[char] = windows[char] + 1
		if windows[char] == need {
			valid++
		}

		for valid == len(needs) && l < r-len(p)+1 {
			lChar := s[l]
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
			ans = append(ans, l)
		}
		r++
	}

	return ans
}
