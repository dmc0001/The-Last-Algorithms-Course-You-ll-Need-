package largestcomponents

func largestComponent(graph map[int][]int) int {
	visited := make(map[int]bool)

	largest := 0
	for node := range graph {
		size := exploreSize(graph, visited, node)
		if size > largest {
			largest = size
		}
	}
	return largest
}

func exploreSize(graph map[int][]int, visited map[int]bool, curr int) int {
	if visited[curr] {
		return 0
	}
	visited[curr] = true

	size := 1
	for _, neighbor := range graph[curr] {
		size += exploreSize(graph, visited, neighbor)
	}
	return size
}
