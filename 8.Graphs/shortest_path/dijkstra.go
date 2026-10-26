package shortestpath

import "heap"

type WeightedEdge struct {
	from   int
	to     int
	weight int
}

type neighbor struct {
	node   int
	weight int
}

func buildWeightedGraph(edges []WeightedEdge) map[int][]neighbor {
	graph := make(map[int][]neighbor)

	for _, edge := range edges {
		graph[edge.from] = append(graph[edge.from], neighbor{edge.to, edge.weight})
		graph[edge.to] = append(graph[edge.to], neighbor{edge.from, edge.weight})
	}
	return graph
}

// dijkstra returns the minimum total weight from src to dst, or -1 if
// dst is unreachable. Weights must be non-negative.
func dijkstra(edges []WeightedEdge, src int, dst int) int {
	graph := buildWeightedGraph(edges)
	// unlike BFS we can only settle a node when we POP it, not when we
	// push it: a node can be pushed via an expensive path before a
	// cheaper one is found
	settled := map[int]bool{}
	h := &heap.Heap{}
	h.Insert([2]int{src, 0}, 0)
	for !h.IsEmpty() {
		v, _ := h.Delete()
		curr := v.([2]int)
		node, dist := curr[0], curr[1]
		if settled[node] {
			continue
		}
		settled[node] = true
		if node == dst {
			return dist
		}
		for _, nb := range graph[node] {
			if !settled[nb.node] {
				h.Insert([2]int{nb.node, dist + nb.weight}, dist+nb.weight)
			}
		}
	}

	return -1
}
