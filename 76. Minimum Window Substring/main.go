package main

import (
	"fmt"
	"os"
)

func main() {
	s, _ := os.ReadFile("s.txt")
	t, _ := os.ReadFile("t.txt")
	/* s := "a"
	t := "aa" */

	fmt.Println(minWindow(string(s), string(t)))
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	min := ^uint(0)
	ans := []int{0, -1}

	left := 0
	right := 0

	needs := make(map[byte]int, 26)
	windows := make(map[byte]int, 26)
	for i := 0; i < len(t); i++ {
		str := t[i]
		needs[str] = needs[str] + 1
	}

	valid := 0
	for right < len(s) {
		char := s[right]
		need, exist := needs[char]
		if !exist {
			right++
			continue
		}
		//update window
		windows[char] = windows[char] + 1

		// judge whether move left ptr
		if windows[char] == need {
			valid++
		}

		//start moving left
		for valid >= len(needs) && left <= right {
			toBeRemoved := s[left]
			needsToBeRemoved, existToBeRemoved := needs[toBeRemoved]
			if existToBeRemoved && windows[toBeRemoved]-1 < needsToBeRemoved {
				if uint((right - left + 1)) < min {
					min = uint(right - left + 1)
					ans[0] = left
					ans[1] = right
				}
				left++
				windows[toBeRemoved] = windows[toBeRemoved] - 1
				valid--
				break
			}
			left++
			if existToBeRemoved {
				windows[toBeRemoved] = windows[toBeRemoved] - 1
			}
		}

		right++
	}

	return s[ans[0] : ans[1]+1]
}
