package main

import (
	"fmt"
	"strconv"
)

func main() {
	ans := isMatch("aab", "c*a*b")
	fmt.Println(ans)
}

func isMatch(s string, p string) bool {
	memo = make(map[string]bool, len(s)*len(p))
	return dp(s, p, 0, 0)
}

var memo map[string]bool

func dp(s, p string, i, j int) bool {
	if j == len(p) {
		return i == len(s)
	}

	if i == len(s) {
		if (len(p)-j)%2 == 1 {
			return false
		}

		for j++; j < len(p); j += 2 {
			if p[j] != '*' {
				return false
			}
		}

		return true
	}

	str := strconv.Itoa(i) + "|" + strconv.Itoa(j)
	val, exist := memo[str]
	if exist {
		return val
	}

	res := false
	if s[i] == p[j] || p[j] == '.' {
		if j < len(p)-1 && p[j+1] == '*' {
			res = dp(s, p, i+1, j) || dp(s, p, i, j+2)
		} else {
			res = dp(s, p, i+1, j+1)
		}
	} else {
		if j < len(p)-1 && p[j+1] == '*' {
			res = dp(s, p, i, j+2)
		} else {
			res = false
		}
	}
	memo[str] = res
	return res
}
