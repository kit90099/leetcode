package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	ans := networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2)
	fmt.Println(ans)
}

type State struct {
	id                int
	index             int
	distanceFromStart int
}

type PriorityQueue []*State

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(a, b int) bool {
	return (*pq)[a].distanceFromStart < (*pq)[b].distanceFromStart
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

func (pq *PriorityQueue) Update(s *State, priority int) {
	(*s).distanceFromStart = priority
	heap.Fix(pq, s.index)
}

func dijkstra(start int, graph [][][]int) []int {
	distanceTo := make([]int, len(graph))
	for i := 0; i < len(distanceTo); i++ {
		distanceTo[i] = math.MaxInt
	}
	distanceTo[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &State{
		id:                start,
		distanceFromStart: 0,
	})

	for len(pq) != 0 {
		curState := heap.Pop(&pq).(*State)
		id := curState.id

		if distanceTo[id] < curState.distanceFromStart {
			continue
		}

		for _, t := range graph[id] {
			distanceToNext := distanceTo[id] + t[1]
			if distanceTo[t[0]] > distanceToNext {
				distanceTo[t[0]] = distanceToNext
				heap.Push(&pq, &State{
					id:                t[0],
					distanceFromStart: distanceToNext,
				})
			}
		}
	}

	return distanceTo
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := buildGraph(n, times)

	distanceTo := dijkstra(k, graph)
	max := -1
	for i := 1; i < len(distanceTo); i++ {
		if distanceTo[i] == math.MaxInt {
			return -1
		}
		if i != k && distanceTo[i] > max {
			max = distanceTo[i]
		}
	}
	return max
}

func buildGraph(n int, times [][]int) [][][]int {
	graph := make([][][]int, n+1)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([][]int, 0)
	}

	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], []int{t[1], t[2]})
	}

	return graph
}
