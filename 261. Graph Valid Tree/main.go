package main

import "fmt"

func main() {
	ans := validTree(5, [][]int{{1, 2}, {2, 4}, {2, 5}, {1, 3}})
	fmt.Println(ans)
}

type UF struct {
	count   int
	parents []int
}

func constructUF(n int) *UF {
	parents := make([]int, n)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	return &UF{
		count:   n,
		parents: parents,
	}
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

func validTree(n int, edges [][]int) bool {
	uf := constructUF(n + 1)

	for _, edge := range edges {
		if uf.connected(edge[0], edge[1]) {
			return false
		}

		uf.union(edge[0], edge[1])
	}

	return uf.count == 2
}
