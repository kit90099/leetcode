package main

import "fmt"

func main() {
	ans := canFinish(1, [][]int{})
	fmt.Println(ans)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)

	hasCycle = false
	visited = make([]bool, numCourses)
	onPath = make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}

	return !hasCycle
}

func buildGraph(numCourse int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourse)
	for i := 0; i < numCourse; i++ {
		graph[i] = make([]int, 0)
	}

	for _, prerequisite := range prerequisites {
		from, to := prerequisite[1], prerequisite[0]
		graph[from] = append(graph[from], to)
	}

	return graph
}

var hasCycle bool
var visited []bool
var onPath []bool

func traverse(graph [][]int, s int) {
	if onPath[s] {
		hasCycle = true
	}

	if visited[s] || hasCycle {
		return
	}

	visited[s] = true
	onPath[s] = true
	for _, t := range graph[s] {
		traverse(graph, t)
	}

	onPath[s] = false
}
