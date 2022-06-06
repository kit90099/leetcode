package main

import (
	"container/heap"
	"math"
)

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

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	s := x.(*State)
	s.index = n
	*pq = append(*pq, s)
}

func (pq *PriorityQueue) Pop() any {
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
		curState := heap.Pop(&pq).(State)
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
