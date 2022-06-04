package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := minimumCost(4, [][]int{{1, 2, 3}, {3, 4, 4}})
	fmt.Println(ans)
}

type Edges [][]int

func (edge *Edges) Len() int {
	return len((*edge))
}

func (edge *Edges) Swap(i, j int) {
	(*edge)[i], (*edge)[j] = (*edge)[j], (*edge)[i]
}

func (edge *Edges) Less(i, j int) bool {
	return (*edge)[i][2] < (*edge)[j][2]
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

func minimumCost(n int, connections [][]int) int {
	uf := constructUF(n + 1)

	var edges Edges = connections
	sort.Sort(&edges)

	res := 0
	for _, edge := range edges {
		if uf.connected(edge[0], edge[1]) {
			continue
		}
		uf.union(edge[0], edge[1])
		res += edge[2]
	}

	if uf.count != 2 {
		return -1
	}

	return res
}
