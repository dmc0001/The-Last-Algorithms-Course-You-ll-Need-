package connectedcomponents

func connectedComponentsCount(graph map[int][]int) int {
	visited := make(map[int]bool)
	count := 0
	for node := range graph {
		if explore(graph, visited, node) {
			count++
		}
	}
	return count
}

func explore(graph map[int][]int, visited map[int]bool, curr int) bool {
	if visited[curr] {
		return false
	}
	visited[curr] = true
	neighbors := graph[curr]
	for neighbor := range neighbors {
		explore(graph, visited, neighbor)
	}

	return true
}
