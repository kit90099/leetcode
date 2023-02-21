package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stk := []Detail{}
	ans := make([]int, n)
	i := n - 1
	for i >= 0 {
		for len(stk) > 0 && stk[len(stk)-1].temp <= temperatures[i] {
			stk = stk[:len(stk)-1]
		}

		if len(stk) != 0 {
			ans[i] = stk[len(stk)-1].day - i
		}

		stk = append(stk, Detail{i, temperatures[i]})
		i--
	}

	return ans
}

type Detail struct {
	day  int
	temp int
}
