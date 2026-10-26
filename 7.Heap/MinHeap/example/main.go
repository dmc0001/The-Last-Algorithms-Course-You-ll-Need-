package main

import (
	"fmt"
	"heap"
)

func main() {
	h := &heap.Heap{}

	for _, v := range []int{5, 3, 8, 1, 9, 2, 7} {
		h.Insert(v, v)
	}

	if v, ok := h.Peek(); ok {
		fmt.Println("min after inserts:", v)
	}

	fmt.Print("extract order:")
	for !h.IsEmpty() {
		v, _ := h.Delete()
		fmt.Printf(" %d", v)
	}
	fmt.Println()
}
