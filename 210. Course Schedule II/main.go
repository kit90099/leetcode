package main

import "fmt"

func main() {
	ans := findOrder(1, [][]int{})
	fmt.Println(ans)
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(numCourses, prerequisites)

	hasCycle = false
	visited = make([]bool, numCourses)
	onPath = make([]bool, numCourses)
	result = make([]int, 0)

	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}

	if hasCycle {
		return make([]int, 0)
	}

	l, r := 0, len(result)-1
	for l < r {
		swap(result, l, r)
		l++
		r--
	}

	return result
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
var result []int

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
	result = append(result, s)
}

func swap(arr []int, posA int, posB int) {
	temp := arr[posA]
	arr[posA] = arr[posB]
	arr[posB] = temp
}
