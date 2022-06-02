package main

type UF struct {
	count   int
	parents []int
}

func (uf *UF) union(p int, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)

	if rootP == rootQ {
		return
	}

	uf.parents[rootQ] = rootP
	uf.count--
}

func (uf *UF) connected(p int, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)

	return rootP == rootQ
}

func (uf *UF) find(x int) int {
	if uf.parents[x] != x {
		uf.parents[x] = uf.find(uf.parents[x])
	}

	return uf.parents[x]
}

func equationsPossible(equations []string) bool {
	parents := make([]int, 26)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	uf := UF{
		count:   26,
		parents: parents,
	}

	for _, equation := range equations {
		if equation[1] == '=' {
			uf.union(int(equation[0])-97, int(equation[3])-97)
		}
	}

	for _, equation := range equations {
		if equation[1] == '!' {
			if uf.connected(int(equation[0]-97), int(equation[3]-97)) {
				return false
			}
		}
	}

	return true
}
