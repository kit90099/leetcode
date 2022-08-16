package main

import "strconv"

func diffWaysToCompute(expression string) []int {
	res := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		c := expression[i]

		if c == '+' || c == '-' || c == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, l := range left {
				for _, r := range right {
					if c == '+' {
						res = append(res, l+r)
					} else if c == '-' {
						res = append(res, l-r)
					} else if c == '*' {
						res = append(res, l*r)
					}
				}
			}
		}
	}

	if len(res) == 0 {
		val, _ := strconv.Atoi(expression)
		res = append(res, val)
	}

	return res
}
