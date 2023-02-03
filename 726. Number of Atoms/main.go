package main

import (
	"sort"
	"strconv"
	"unicode"
)

func countOfAtoms(formula string) string {
	i, n := 0, len(formula)
	stk := []map[string]int{{}}
	parseAtom := func() string {
		start := i
		i++
		if i < n && unicode.IsLower((rune(formula[i]))) {
			i++
		}
		return formula[start:i]
	}
	parseNum := func() (num int) {
		if i == n || !unicode.IsDigit(rune(formula[i])) {
			return 1
		}
		for ; i < n && unicode.IsDigit(rune(formula[i])); i++ {
			num = num*10 + int(formula[i]-'0')
		}
		return
	}
	for i < n {
		if c := formula[i]; c == '(' {
			i++
			stk = append(stk, map[string]int{})
		} else if c == ')' {
			i++
			num := parseNum()
			m := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			last := stk[len(stk)-1]
			for a, v := range m {
				last[a] += v * num
			}
		} else {
			atom := parseAtom()
			num := parseNum()
			stk[len(stk)-1][atom] += num
		}
	}

	finalMap := stk[0]
	type Pair struct {
		k string
		v int
	}
	sList := []Pair{}
	for k, v := range finalMap {
		sList = append(sList, Pair{k, v})
	}
	sort.Slice(sList, func(i, j int) bool {
		return sList[i].k < sList[j].k
	})
	ans := ""
	for _, p := range sList {
		ans += p.k
		if p.v > 1 {
			ans += strconv.Itoa(p.v)
		}
	}

	return ans
}

func main() {
	countOfAtoms("K4(ON(SO3)2)2")
}
