package main

import "fmt"

func main() {
	ans := isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}})
	fmt.Println(ans)
}

var color []bool
var visited []bool

var ok bool

func isBipartite(graph [][]int) bool {
	color = make([]bool, len(graph))
	visited = make([]bool, len(graph))
	ok = true

	for i := 0; i < len(graph); i++ {
		if visited[i] == false {
			traverse(graph, i)
		}
	}

	return ok
}

func traverse(graph [][]int, v int) {
	if !ok {
		return
	}

	visited[v] = true
	for _, w := range graph[v] {
		if visited[w] {
			if color[v] == color[w] {
				ok = false
			}
		} else {
			color[w] = !color[v]
			traverse(graph, w)
		}
	}
}
