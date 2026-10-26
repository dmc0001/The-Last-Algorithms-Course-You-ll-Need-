package shortestpath

import "queue"

type Pair struct {
	first  int
	second int
}

func buildGraph(edges []Pair) map[int][]Pair {
	graph := make(map[int][]Pair)

	for _, edge := range edges {
		a, b := edge.first, edge.second
		graph[a] = append(graph[a], Pair{first: b, second: 1})
		graph[b] = append(graph[b], Pair{first: b, second: 1})
	}
	return graph
}

func shortestPath(edges []Pair, src int, dst int) int {
	graph := buildGraph(edges)
	visited := map[int]bool{src: true}
	q := queue.Queue{}
	q.Enqueue([2]int{src, 0})
	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		curr := v.([2]int)
		node, dist := curr[0], curr[1]
		if node == dst {
			return dist
		}
		for _, neighbor := range graph[node] {
			if !visited[neighbor.first] {
				visited[neighbor.first] = true
				q.Enqueue([2]int{neighbor.first, dist + neighbor.second})
			}

		}

	}

	return -1

}
