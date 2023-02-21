package main

import "fmt"

func removeDuplicateLetters(s string) string {
	stk := []rune{}
	inStk := make(map[rune]bool)
	count := make(map[rune]int)

	for _, c := range s {
		count[c]++
	}

	for _, c := range s {
		count[c]--

		if inStk[c] {
			continue
		}

		for len(stk) > 0 && stk[len(stk)-1] > c {
			if count[stk[len(stk)-1]] == 0 {
				break
			}
			ele := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			inStk[ele] = false
		}

		stk = append(stk, c)
		inStk[c] = true
	}

	return string(stk)
}

func main() {
	fmt.Println(removeDuplicateLetters("bcac"))
}
