package main

var paths [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	paths = make([][]int, 0)
	path := make([]int, 0)
	traverse(graph, 0, &path)
	return paths
}

func traverse(graph [][]int, s int, path *[]int) {
	*path = append(*path, s)

	n := len(graph)
	if s == n-1 {
		temp := make([]int, len(*path))
		copy(temp, *path)
		paths = append(paths, temp)
	}

	for _, v := range graph[s] {
		traverse(graph, v, path)
	}

	(*path) = (*path)[:len(*path)-1]
}
