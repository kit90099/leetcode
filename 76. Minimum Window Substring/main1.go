package main

import (
	"fmt"
	"os"
)

func main() {
	s, _ := os.ReadFile("s.txt")
	t, _ := os.ReadFile("t.txt")

	fmt.Println(minWindow1(string(s), string(t)))
}

func minWindow1(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	dictT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		dictT[s[i]] = dictT[s[i]] + 1
	}

	required := len(dictT)

	l, r := 0, 0

	formed := 0

	windowCounts := make(map[byte]int)

	ans := []int{-1, 0, 0}

	for r < len(s) {
		c := s[r]
		windowCounts[c] = windowCounts[c] + 1

		dictTCount, dictTOk := dictT[c]
		if dictTOk && windowCounts[c] == dictTCount {
			formed++
		}

		for l <= r && formed == required {
			c = s[l]

			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}

			windowCounts[c] = windowCounts[c] - 1
			dictTCount, dictTOk = dictT[c]
			if dictTOk && windowCounts[c] < dictTCount {
				formed--
			}

			l++
		}
	}

	if ans[0] == -1 {
		return ""
	} else {
		return s[ans[1] : ans[2]+1]
	}
}
