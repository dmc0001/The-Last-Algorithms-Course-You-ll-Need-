package main

import (
	"fmt"

	graphtraversal "graphs/graph_traversel"
	undirectedgraph "graphs/undirected_graph"
)

func main() {
	g := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {5},
		4: {6},
		5: {},
		6: {},
	}
	fmt.Println("===================DFS===================")
	graphtraversal.DepthFirstSearch(g, 1)
	fmt.Println("===================BFS===================")
	graphtraversal.BreathFirstSearch(g, 1)

	fmt.Println("===============Undirected Path===============")
	// Undirected graph:
	//   1 - 2 - 3
	//   |       |
	//   4       5
	//
	//   6 - 7   (separate component)
	edges := []undirectedgraph.Pair{
		undirectedgraph.NewPair(1, 2),
		undirectedgraph.NewPair(2, 3),
		undirectedgraph.NewPair(1, 4),
		undirectedgraph.NewPair(3, 5),
		undirectedgraph.NewPair(6, 7),
	}

	fmt.Println("adjacency list:", undirectedgraph.BuildGraph(edges))

	cases := [][2]int{
		{1, 5}, // true  (1->2->3->5)
		{4, 5}, // true  (4->1->2->3->5)
		{1, 6}, // false (different component)
		{6, 7}, // true
		{2, 2}, // true  (same node)
	}
	for _, c := range cases {
		dfs := undirectedgraph.HasPathDFS(edges, c[0], c[1])
		bfs := undirectedgraph.HasPathBFS(edges, c[0], c[1])
		fmt.Printf("%d -> %d : DFS=%v BFS=%v\n", c[0], c[1], dfs, bfs)
	}
}
