package undirectedgraph

import (
	"queue"
)

func hasPathDFS(graph map[int][]int, src int, dst int, visited map[int]bool) bool {
	if src == dst {
		return true
	}
	if visited[src] {
		return false
	}
	visited[src] = true

	for _, neighbor := range graph[src] {
		if hasPathDFS(graph, neighbor, dst, visited) {
			return true
		}
	}
	return false
}

func hasPathBFS(graph map[int][]int, src int, dst int) bool {
	q := queue.Queue{}
	q.Enqueue(src)
	seen := map[int]bool{src: true}

	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		curr := v.(int)
		if curr == dst {
			return true
		}
		for _, neighbor := range graph[curr] {
			if !seen[neighbor] {
				seen[neighbor] = true
				q.Enqueue(neighbor)
			}
		}
	}
	return false
}

type Pair struct {
	first  int
	second int
}

func buildGraph(edges []Pair) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge.first, edge.second
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	return graph
}

// NewPair is a small helper so callers in other packages can build edges
// even though the Pair fields are unexported.
func NewPair(a, b int) Pair {
	return Pair{first: a, second: b}
}

// BuildGraph builds an undirected adjacency list from a list of edges.
func BuildGraph(edges []Pair) map[int][]int {
	return buildGraph(edges)
}

// HasPathDFS reports whether nodeB is reachable from nodeA (depth-first).
func HasPathDFS(edges []Pair, nodeA int, nodeB int) bool {
	graph := buildGraph(edges)
	return hasPathDFS(graph, nodeA, nodeB, map[int]bool{})
}

// HasPathBFS reports whether nodeB is reachable from nodeA (breadth-first).
func HasPathBFS(edges []Pair, nodeA int, nodeB int) bool {
	graph := buildGraph(edges)
	return hasPathBFS(graph, nodeA, nodeB)
}
