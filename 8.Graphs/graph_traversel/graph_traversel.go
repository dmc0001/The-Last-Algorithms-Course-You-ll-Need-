package graphtraversal

import (
	"fmt"
	"queue"
)

func DepthFirstSearch(graph map[int][]int, source int) {
	fmt.Println(source)
	neighbors := graph[source]
	if len(neighbors) == 0 {
		return
	}
	for _, neighbor := range neighbors {
		DepthFirstSearch(graph, neighbor)
	}
}

func BreathFirstSearch(graph map[int][]int, source int) {
	q := queue.Queue{}
	q.Enqueue(source)
	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		current := v.(int)
		fmt.Println(current)
		neighbors := graph[current]
		for _, neighbor := range neighbors {
			q.Enqueue(neighbor)
		}
	}
}
