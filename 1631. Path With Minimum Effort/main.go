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
	x                 int
	y                 int
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

var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func adj(matrix [][]int, x int, y int) [][]int {
	m, n := len(matrix), len(matrix[0])
	neighbour := make([][]int, 0)
	for _, d := range dirs {
		nx := x + d[0]
		ny := y + d[1]
		if nx >= m || nx < 0 || ny >= n || ny < 0 {
			continue
		}
		neighbour = append(neighbour, []int{nx, ny})
	}
	return neighbour
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	distanceTo := make([][]int, m)
	for i := 0; i < m; i++ {
		distanceTo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distanceTo[i][j] = math.MaxInt
		}
	}

	var pq PriorityQueue = []*State{}
	heap.Push(&pq, &State{
		x:                 0,
		y:                 0,
		distanceFromStart: 0,
	})

	for len(pq) != 0 {
		s := heap.Pop(&pq).(State)
		x, y, effort := s.x, s.y, s.distanceFromStart

		if x == m-1 && y == n-1 {
			return distanceTo[m-1][n-1]
		}

		if effort > distanceTo[x][y] {
			continue
		}

		adjNode := adj(heights, s.x, s.y)
		for _, n := range adjNode {
			effortToNext := max(
				distanceTo[n[0]][n[1]],
				abs(heights[n[0]][n[1]]-heights[x][y]),
			)

		}
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
