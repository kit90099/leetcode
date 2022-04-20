package main

import "fmt"

func main() {
	ans := longestPalindrome("babad")
	fmt.Print(ans)
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	ans := ""

	for i := 0; i < len(s); i++ {
		odd := palindrome(s, i, i)
		even := palindrome(s, i, i+1)

		if len(odd) > len(ans) {
			ans = odd
		}

		if len(even) > len(ans) {
			ans = even
		}
	}

	return ans
}

func palindrome(s string, l int, r int) string {
	if r >= len(s) || l != r && s[l] != s[r] {
		return ""
	}

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return s[l+1 : r]
}
