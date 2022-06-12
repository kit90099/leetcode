package main

import (
	"container/heap"
	"fmt"
)

func main() {
	ans := maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2)
	fmt.Println(ans)
}

type State struct {
	idx           int
	index         int
	probFromStart float64
}

type Edge struct {
	from   int
	to     int
	weight float64
}

type PriorityQueue []*State

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(a, b int) bool {
	return (*pq)[a].probFromStart > (*pq)[b].probFromStart
}

func (pq *PriorityQueue) Swap(a, b int) {
	(*pq)[a], (*pq)[b] = (*pq)[b], (*pq)[a]
	(*pq)[a].index = a
	(*pq)[b].index = b
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	s := x.(*State)
	s.index = n
	*pq = append(*pq, s)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(s *State, priority float64) {
	(*s).probFromStart = priority
	heap.Fix(pq, s.index)
}

var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func adj(graph [][]Edge, s int) []Edge {
	return graph[s]
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	probTo := make([]float64, n)
	for i := 0; i < n; i++ {
		probTo[i] = float64(-1)
	}
	probTo[start] = 1

	var pq PriorityQueue = []*State{}
	heap.Push(&pq, &State{
		idx:           start,
		probFromStart: 1,
	})

	graph := buildGraph(n, edges, succProb)

	for len(pq) != 0 {
		s := heap.Pop(&pq).(*State)
		idx, effort := s.idx, s.probFromStart

		if idx == end {
			return probTo[end]
		}

		if effort < probTo[end] {
			continue
		}

		adjNode := adj(graph, idx)
		for _, n := range adjNode {
			next := n.to
			probToNext := probTo[idx] * n.weight

			if probTo[next] < probToNext {
				probTo[next] = probToNext
				heap.Push(&pq, &State{
					idx:           next,
					probFromStart: probToNext,
				})
			}
		}
	}

	return -1
}

func buildGraph(n int, edges [][]int, succProb []float64) [][]Edge {
	graph := make([][]Edge, n)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]Edge, 0)
	}

	for i := 0; i < len(edges); i++ {
		edge, prob := edges[i], succProb[i]

		graph[edge[0]] = append(graph[edge[0]], Edge{from: edge[0], to: edge[1], weight: prob})
		graph[edge[1]] = append(graph[edge[1]], Edge{from: edge[1], to: edge[0], weight: prob})
	}

	return graph
}
