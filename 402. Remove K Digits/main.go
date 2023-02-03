package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stk := []byte{}
	for i := range num {
		c := num[i]
		for k > 0 && len(stk) > 0 && stk[len(stk)-1] > c {
			stk = stk[:len(stk)-1]
			k--
		}
		stk = append(stk, c)
	}
	stk = stk[:len(stk)-k]
	ans := strings.TrimLeft(string(stk), "0")
	if ans == "" {
		return "0"
	}
	return ans
}

func main() {
	fmt.Println(removeKdigits("987654321", 3))
}
