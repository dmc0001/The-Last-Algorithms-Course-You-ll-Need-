package main

import (
	"fmt"
	"stack"
)

func main() {
	s := &stack.Stack{}

	s.Push(1)
	s.Push(1.23)
	s.Push("Haitham")
	s.Push(13)
	s.Push(14)

	if v, ok := s.Peek(); ok {
		fmt.Printf("peek is: %v\n", v)
	}

	for !s.IsEmpty() {
		v, _ := s.Pop()
		fmt.Println(v)
	}
}
