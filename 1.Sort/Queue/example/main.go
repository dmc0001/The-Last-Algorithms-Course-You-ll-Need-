package main

import (
	"fmt"
	"queue"
)

func main() {
	q := &queue.Queue{}

	q.Enqueue(1)
	q.Enqueue(12)
	q.Enqueue(12)

	if v, ok := q.Peek(); ok {
		fmt.Printf("peek is: %v\n", v)
	}

	q.Enqueue(13)
	q.Enqueue(14)

	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		fmt.Println(v)
	}
}
